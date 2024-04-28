package chardet

type CodingStateMachineDict struct {
	classTable   []int
	classFactor  int
	stateTable   []int
	charLenTable []int
	name         string
	language     string
}
