package chardet

type CodingStateMachine struct {
	model      CodingStateMachineDict
	curBytePos int
	curCharLen int
	curState   MachineState
	active     bool
}

func NewCodingStateMachine(sm CodingStateMachineDict) *CodingStateMachine {
	c := CodingStateMachine{
		model:      sm,
		curState:   START,
		curCharLen: 0,
		curBytePos: 0,
		active:     true,
	}
	c.reset()
	return &c
}

func (m *CodingStateMachine) reset() {
	m.curState = START
}

func (m *CodingStateMachine) nextState(c int) MachineState {
	byteClass := m.model.classTable[c]
	if m.curState == START {
		m.curBytePos = 0
		m.curCharLen = m.model.charLenTable[byteClass]
	}
	curState := m.curState*m.model.classFactor + byteClass
	m.curState = m.model.stateTable[curState]
	m.curBytePos += 1
	return m.curState
}

func (m *CodingStateMachine) getCurrentCharLen() int {
	return m.curCharLen
}

func (m *CodingStateMachine) getCodingStateMachine() string {
	return m.model.name
}

func (m *CodingStateMachine) getLanguage() string {
	return m.model.language
}
