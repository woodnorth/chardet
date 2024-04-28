package chardet

type SingleByteCharSetModel struct {
	charsetName          string
	language             string
	charToOrdermap       map[int]int
	languageModel        map[int]map[int]int
	typicalPositiveRatio float64
	keepAsciiLetters     bool
	alphabet             string
}

const (
	SAMPLE_SIZE                 = 64
	SB_ENOUGH_REL_THRESHOLD     = 1024 // 0.25 * SAMPLE_SIZE^2
	POSITIVE_SHORTCUT_THRESHOLD = 0.95
	NEGATIVE_SHORTCUT_THRESHOLD = 0.05
)

type SingleByteCharSetProber struct {
	CommonCharsetProber
	model       SingleByteCharSetModel
	isReserved  bool
	nameProber  CharsetProber
	lastOrder   int
	seqCounters []int
	totalSeqs   int
	totalChar   int
	controlChar int
	freqChar    int
}

func NewSingleByteCharSetProber(model SingleByteCharSetModel, isReserved bool, nameProber CharsetProber) *SingleByteCharSetProber {
	var s = SingleByteCharSetProber{
		model:       model,
		isReserved:  isReserved,
		lastOrder:   255,
		nameProber:  nameProber,
		seqCounters: make([]int, 0),
		totalSeqs:   0,
		totalChar:   0,
		controlChar: 0,
		freqChar:    0,
	}
	s.CommonCharsetProber.reset()
	return &s
}

func (s *SingleByteCharSetProber) reset() {
	s.CommonCharsetProber.reset()
	s.lastOrder = 255
	s.seqCounters = make([]int, 4)
	s.totalSeqs = 0
	s.totalChar = 0
	s.controlChar = 0
	s.freqChar = 0
}

func (s *SingleByteCharSetProber) charName() string {
	if s.nameProber != nil {
		return s.nameProber.charName()
	}
	return s.model.charsetName
}

func (s *SingleByteCharSetProber) language() string {
	if s.nameProber != nil {
		return s.nameProber.language()
	}
	return s.model.language
}

func (s *SingleByteCharSetProber) feed(data []byte) ProbingState {
	var bytes []byte
	if !s.model.keepAsciiLetters {
		bytes = filterInternationalWords(data)
	} else {
		bytes = removeXmlTags(data)
	}
	if len(bytes) == 0 {
		return s.state_
	}
	for _, c := range bytes {
		order, ok := s.model.charToOrdermap[int(c)]
		if !ok {
			order = UNDEFINED
		}
		if order < CONTROL {
			s.totalChar += 1
		}
		if order < SAMPLE_SIZE {
			s.freqChar += 1
			if s.lastOrder < SAMPLE_SIZE {
				s.totalSeqs += 1
				lmCat := s.model.languageModel[order][s.lastOrder]
				if !s.isReserved {
					lmCat = s.model.languageModel[s.lastOrder][order]
				}
				s.seqCounters[lmCat] += 1
			}
		}
		s.lastOrder = order
	}
	if s.state_ == DETECTING {
		if s.totalSeqs > SB_ENOUGH_REL_THRESHOLD {
			confidence := s.getConfidence()
			if confidence > POSITIVE_SHORTCUT_THRESHOLD {
				s.state_ = FOUND_IT
			} else if confidence < NEGATIVE_SHORTCUT_THRESHOLD {
				s.state_ = NOT_ME
			}
		}
	}
	return s.state_

}

func (s *SingleByteCharSetProber) getConfidence() float64 {
	r := 0.01
	if s.totalSeqs > 0 {
		r = (float64(s.seqCounters[POSITIVE]) + 0.25*float64(s.seqCounters[LIKELY])) / float64(s.totalSeqs) / s.model.typicalPositiveRatio
		r = r * float64((s.totalChar - s.controlChar)) / float64(s.totalChar)
		r = r * float64(s.freqChar) / float64(s.totalChar)
		if r >= 1.0 {
			r = 0.99
		}
	}
	return r
}
