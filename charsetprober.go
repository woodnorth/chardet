package chardet

import (
	"regexp"
)

const (
	SHORTCUT_THRESHOLD = 0.95
)

var (
	INTERNATIONAL_WORDS_PATTERN, _ = regexp.Compile(`[a-zA-Z]*[^\x00-\x7F]+[a-zA-Z]*[^a-zA-Z\x80-\xFF]?`)
)

type CharsetProber interface {
	reset()
	charName() string
	language() string
	feed([]byte) ProbingState
	state() ProbingState
	getConfidence() float64
	isActive() bool
	disActive()
	enActive()
}

type CommonCharsetProber struct {
	active     bool
	state_     ProbingState
	langFilter LanguageFilter
}

func (c *CommonCharsetProber) init(filetr LanguageFilter) {
	c.state_ = DETECTING
	c.active = true
	c.langFilter = filetr
}

func (c *CommonCharsetProber) reset() {
	c.state_ = DETECTING
	c.active = true
}

func (c *CommonCharsetProber) feed([]byte) ProbingState {
	return 0
}

func (c *CommonCharsetProber) charName() string {
	return ""
}

func (c *CommonCharsetProber) language() string {
	return ""
}

func (c *CommonCharsetProber) state() ProbingState {
	return c.state_
}

func (c *CommonCharsetProber) getConfidence() float64 {
	return 0
}

func (c *CommonCharsetProber) isActive() bool {
	return c.active
}

func (c *CommonCharsetProber) disActive() {
	c.active = false
}

func (c *CommonCharsetProber) enActive() {
	c.active = true
}

func filterHighByteOnly(data []byte) []byte {
	re, _ := regexp.Compile("[\x00-\x7F]+")
	return re.ReplaceAll(data, []byte(" "))
}

func removeXmlTags(buf []byte) []byte {
	res := make([]byte, 0)
	prev := 0
	inTag := false
	for cur, c := range buf {
		if c == '>' {
			prev = cur + 1
			inTag = false
		} else if c == '<' {
			if cur > prev && !inTag {
				res = append(res, buf[prev:cur]...)
				res = append(res, []byte(" ")...)
			}
			inTag = true
		}
	}
	if !inTag {
		res = append(res, buf[prev:]...)
	}
	return res
}

func isEnLetter(s byte) bool {
	return (s >= 97 && s <= 122) || (s >= 65 && s <= 90)
}

func filterInternationalWords(buf []byte) []byte {
	res := make([]byte, 0)
	words := INTERNATIONAL_WORDS_PATTERN.FindAll(buf, -1)
	if words != nil {
		for _, word := range words {
			res = append(res, word...)
			lastChar := word[len(word)-1]
			if isEnLetter(lastChar) && lastChar < '\x80' {
				lastChar = ' '
			}
			res = append(res, lastChar)
		}
	}
	return res
}
