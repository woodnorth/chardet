package chardet

type Big5Prober struct {
	MultiByteCharSetProber
}

func NewBig5Prober() *Big5Prober {
	var g Big5Prober
	g.init(NONE)
	g.codingSm = NewCodingStateMachine(BIG5_SM_MODEL)
	g.distributionAnalyzer = NewBig5DistributionAnalysis()
	g.reset()
	return &g
}

func (b *Big5Prober) charName() string {
	return "Big5"
}

func (b *Big5Prober) language() string {
	return "Chinese"
}
