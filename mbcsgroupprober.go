package chardet

type MBCSGroupProber struct {
	CharSetGroupProber
}

func NewMBCSGroupProber(filter LanguageFilter) *MBCSGroupProber {
	var m MBCSGroupProber
	m.init(filter)
	m.probers = append(m.probers, NewUTF8Prober())
	m.probers = append(m.probers, NewSJISProber())
	m.probers = append(m.probers, NewEUCJPProber())
	m.probers = append(m.probers, NewGB2312Prober())
	m.probers = append(m.probers, NewEUCKRProber())
	m.probers = append(m.probers, NewCP949Prober())
	m.probers = append(m.probers, NewBig5Prober())
	m.probers = append(m.probers, NewEUCTWProber())
	m.probers = append(m.probers, NewJOHABProber())
	m.CharSetGroupProber.reset()
	return &m
}
