package chardet

type EUCTWProber struct {
	MultiByteCharSetProber
}

func NewEUCTWProber() *EUCTWProber {
	var e EUCTWProber
	e.codingSm = NewCodingStateMachine(EUCTW_SM_MODEL)
	e.distributionAnalyzer = NewEUCTWDistributionAnalysis()
	e.reset()
	return &e
}

func (e *EUCTWProber) charName() string {
	return "EUC-TW"
}

func (e *EUCTWProber) language() string {
	return "Taiwan"
}
