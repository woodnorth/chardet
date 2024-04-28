package chardet

type JOHABProber struct {
	MultiByteCharSetProber
}

func NewJOHABProber() *JOHABProber {
	var e JOHABProber
	e.codingSm = NewCodingStateMachine(JOHAB_SM_MODEL)
	e.distributionAnalyzer = NewJOHABDistributionAnalysis()
	e.reset()
	return &e
}

func (e *JOHABProber) charName() string {
	return "Johab"
}

func (e *JOHABProber) language() string {
	return "Korean"
}
