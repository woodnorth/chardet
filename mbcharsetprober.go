package chardet

type MultiByteCharSetProber struct {
	CommonCharsetProber
	codingSm             *CodingStateMachine
	distributionAnalyzer LanguageAnalysis
	lastChar             []byte
}

func (m *MultiByteCharSetProber) init(filter LanguageFilter) {
	m.CommonCharsetProber.init(filter)
	m.lastChar = make([]byte, 2)
}

func (m *MultiByteCharSetProber) reset() {
	m.CommonCharsetProber.reset()
	if m.codingSm != nil {
		m.codingSm.reset()
	}
	if m.distributionAnalyzer != nil {
		m.distributionAnalyzer.reset()
	}
	m.lastChar = []byte("\x00\x00")
}

func (m *MultiByteCharSetProber) feed(data []byte) ProbingState {
	for i, b := range data {
		codingState := m.codingSm.nextState(int(b))
		if codingState == ERROR {
			m.state_ = NOT_ME
			break
		}
		if codingState == ITS_ME {
			m.state_ = FOUND_IT
			break
		}
		if codingState == START {
			charLen := m.codingSm.getCurrentCharLen()
			if i == 0 {
				m.lastChar[1] = b
				m.distributionAnalyzer.feed(m.lastChar, charLen)
			} else {
				m.distributionAnalyzer.feed(data[i-1:i+1], charLen)
			}
		}
	}
	m.lastChar[0] = data[len(data)-1]

	if m.state_ == DETECTING {
		if m.distributionAnalyzer.getEnoughData() && m.getConfidence() > SHORTCUT_THRESHOLD {
			m.state_ = FOUND_IT
		}
	}
	return m.state_
}

func (m *MultiByteCharSetProber) getConfidence() float64 {
	return m.distributionAnalyzer.getConfidence()
}
