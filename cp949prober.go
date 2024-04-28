package chardet

type CP949Prober struct {
	MultiByteCharSetProber
}

func NewCP949Prober() *CP949Prober {
	var e CP949Prober
	e.codingSm = NewCodingStateMachine(CP949_SM_MODEL)
	e.distributionAnalyzer = NewEUCKRDistributionAnalysis()
	e.reset()
	return &e
}

func (e *CP949Prober) charName() string {
	return "CP949"
}

func (e *CP949Prober) language() string {
	return "Korean"
}
