package chardet

const (
	ODD         = 8 // character that is unlikely to appear
	CLASS_NUM_M = 9 // total classes
)

// The change from Latin1 is that we explicitly look for extended characters
// that are infrequently-occurring symbols, and consider them to always be
// improbable. This should let MacRoman get out of the way of more likely
// encodings in most situations.

// fmt: off
var MacRoman_CharToClass = []int{
	OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // 00 - 07
	OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // 08 - 0F
	OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // 10 - 17
	OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // 18 - 1F
	OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // 20 - 27
	OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // 28 - 2F
	OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // 30 - 37
	OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // 38 - 3F
	OTH, ASC, ASC, ASC, ASC, ASC, ASC, ASC, // 40 - 47
	ASC, ASC, ASC, ASC, ASC, ASC, ASC, ASC, // 48 - 4F
	ASC, ASC, ASC, ASC, ASC, ASC, ASC, ASC, // 50 - 57
	ASC, ASC, ASC, OTH, OTH, OTH, OTH, OTH, // 58 - 5F
	OTH, ASS, ASS, ASS, ASS, ASS, ASS, ASS, // 60 - 67
	ASS, ASS, ASS, ASS, ASS, ASS, ASS, ASS, // 68 - 6F
	ASS, ASS, ASS, ASS, ASS, ASS, ASS, ASS, // 70 - 77
	ASS, ASS, ASS, OTH, OTH, OTH, OTH, OTH, // 78 - 7F
	ACV, ACV, ACO, ACV, ACO, ACV, ACV, ASV, // 80 - 87
	ASV, ASV, ASV, ASV, ASV, ASO, ASV, ASV, // 88 - 8F
	ASV, ASV, ASV, ASV, ASV, ASV, ASO, ASV, // 90 - 97
	ASV, ASV, ASV, ASV, ASV, ASV, ASV, ASV, // 98 - 9F
	OTH, OTH, OTH, OTH, OTH, OTH, OTH, ASO, // A0 - A7
	OTH, OTH, ODD, ODD, OTH, OTH, ACV, ACV, // A8 - AF
	OTH, OTH, OTH, OTH, OTH, OTH, OTH, OTH, // B0 - B7
	OTH, OTH, OTH, OTH, OTH, OTH, ASV, ASV, // B8 - BF
	OTH, OTH, ODD, OTH, ODD, OTH, OTH, OTH, // C0 - C7
	OTH, OTH, OTH, ACV, ACV, ACV, ACV, ASV, // C8 - CF
	OTH, OTH, OTH, OTH, OTH, OTH, OTH, ODD, // D0 - D7
	ASV, ACV, ODD, OTH, OTH, OTH, OTH, OTH, // D8 - DF
	OTH, OTH, OTH, OTH, OTH, ACV, ACV, ACV, // E0 - E7
	ACV, ACV, ACV, ACV, ACV, ACV, ACV, ACV, // E8 - EF
	ODD, ACV, ACV, ACV, ACV, ASV, ODD, ODD, // F0 - F7
	ODD, ODD, ODD, ODD, ODD, ODD, ODD, ODD, // F8 - FF
}

// 0 : illegal
// 1 : very unlikely
// 2 : normal
// 3 : very likely
var MacRomanClassModel = []int{
	// UDF OTH ASC ASS ACV ACO ASV ASO ODD
	0, 0, 0, 0, 0, 0, 0, 0, 0, // UDF
	0, 3, 3, 3, 3, 3, 3, 3, 1, // OTH
	0, 3, 3, 3, 3, 3, 3, 3, 1, // ASC
	0, 3, 3, 3, 1, 1, 3, 3, 1, // ASS
	0, 3, 3, 3, 1, 2, 1, 2, 1, // ACV
	0, 3, 3, 3, 3, 3, 3, 3, 1, // ACO
	0, 3, 1, 3, 1, 1, 1, 3, 1, // ASV
	0, 3, 1, 3, 1, 1, 3, 3, 1, // ASO
	0, 1, 1, 1, 1, 1, 1, 1, 1, // ODD
}

// fmt: on

type MacRomanProber struct {
	CommonCharsetProber
	lastCharClass int
	freqCounter   []float64
}

func NewMacRomanProber() CharsetProber {
	l := &MacRomanProber{
		freqCounter: make([]float64, 0),
	}
	l.init(NONE)
	l.lastCharClass = OTH
	l.reset()
	return l
}

func (l *MacRomanProber) reset() {
	l.lastCharClass = OTH
	l.freqCounter = make([]float64, FREQ_CAT_NUM)
	l.CommonCharsetProber.reset()
}

func (l *MacRomanProber) charName() string {
	return "MacRoman"
}

func (l *MacRomanProber) language() string {
	return ""
}

func (l *MacRomanProber) feed(data []byte) ProbingState {
	buf := removeXmlTags(data)
	for _, c := range buf {
		charClass := MacRoman_CharToClass[int(c)]
		freq := MacRomanClassModel[l.lastCharClass*CLASS_NUM_M+charClass]
		if freq == 0 {
			l.state_ = NOT_ME
			break
		}
		l.freqCounter[freq] += 1
		l.lastCharClass = charClass
	}
	return l.state_
}

func (l *MacRomanProber) getConfidence() float64 {
	if l.state_ == NOT_ME {
		return 0.01
	}
	var confidence float64
	var total float64
	for _, t := range l.freqCounter {
		total += t
	}
	if float64(total) < 0.01 {
		confidence = 0
	} else {
		confidence = (l.freqCounter[3] - l.freqCounter[1]*20.0) / total
	}
	confidence = max(confidence, 0.0)

	confidence *= 0.73
	return confidence
}
