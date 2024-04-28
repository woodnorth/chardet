package chardet

type ResultDict struct {
	encoding   string
	confidence float64
	language   string
}

func (r *ResultDict) GetEncoding() string {
	return r.encoding
}

func (r *ResultDict) GetConfidence() float64 {
	return r.confidence
}

func (r *ResultDict) GetLanguage() string {
	return r.language
}
