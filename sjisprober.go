package chardet

type SJISProber struct {
	MultiByteCharSetProber
	contextAnalyzer *SJISContextAnalysis
}

func NewSJISProber() *SJISProber {
	var s SJISProber
	s.MultiByteCharSetProber.init(NONE)
	s.codingSm = NewCodingStateMachine(SJIS_SM_MODEL)
	s.distributionAnalyzer = NewSJISDistributionAnalysis()
	s.contextAnalyzer = NewSJISContextAnalysis()
	s.reset()
	return &s
}

func (s *SJISProber) reset() {
	s.MultiByteCharSetProber.reset()
	s.contextAnalyzer.reset()
}

func (s *SJISProber) charName() string {
	return s.contextAnalyzer.charsetName
}

func (s *SJISProber) language() string {
	return "Japanese"
}

func (s *SJISProber) feed(data []byte) ProbingState {
	for i, b := range data {
		codingState := s.codingSm.nextState(int(b))
		if codingState == ERROR {
			s.state_ = NOT_ME
			break
		}
		if codingState == ITS_ME {
			s.state_ = FOUND_IT
			break
		}
		if codingState == START {
			charLen := s.codingSm.getCurrentCharLen()
			if i == 0 {
				s.lastChar[1] = b
				s.contextAnalyzer.feed(s.lastChar[2-charLen:], charLen)
				s.distributionAnalyzer.feed(s.lastChar, charLen)
			} else {
				s.contextAnalyzer.feed(data[i+1-charLen:i+3-charLen], charLen)
				s.distributionAnalyzer.feed(data[i-1:i+1], charLen)
			}
		}
	}
	s.lastChar[0] = data[len(data)-1]
	if s.state_ == DETECTING {
		if s.contextAnalyzer.getEnoughData() && s.getConfidence() > SHORTCUT_THRESHOLD {
			s.state_ = FOUND_IT
		}
	}
	return s.state_
}

func (s *SJISProber) getConfidence() float64 {
	return max(s.contextAnalyzer.getConfidence(), s.distributionAnalyzer.getConfidence())
}
