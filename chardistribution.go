package chardet

const (
	ENOUGH_DATA_THRESHOLD  = 1024
	SURE_YES               = 0.99
	SURE_NO                = 0.01
	MINIMUM_DATA_THRESHOLD = 3
)

type LanguageAnalysis interface {
	reset()
	getEnoughData() bool
	feed(data []byte, charLen int)
	getConfidence() float64
}

type CharDistributionAnalysis struct {
	charToFreqOrder          []int
	tableSize                int
	typicalDistributionRatio float64
	done                     bool
	totalChars               int
	freqChars                int
}

func (c *CharDistributionAnalysis) init() {
	c.charToFreqOrder = make([]int, 0)
	c.reset()
}

func (c *CharDistributionAnalysis) reset() {
	c.done = false
	c.totalChars = 0
	c.freqChars = 0
}

func (c *CharDistributionAnalysis) feed(data []byte, charLen int) {
	order := 0
	if charLen == 2 {
		order = c.getOrder(data)
	} else {
		order = -1
	}
	if order >= 0 {
		c.totalChars += 1
		if order < c.tableSize {
			if c.charToFreqOrder[order] < 512 {
				c.freqChars += 1
			}
		}
	}
}

func (c *CharDistributionAnalysis) getConfidence() float64 {
	if c.totalChars <= 0 || c.freqChars <= MINIMUM_DATA_THRESHOLD {
		return SURE_NO
	}
	if c.totalChars != c.freqChars {
		r := float64(c.freqChars) / (float64(c.totalChars-c.freqChars) * c.typicalDistributionRatio)
		if r < SURE_YES {
			return r
		}
	}
	return SURE_YES
}

func (c *CharDistributionAnalysis) getEnoughData() bool {
	return c.totalChars > ENOUGH_DATA_THRESHOLD
}

func (c *CharDistributionAnalysis) getOrder(data []byte) int {
	return -1
}

type Big5DistributionAnalysis struct {
	CharDistributionAnalysis
}

func NewBig5DistributionAnalysis() LanguageAnalysis {
	g := &Big5DistributionAnalysis{}
	g.init()
	g.charToFreqOrder = BIG5_CHAR_TO_FREQ_ORDER
	g.tableSize = BIG5_TABLE_SIZE
	g.typicalDistributionRatio = BIG5_TYPICAL_DISTRIBUTION_RATIO
	return g
}

func (c *Big5DistributionAnalysis) feed(data []byte, charLen int) {
	order := 0
	if charLen == 2 {
		order = c.getOrder(data)
	} else {
		order = -1
	}
	if order >= 0 {
		c.totalChars += 1
		if order < c.tableSize {
			if c.charToFreqOrder[order] < 512 {
				c.freqChars += 1
			}
		}
	}
}

func (g *Big5DistributionAnalysis) getOrder(data []byte) int {
	firstChar, secondChar := data[0], data[1]
	if firstChar >= 0xA4 {
		if secondChar >= 0xA1 {
			return 157*int(firstChar-0xA4) + int(secondChar-0xA1) + 63
		}
		return 157*int(firstChar-0xA4) + int(secondChar-0x40)
	}
	return -1
}

type GB2312DistributionAnalysis struct {
	CharDistributionAnalysis
}

func NewGB2312DistributionAnalysis() LanguageAnalysis {
	g := &GB2312DistributionAnalysis{}
	g.init()
	g.charToFreqOrder = GB2312_CHAR_TO_FREQ_ORDER
	g.tableSize = GB2312_TABLE_SIZE
	g.typicalDistributionRatio = GB2312_TYPICAL_DISTRIBUTION_RATIO
	return g
}

func (c *GB2312DistributionAnalysis) feed(data []byte, charLen int) {
	order := 0
	if charLen == 2 {
		order = c.getOrder(data)
	} else {
		order = -1
	}
	if order >= 0 {
		c.totalChars += 1
		if order < c.tableSize {
			if c.charToFreqOrder[order] < 512 {
				c.freqChars += 1
			}
		}
	}
}

func (g *GB2312DistributionAnalysis) getOrder(data []byte) int {
	firstChar, secondChar := data[0], data[1]
	if firstChar >= 0xB0 && secondChar >= 0xA1 {
		return 94*int(firstChar-0xB0) + int(secondChar-0xA1)
	}
	return -1
}

type SJISDistributionAnalysis struct {
	CharDistributionAnalysis
}

func NewSJISDistributionAnalysis() LanguageAnalysis {
	g := &SJISDistributionAnalysis{}
	g.init()
	g.charToFreqOrder = JIS_CHAR_TO_FREQ_ORDER
	g.tableSize = JIS_TABLE_SIZE
	g.typicalDistributionRatio = JIS_TYPICAL_DISTRIBUTION_RATIO
	return g
}

func (c *SJISDistributionAnalysis) feed(data []byte, charLen int) {
	order := 0
	if charLen == 2 {
		order = c.getOrder(data)
	} else {
		order = -1
	}
	if order >= 0 {
		c.totalChars += 1
		if order < c.tableSize {
			if c.charToFreqOrder[order] < 512 {
				c.freqChars += 1
			}
		}
	}
}

