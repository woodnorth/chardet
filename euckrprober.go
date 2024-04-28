package chardet

type EUCKRProber struct {
	MultiByteCharSetProber
}

func NewEUCKRProber() *EUCKRProber {
	var e EUCKRProber
	e.codingSm = NewCodingStateMachine(EUCKR_SM_MODEL)
	e.distributionAnalyzer = NewEUCKRDistributionAnalysis()
	e.reset()
	return &e
}

func (e *EUCKRProber) charName() string {
	return "EUC-KR"
}

func (e *EUCKRProber) language() string {
	return "Korean"
}
