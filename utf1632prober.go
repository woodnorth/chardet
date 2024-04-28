package chardet

const (
	MIN_CHARS_FOR_DETECTION = 20
	EXPECTED_RATIO          = 0.94
)

type UTF1632Prober struct {
	CommonCharsetProber
	position                           int
	zerosAtMod                         []int
	nonZerosAtMod                      []int
	quad                               []int
	invalidUtf16be                     bool
	invalidUtf16le                     bool
	invalidUtf32be                     bool
	invalidUtf32le                     bool
	firstHalfSurrogatePairDetected16be bool
	firstHalfSurrogatePairDetected16le bool
}

func NewUTF1632Prober() UTF1632Prober {
	var u UTF1632Prober
	u.init(NONE)
	u.zerosAtMod = make([]int, 4)
	u.nonZerosAtMod = make([]int, 4)
	u.state_ = DETECTING
	u.quad = make([]int, 4)
	u.reset()
	return u
}

func (u *UTF1632Prober) approx32bitChars() float64 {
	return max(1.0, float64(u.position/4.0))
}

func (u *UTF1632Prober) approx16bitChars() float64 {
	return max(1.0, float64(u.position/2.0))
}

func (u *UTF1632Prober) isLikelyUtf32be() bool {
	approxChars := u.approx32bitChars()
	return approxChars >= MIN_CHARS_FOR_DETECTION && (float64(u.zerosAtMod[0])/approxChars > EXPECTED_RATIO &&
		float64(u.zerosAtMod[1])/approxChars > EXPECTED_RATIO &&
		float64(u.zerosAtMod[2])/approxChars > EXPECTED_RATIO &&
		float64(u.nonZerosAtMod[3])/approxChars > EXPECTED_RATIO &&
		!u.invalidUtf32be)
}

func (u *UTF1632Prober) isLikelyUtf32le() bool {
	approxChars := u.approx32bitChars()
	return approxChars >= MIN_CHARS_FOR_DETECTION && (float64(u.nonZerosAtMod[0])/approxChars > EXPECTED_RATIO &&
		float64(u.zerosAtMod[1])/approxChars > EXPECTED_RATIO &&
		float64(u.zerosAtMod[2])/approxChars > EXPECTED_RATIO &&
		float64(u.zerosAtMod[3])/approxChars > EXPECTED_RATIO &&
		!u.invalidUtf32le)
}

func (u *UTF1632Prober) isLikelyUtf16be() bool {
	approxChars := u.approx16bitChars()
	return approxChars >= MIN_CHARS_FOR_DETECTION && (float64(u.nonZerosAtMod[1]+u.nonZerosAtMod[3])/approxChars > EXPECTED_RATIO &&
		float64(u.zerosAtMod[0]+u.zerosAtMod[2])/approxChars > EXPECTED_RATIO &&
		!u.invalidUtf16be)
}

func (u *UTF1632Prober) isLikelyUtf16le() bool {
	approxChars := u.approx16bitChars()
	return approxChars >= MIN_CHARS_FOR_DETECTION && (float64(u.nonZerosAtMod[0]+u.nonZerosAtMod[2])/approxChars > EXPECTED_RATIO &&
		float64(u.zerosAtMod[1]+u.zerosAtMod[3])/approxChars > EXPECTED_RATIO &&
		!u.invalidUtf16le)
}

func (u *UTF1632Prober) charName() string {
	if u.isLikelyUtf32be() {
		return "utf-32be"
	}
	if u.isLikelyUtf32le() {
		return "utf-32le"
	}
	if u.isLikelyUtf16be() {
		return "utf-16be"
	}
	if u.isLikelyUtf16le() {
		return "utf-16le"
	}
	return "utf-16"
}

func (u *UTF1632Prober) language() string {
	return ""
}

func (u *UTF1632Prober) validateUtf32Characters(quad []int) {
	if quad[0] != 0 || quad[1] > 0x01 || (quad[0] == 0 && quad[1] == 0 && quad[2] <= 0xDF && quad[2] >= 0xD8) {
		u.invalidUtf32be = true
	}
	if quad[3] != 0 || quad[2] > 0x10 || (quad[3] == 0 && quad[2] == 0 && quad[1] <= 0xDF && quad[1] >= 0xD8) {
		u.invalidUtf32le = true
	}
}

func (u *UTF1632Prober) validateUtf16Characters(pair []int) {
	if !u.firstHalfSurrogatePairDetected16be {
		if pair[0] >= 0xD8 && pair[0] <= 0xDB {
			u.firstHalfSurrogatePairDetected16be = true
		} else if pair[0] >= 0xDC && pair[0] <= 0xDF {
			u.invalidUtf16be = true
		}
	} else {
		if pair[0] >= 0xDC && pair[0] <= 0xDF {
			u.firstHalfSurrogatePairDetected16be = false
		} else {
			u.invalidUtf16be = true
		}
	}

	if !u.firstHalfSurrogatePairDetected16le {
		if pair[1] >= 0xD8 && pair[1] <= 0xDB {
			u.firstHalfSurrogatePairDetected16le = true
		} else if pair[1] >= 0xDC && pair[1] <= 0xDF {
			u.invalidUtf16le = true
		}
	} else {
		if pair[1] >= 0xDC && pair[1] <= 0xDF {
			u.firstHalfSurrogatePairDetected16le = false
		} else {
			u.invalidUtf16le = true
		}
	}
}

func (u *UTF1632Prober) feed(data []byte) ProbingState {
	for _, c := range data {
		mod4 := u.position % 4
		u.quad[mod4] = int(c)
		if mod4 == 3 {
			u.validateUtf32Characters(u.quad)
			u.validateUtf16Characters(u.quad[0:2])
			u.validateUtf16Characters(u.quad[2:4])
		}
		if c == 0 {
			u.zerosAtMod[mod4] += 1
		} else {
			u.nonZerosAtMod[mod4] += 1
		}
		u.position += 1
	}
	u.state()
	return u.state_
}

func (u *UTF1632Prober) state() ProbingState {
	if u.state_ == NOT_ME || u.state_ == FOUND_IT {
		return u.state_
	}
	if u.getConfidence() > 0.8 {
		u.state_ = FOUND_IT
	} else if u.position > 4*1024 {
		u.state_ = NOT_ME
	}
	return u.state_
}

func (u *UTF1632Prober) getConfidence() float64 {
	if u.isLikelyUtf16le() || u.isLikelyUtf16be() || u.isLikelyUtf32le() || u.isLikelyUtf32be() {
		return 0.85
	}
	return 0
}