func (g *SJISDistributionAnalysis) getOrder(data []byte) int {
	firstChar, secondChar := data[0], data[1]
	var order int
	if firstChar >= 0x81 && firstChar <= 0x9F {
		order = 188 * int(firstChar-0x81)
	} else if firstChar >= 0xE0 && firstChar <= 0xEF {
		order = 188 * int(firstChar-0xE0+31)
	} else {
		return -1
	}
	order = order + int(secondChar-0x40)
	if secondChar > 0x7F {
		order = -1
	}
	return order
}

type EUCTWDistributionAnalysis struct {
	CharDistributionAnalysis
}

func NewEUCTWDistributionAnalysis() LanguageAnalysis {
	g := &EUCTWDistributionAnalysis{}
	g.init()
	g.charToFreqOrder = EUCTW_CHAR_TO_FREQ_ORDER
	g.tableSize = EUCTW_TABLE_SIZE
	g.typicalDistributionRatio = EUCTW_TYPICAL_DISTRIBUTION_RATIO
	return g
}

func (c *EUCTWDistributionAnalysis) feed(data []byte, charLen int) {
	order := 0
	if charLen == 2 {
		order = c.getOrder(data)
	} else {
		order = -1
	}
	if order >= 0 {
		c.totalChars += 1
		if order < c.tableSize {
			if c.charToFreqOrder[order] < 512 {
				c.freqChars += 1
			}
		}
	}
}

func (g *EUCTWDistributionAnalysis) getOrder(data []byte) int {
	firstChar := data[0]
	if firstChar >= 0xC4 {
		return 94*int(firstChar-0xC4) + int(data[1]-0xA1)
	}
	return -1
}

type EUCKRDistributionAnalysis struct {
	CharDistributionAnalysis
}

func NewEUCKRDistributionAnalysis() LanguageAnalysis {
	g := &EUCKRDistributionAnalysis{}
	g.init()
	g.charToFreqOrder = EUCKR_CHAR_TO_FREQ_ORDER
	g.tableSize = EUCKR_TABLE_SIZE
	g.typicalDistributionRatio = EUCKR_TYPICAL_DISTRIBUTION_RATIO
	return g
}

func (c *EUCKRDistributionAnalysis) feed(data []byte, charLen int) {
	order := 0
	if charLen == 2 {
		order = c.getOrder(data)
	} else {
		order = -1
	}
	if order >= 0 {
		c.totalChars += 1
		if order < c.tableSize {
			if c.charToFreqOrder[order] < 512 {
				c.freqChars += 1
			}
		}
	}
}

func (g *EUCKRDistributionAnalysis) getOrder(data []byte) int {
	firstChar := data[0]
	if firstChar >= 0xB0 {
		return 94*int(firstChar-0xB0) + int(data[1]-0xA1)
	}
	return -1
}

type JOHABDistributionAnalysis struct {
	CharDistributionAnalysis
}

func NewJOHABDistributionAnalysis() LanguageAnalysis {
	g := &JOHABDistributionAnalysis{}
	g.init()
	g.charToFreqOrder = EUCKR_CHAR_TO_FREQ_ORDER
	g.tableSize = EUCKR_TABLE_SIZE
	g.typicalDistributionRatio = EUCKR_TYPICAL_DISTRIBUTION_RATIO
	return g
}

func (c *JOHABDistributionAnalysis) feed(data []byte, charLen int) {
	order := 0
	if charLen == 2 {
		order = c.getOrder(data)
	} else {
		order = -1
	}
	if order >= 0 {
		c.totalChars += 1
		if order < c.tableSize {
			if c.charToFreqOrder[order] < 512 {
				c.freqChars += 1
			}
		}
	}
}

func (g *JOHABDistributionAnalysis) getOrder(data []byte) int {
	firstChar := data[0]
	if firstChar >= 0x88 && firstChar < 0xD4 {
		code := int(firstChar)*256 + int(data[1])
		o, ok := JOHAB_TO_EUCKR_ORDER_TABLE[code]
		if ok {
			return o
		}
		return -1
	}
	return -1
}

type EUCJPDistributionAnalysis struct {
	CharDistributionAnalysis
}

func NewEUCJPDistributionAnalysis() LanguageAnalysis {
	g := &EUCJPDistributionAnalysis{}
	g.init()
	g.charToFreqOrder = JIS_CHAR_TO_FREQ_ORDER
	g.tableSize = JIS_TABLE_SIZE
	g.typicalDistributionRatio = JIS_TYPICAL_DISTRIBUTION_RATIO
	return g
}

func (c *EUCJPDistributionAnalysis) feed(data []byte, charLen int) {
	order := 0
	if charLen == 2 {
		order = c.getOrder(data)
	} else {
		order = -1
	}
	if order >= 0 {
		c.totalChars += 1
		if order < c.tableSize {
			if c.charToFreqOrder[order] < 512 {
				c.freqChars += 1
			}
		}
	}
}

func (g *EUCJPDistributionAnalysis) getOrder(data []byte) int {
	firstChar := data[0]
	if firstChar >= 0xA0 {
		return 94*int(firstChar-0xA1) + int(data[1]-0xA1)
	}
	return -1
}
