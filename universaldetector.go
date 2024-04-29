package chardet

import (
	"bytes"
	"regexp"
	"strings"
)

var (
	BOM_UTF8     = []byte("\xef\xbb\xbf")
	BOM_LE       = []byte("\xff\xfe")
	BOM_BE       = []byte("\xfe\xff")
	BOM_UTF32_LE = []byte("\xff\xfe\x00\x00")
	BOM_UTF32_BE = []byte("\x00\x00\xfe\xff")
)

var (
	MINIMUM_THRESHOLD     = 0.20
	HIGH_BYTE_DETECTOR, _ = regexp.Compile(`[^\x00-\x7F]`)
	ESC_DETECTOR, _       = regexp.Compile(`(\033|~{)`)
	WIN_BYTE_DETECTOR, _  = regexp.Compile(`[\x80-\x9F]`)
)

var ISO_WIN_MAP = map[string]string{
	"iso-8859-1":  "Windows-1252",
	"iso-8859-2":  "Windows-1250",
	"iso-8859-5":  "Windows-1251",
	"iso-8859-6":  "Windows-1256",
	"iso-8859-7":  "Windows-1253",
	"iso-8859-8":  "Windows-1255",
	"iso-8859-9":  "Windows-1254",
	"iso-8859-13": "Windows-1257",
}

var LEGACY_MAP = map[string]string{
	"ascii":      "Windows-1252",
	"iso-8859-1": "Windows-1252",
	"tis-620":    "ISO-8859-11",
	"iso-8859-9": "Windows-1254",
	"gb2312":     "GB18030",
	"euc-kr":     "CP949",
	"utf-16le":   "UTF-16",
}

type UniversalDetector struct {
	hasWinBytes        bool
	done               bool
	gotData            bool
	inputState         InputState
	langFilter         LanguageFilter
	lastChar           []byte
	shouldRenameLegacy bool
	charsetProbers     []CharsetProber
	result             ResultDict
	escCharsetProber   *EscCharSetProber
	utf1632Prober      *UTF1632Prober
	allProbers         []CharsetProber
}

func NewUniversalDetector(filter LanguageFilter, shouldRenameLegacy bool) *UniversalDetector {
	var u UniversalDetector
	u.charsetProbers = make([]CharsetProber, 0)
	u.shouldRenameLegacy = shouldRenameLegacy
	u.langFilter = filter
	u.Reset()
	return &u
}

func (u *UniversalDetector) Reset() {
	u.result = ResultDict{
		encoding:   "",
		confidence: 0,
		language:   "",
	}
	u.done = false
	u.gotData = false
	u.hasWinBytes = false
	u.inputState = PURE_ASCII
	u.lastChar = []byte("")
	if u.escCharsetProber != nil {
		u.escCharsetProber.reset()
	}
	if u.utf1632Prober != nil {
		u.utf1632Prober.reset()
	}
	for _, prober := range u.charsetProbers {
		prober.reset()
	}
}

func (u *UniversalDetector) InputState() InputState {
	return u.inputState
}

func (u *UniversalDetector) HasWinBytes() bool {
	return u.hasWinBytes
}

func (u *UniversalDetector) CharSetProbers() []CharsetProber {
	return u.charsetProbers
}

