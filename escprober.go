package chardet

type EscCharSetProber struct {
	CommonCharsetProber
	codingSm         []*CodingStateMachine
	activeSmCount    int
	detectedCharset  string
	detectedLanguage string
	state            ProbingState
}

func NewEscCharSetProber(languageFilter LanguageFilter) EscCharSetProber {
	var e EscCharSetProber
	e.init(languageFilter)
	e.codingSm = make([]*CodingStateMachine, 0)
	e.langFilter = languageFilter
	if languageFilter&CHINESE_SIMPLIFIED != 0 {
		e.codingSm = append(e.codingSm, NewCodingStateMachine(HZ_SM_MODEL))
		e.codingSm = append(e.codingSm, NewCodingStateMachine(ISO2022CN_SM_MODEL))
	}
	if languageFilter&JAPANESE != 0 {
		e.codingSm = append(e.codingSm, NewCodingStateMachine(ISO2022JP_SM_MODEL))
	}
	if languageFilter&KOREAN != 0 {
		e.codingSm = append(e.codingSm, NewCodingStateMachine(ISO2022KR_SM_MODEL))
	}
	e.activeSmCount = 0
	e.state = DETECTING
	e.reset()
	return e
}

func (e *EscCharSetProber) reset() {
	e.CommonCharsetProber.reset()
	for _, sm := range e.codingSm {
		sm.active = true
		sm.reset()
	}
	e.activeSmCount = len(e.codingSm)
	e.detectedCharset = ""
	e.detectedLanguage = ""
}

func (e *EscCharSetProber) charName() string {
	return e.detectedCharset
}

func (e *EscCharSetProber) language() string {
	return e.detectedLanguage
}

func (e *EscCharSetProber) getConfidence() float64 {
	if e.detectedCharset != "" {
		return 0.99
	}
	return 0
}

func (e *EscCharSetProber) feed(data []byte) ProbingState {
	for _, c := range data {
		for _, sm := range e.codingSm {
			if !sm.active {
				continue
			}
			codingState := sm.nextState(int(c))
			if codingState == ERROR {
				sm.active = false
				e.activeSmCount -= 1
				if e.activeSmCount <= 0 {
					e.state = NOT_ME
					return e.state
				}
			} else if codingState == ITS_ME {
				e.state = FOUND_IT
				e.detectedCharset = sm.getCodingStateMachine()
				e.detectedLanguage = sm.getLanguage()
				return e.state
			}
		}
	}
	return e.state
}
