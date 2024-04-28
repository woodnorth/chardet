package chardet

type SBCSGroupProber struct {
	CharSetGroupProber
}

func NewSBCSGroupProber() *SBCSGroupProber {
	var g SBCSGroupProber
	g.CharSetGroupProber.init(NONE)
	hebrewProber := NewHebrewProber()

	logicalHeBrewProber := NewSingleByteCharSetProber(WINDOWS_1255_HEBREW_MODEL, false, hebrewProber)
	visualHeBrewProber := NewSingleByteCharSetProber(WINDOWS_1255_HEBREW_MODEL, true, hebrewProber)
	hebrewProber.setModelProbers(logicalHeBrewProber, visualHeBrewProber)

	g.probers = append(g.probers, NewSingleByteCharSetProber(WINDOWS_1251_RUSSIAN_MODEL, false, nil))
	g.probers = append(g.probers, NewSingleByteCharSetProber(KOI8_R_RUSSIAN_MODEL, false, nil))
	g.probers = append(g.probers, NewSingleByteCharSetProber(ISO_8859_5_RUSSIAN_MODEL, false, nil))
	g.probers = append(g.probers, NewSingleByteCharSetProber(MACCYRILLIC_RUSSIAN_MODEL, false, nil))
	g.probers = append(g.probers, NewSingleByteCharSetProber(IBM866_RUSSIAN_MODEL, false, nil))
	g.probers = append(g.probers, NewSingleByteCharSetProber(IBM855_RUSSIAN_MODEL, false, nil))
	g.probers = append(g.probers, NewSingleByteCharSetProber(ISO_8859_7_GREEK_MODEL, false, nil))
	g.probers = append(g.probers, NewSingleByteCharSetProber(WINDOWS_1253_GREEK_MODEL, false, nil))
	g.probers = append(g.probers, NewSingleByteCharSetProber(ISO_8859_5_BULGARIAN_MODEL, false, nil))
	g.probers = append(g.probers, NewSingleByteCharSetProber(WINDOWS_1251_BULGARIAN_MODEL, false, nil))

	g.probers = append(g.probers, NewSingleByteCharSetProber(TIS_620_THAI_MODEL, false, nil))
	g.probers = append(g.probers, NewSingleByteCharSetProber(ISO_8859_9_TURKISH_MODEL, false, nil))

	g.probers = append(g.probers, hebrewProber)
	g.probers = append(g.probers, logicalHeBrewProber)
	g.probers = append(g.probers, visualHeBrewProber)
	g.reset()
	return &g
}