func (u *UniversalDetector) Feed(data []byte) {
	if u.done {
		return
	}
	if len(data) == 0 {
		return
	}
	if !u.gotData {
		if bytes.HasPrefix(data, BOM_UTF8) {
			u.result = ResultDict{
				confidence: 1.0,
				encoding:   "UTF-8",
				language:   "",
			}
		} else if bytes.HasPrefix(data, BOM_UTF32_LE) || bytes.HasPrefix(data, BOM_UTF32_BE) {
			u.result = ResultDict{
				confidence: 1.0,
				encoding:   "UTF-32",
				language:   "",
			}
		} else if bytes.HasPrefix(data, []byte("\xFE\xFF\x00\x00")) {
			u.result = ResultDict{
				confidence: 1.0,
				encoding:   "X-ISO-10646-UCS-4-3412",
				language:   "",
			}
		} else if bytes.HasPrefix(data, []byte("\x00\x00\xFF\xFE")) {
			u.result = ResultDict{
				confidence: 1.0,
				encoding:   "X-ISO-10646-UCS-4-2143",
				language:   "",
			}
		} else if bytes.HasPrefix(data, BOM_LE) || bytes.HasPrefix(data, BOM_BE) {
			u.result = ResultDict{
				confidence: 1.0,
				encoding:   "UTF-16",
				language:   "",
			}
		}
		u.gotData = true
		if u.result.encoding != "" {
			u.done = true
			return
		}
	}

	//判断是否是非ascii
	if u.inputState == PURE_ASCII {
		if HIGH_BYTE_DETECTOR.Match(data) {
			u.inputState = HIGH_BYTE
		} else if u.inputState == PURE_ASCII && ESC_DETECTOR.Match(data) {
			u.inputState = ESC_ASCII
		}
	}
	u.lastChar = data[len(data)-1:]
	if u.utf1632Prober == nil {
		b := NewUTF1632Prober()
		u.utf1632Prober = &b
	}

	if u.utf1632Prober.state_ == DETECTING {
		if u.utf1632Prober.feed(data) == FOUND_IT {
			u.result = ResultDict{
				encoding:   u.utf1632Prober.charName(),
				confidence: u.utf1632Prober.getConfidence(),
				language:   "",
			}
			u.done = true
			return
		}
	}

	if u.inputState == ESC_ASCII {
		if u.escCharsetProber == nil {
			b := NewEscCharSetProber(u.langFilter)
			u.escCharsetProber = &b
		}
		if u.escCharsetProber.feed(data) == FOUND_IT {
			u.result = ResultDict{
				encoding:   u.escCharsetProber.charName(),
				confidence: u.escCharsetProber.getConfidence(),
				language:   u.escCharsetProber.language(),
			}
			u.done = true
		}
	} else if u.inputState == HIGH_BYTE {
		if len(u.charsetProbers) == 0 {
			mbc := NewMBCSGroupProber(u.langFilter)
			u.charsetProbers = append(u.charsetProbers, mbc)
			u.allProbers = append(u.allProbers, mbc.getProbers()...)
			if u.langFilter&NON_CJK != 0 {
				sbc := NewSBCSGroupProber()
				u.charsetProbers = append(u.charsetProbers, sbc)
				u.allProbers = append(u.allProbers, sbc.getProbers()...)
			}
			latin1 := NewLatin1Prober()
			u.charsetProbers = append(u.charsetProbers, latin1)
			roman := NewMacRomanProber()
			u.charsetProbers = append(u.charsetProbers, roman)
			u.allProbers = append(u.allProbers, latin1)
			u.allProbers = append(u.allProbers, roman)
		}
		for _, prober := range u.charsetProbers {
			if prober.feed(data) == FOUND_IT {
				u.result = ResultDict{
					encoding:   prober.charName(),
					confidence: prober.getConfidence(),
					language:   prober.language(),
				}
				u.done = true
				break
			}
		}
		if winByteMatch(data) {
			u.hasWinBytes = true
		}
	}
}

func (u *UniversalDetector) Close() ResultDict {
	if u.done {
		return u.result
	}
	u.done = true

	if !u.gotData {

	} else if u.inputState == PURE_ASCII {
		u.result = ResultDict{
			encoding:   "utf-8",
			confidence: 1.0,
			language:   "",
		}
	} else if u.inputState == HIGH_BYTE {
		maxProberConfidence := 0.0
		var maxprober CharsetProber
		for _, prober := range u.charsetProbers {
			if prober.getConfidence() > maxProberConfidence {
				maxProberConfidence = prober.getConfidence()
				maxprober = prober
			}
		}
		if maxprober != nil && maxProberConfidence > MINIMUM_THRESHOLD {
			name := maxprober.charName()
			lowerCharsetName := strings.ToLower(maxprober.charName())
			if strings.HasPrefix(lowerCharsetName, "iso-8859") {
				if u.hasWinBytes {
					nameTmp, ok := ISO_WIN_MAP[lowerCharsetName]
					if ok {
						name = nameTmp
					}
				}
			}
			if u.shouldRenameLegacy {
				nameTmp, ok := LEGACY_MAP[strings.ToLower(name)]
				if ok {
					name = nameTmp
				}
			}
			u.result = ResultDict{
				encoding:   name,
				confidence: maxProberConfidence,
				language:   maxprober.language(),
			}
		}
	}
	return u.result
}

func winByteMatch(data []byte) bool {
	for _, x := range data {
		if x >= 0x80 && x <= 0x9F {
			return true
		}
	}
	return false
}
