package chardet

const (
	SPACE = 0x20
	// windows-1255 / ISO-8859-8 code points of interest
	FINAL_KAF    = 0xEA
	NORMAL_KAF   = 0xEB
	FINAL_MEM    = 0xED
	NORMAL_MEM   = 0xEE
	FINAL_NUN    = 0xEF
	NORMAL_NUN   = 0xF0
	FINAL_PE     = 0xF3
	NORMAL_PE    = 0xF4
	FINAL_TSADI  = 0xF5
	NORMAL_TSADI = 0xF6

	// Minimum Visual vs Logical final letter score difference.
	// If the difference is below this, don't rely solely on the final letter score
	// distance.
	MIN_FINAL_CHAR_DISTANCE = 5

	// Minimum Visual vs Logical model score difference.
	// If the difference is below this, don't rely at all on the model score
	// distance.
	MIN_MODEL_DISTANCE = 0.01

	VISUAL_HEBREW_NAME  = "ISO-8859-8"
	LOGICAL_HEBREW_NAME = "windows-1255"
)

type HebrewProber struct {
	CommonCharsetProber
	finalCharLogicalScore int
	finalCharVisualScore  int
	prev                  int
	beforePrev            int
	logicalProber         *SingleByteCharSetProber
	visualProber          *SingleByteCharSetProber
}

func NewHebrewProber() *HebrewProber {
	var h HebrewProber
	h.finalCharVisualScore = 0
	h.finalCharLogicalScore = 0
	h.prev = SPACE
	h.beforePrev = SPACE
	h.logicalProber = nil
	h.visualProber = nil
	return &h
}

func (h *HebrewProber) reset() {
	h.finalCharLogicalScore = 0
	h.finalCharVisualScore = 0
	h.prev = SPACE
	h.beforePrev = SPACE
}

func (h *HebrewProber) setModelProbers(logicalProber *SingleByteCharSetProber, visualProber *SingleByteCharSetProber) {
	h.logicalProber = logicalProber
	h.visualProber = visualProber
}

func (h *HebrewProber) isFinal(c int) bool {
	return c == FINAL_KAF || c == FINAL_PE || c == FINAL_MEM || c == FINAL_NUN || c == FINAL_TSADI
}

func (h *HebrewProber) isNonFinal(c int) bool {
	return c == NORMAL_KAF || c == NORMAL_MEM || c == NORMAL_NUN || c == NORMAL_PE
}

func (h *HebrewProber) feed(data []byte) ProbingState {
	if h.state_ == NOT_ME {
		return NOT_ME
	}
	data = filterHighByteOnly(data)
	for _, c := range data {
		if c == SPACE {
			if h.beforePrev != SPACE {
				if h.isFinal(h.prev) {
					h.finalCharLogicalScore += 1
				} else if h.isNonFinal(h.prev) {
					h.finalCharVisualScore += 1
				}
			}
		} else {
			if h.beforePrev == SPACE && h.isFinal(h.prev) && c != SPACE {
				h.finalCharVisualScore += 1
			}
		}
	}
	return DETECTING
}

func (h *HebrewProber) charName() string {
	finalSub := h.finalCharLogicalScore - h.finalCharVisualScore
	if finalSub >= MIN_FINAL_CHAR_DISTANCE {
		return LOGICAL_HEBREW_NAME
	}
	if finalSub <= -MIN_FINAL_CHAR_DISTANCE {
		return VISUAL_HEBREW_NAME
	}
	modelSub := h.logicalProber.getConfidence() - h.visualProber.getConfidence()
	if modelSub > MIN_MODEL_DISTANCE {
		return LOGICAL_HEBREW_NAME
	}
	if modelSub < -MIN_MODEL_DISTANCE {
		return VISUAL_HEBREW_NAME
	}
	if finalSub < 0.0 {
		return VISUAL_HEBREW_NAME
	}
	return LOGICAL_HEBREW_NAME
}

func (h *HebrewProber) language() string {
	return "Hebrew"
}

func (h *HebrewProber) state() ProbingState {
	if h.logicalProber.state_ == NOT_ME && h.visualProber.state_ == NOT_ME {
		return NOT_ME
	}
	return DETECTING
}
