package chardet

import "math"

const (
	ONE_CHAR_PROB = 0.5
)

type UTF8Prober struct {
	CommonCharsetProber
	codingSm   *CodingStateMachine
	numMbChars int
}

func NewUTF8Prober() *UTF8Prober {
	var u UTF8Prober
	u.init(NONE)
	u.numMbChars = 0
	u.codingSm = NewCodingStateMachine(UTF8_SM_MODEL)
	u.reset()
	return &u
}

func (u *UTF8Prober) reset() {
	u.CommonCharsetProber.reset()
	u.codingSm.reset()
	u.numMbChars = 0
}

func (u *UTF8Prober) charName() string {
	return "utf-8"
}

func (u *UTF8Prober) language() string {
	return ""
}

func (u *UTF8Prober) feed(data []byte) ProbingState {
	for _, c := range data {
		codingState := u.codingSm.nextState(int(c))
		if codingState == ERROR {
			u.state_ = NOT_ME
			break
		}
		if codingState == ITS_ME {
			u.state_ = FOUND_IT
			break
		}
		if codingState == START {
			if u.codingSm.getCurrentCharLen() >= 2 {
				u.numMbChars += 1
			}
		}
	}

	if u.state_ == DETECTING {
		if u.getConfidence() > SHORTCUT_THRESHOLD {
			u.state_ = FOUND_IT
		}
	}
	return u.state_
}

func (u *UTF8Prober) getConfidence() float64 {
	unlike := 0.99
	if u.numMbChars < 6 {
		unlike *= math.Pow(ONE_CHAR_PROB, float64(u.numMbChars))
		return 1 - unlike
	}
	return unlike
}
