package chardet

type GB2312Prober struct {
	MultiByteCharSetProber
}

func NewGB2312Prober() *GB2312Prober {
	var g GB2312Prober
	g.init(NONE)
	g.codingSm = NewCodingStateMachine(GB2312_SM_MODEL)
	g.distributionAnalyzer = NewGB2312DistributionAnalysis()
	g.reset()
	return &g
}

func (g *GB2312Prober) charName() string {
	return "GB2312"
}

func (g *GB2312Prober) language() string {
	return "Chinese"
}
