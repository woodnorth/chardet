package chardet

type InputState int

const (
	PURE_ASCII InputState = iota
	ESC_ASCII
	HIGH_BYTE
)

type LanguageFilter int

const (
	NONE                LanguageFilter = 0x00
	CHINESE_SIMPLIFIED  LanguageFilter = 0x01
	CHINESE_TRADITIONAL LanguageFilter = 0x02
	JAPANESE            LanguageFilter = 0x04
	KOREAN              LanguageFilter = 0x08
	NON_CJK             LanguageFilter = 0x10
	ALL                 LanguageFilter = 0x1F
	CHINESE                            = CHINESE_SIMPLIFIED | CHINESE_TRADITIONAL
	CJK                                = CHINESE | JAPANESE | KOREAN
)

type ProbingState int

const (
	DETECTING ProbingState = iota
	FOUND_IT
	NOT_ME
)

type MachineState = int

const (
	START MachineState = iota
	ERROR
	ITS_ME
)

type SequenceLikelihood int

const (
	NEGATIVE SequenceLikelihood = iota
	UNLIKELY                    = 1
	LIKELY                      = 2
	POSITIVE                    = 3
)

type CharacterCategory = int

const (
	UNDEFINED  CharacterCategory = 255
	LINE_BREAK CharacterCategory = 254
	SYMBOL     CharacterCategory = 253
	DIGIT      CharacterCategory = 252
	CONTROL    CharacterCategory = 251
)
