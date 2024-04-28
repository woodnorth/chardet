package chardet

type EUCJPProber struct {
	MultiByteCharSetProber
	contextAnalyzer *EUCJPContextAnalysis
}

func NewEUCJPProber() *EUCJPProber {
	var s EUCJPProber
	s.MultiByteCharSetProber.init(NONE)
	s.codingSm = NewCodingStateMachine(EUCJP_SM_MODEL)
	s.distributionAnalyzer = NewEUCJPDistributionAnalysis()
	s.contextAnalyzer = NewEUCJPContextAnalysis()
	s.reset()
	return &s
}

func (s *EUCJPProber) reset() {
	s.MultiByteCharSetProber.reset()
	s.contextAnalyzer.reset()
}

func (s *EUCJPProber) charName() string {
	return "EUC-JP"
}

func (s *EUCJPProber) language() string {
	return "Japanese"
}

func (s *EUCJPProber) feed(data []byte) ProbingState {
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
				s.contextAnalyzer.feed(s.lastChar, charLen)
				s.distributionAnalyzer.feed(s.lastChar, charLen)
			} else {
				s.contextAnalyzer.feed(data[i-1:i+1], charLen)
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

func (s *EUCJPProber) getConfidence() float64 {
	return max(s.contextAnalyzer.getConfidence(), s.distributionAnalyzer.getConfidence())
}
