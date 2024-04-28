package chardet

//type ResultDict map[string]interface{}
//

type ResultDict struct {
	encoding   string
	confidence float64
	language   string
}
