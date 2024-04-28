package chardet

// fmt: off
var HZ_CLS = []int{
	1, 0, 0, 0, 0, 0, 0, 0, // 00 - 07
	0, 0, 0, 0, 0, 0, 0, 0, // 08 - 0f
	0, 0, 0, 0, 0, 0, 0, 0, // 10 - 17
	0, 0, 0, 1, 0, 0, 0, 0, // 18 - 1f
	0, 0, 0, 0, 0, 0, 0, 0, // 20 - 27
	0, 0, 0, 0, 0, 0, 0, 0, // 28 - 2f
	0, 0, 0, 0, 0, 0, 0, 0, // 30 - 37
	0, 0, 0, 0, 0, 0, 0, 0, // 38 - 3f
	0, 0, 0, 0, 0, 0, 0, 0, // 40 - 47
	0, 0, 0, 0, 0, 0, 0, 0, // 48 - 4f
	0, 0, 0, 0, 0, 0, 0, 0, // 50 - 57
	0, 0, 0, 0, 0, 0, 0, 0, // 58 - 5f
	0, 0, 0, 0, 0, 0, 0, 0, // 60 - 67
	0, 0, 0, 0, 0, 0, 0, 0, // 68 - 6f
	0, 0, 0, 0, 0, 0, 0, 0, // 70 - 77
	0, 0, 0, 4, 0, 5, 2, 0, // 78 - 7f
	1, 1, 1, 1, 1, 1, 1, 1, // 80 - 87
	1, 1, 1, 1, 1, 1, 1, 1, // 88 - 8f
	1, 1, 1, 1, 1, 1, 1, 1, // 90 - 97
	1, 1, 1, 1, 1, 1, 1, 1, // 98 - 9f
	1, 1, 1, 1, 1, 1, 1, 1, // a0 - a7
	1, 1, 1, 1, 1, 1, 1, 1, // a8 - af
	1, 1, 1, 1, 1, 1, 1, 1, // b0 - b7
	1, 1, 1, 1, 1, 1, 1, 1, // b8 - bf
	1, 1, 1, 1, 1, 1, 1, 1, // c0 - c7
	1, 1, 1, 1, 1, 1, 1, 1, // c8 - cf
	1, 1, 1, 1, 1, 1, 1, 1, // d0 - d7
	1, 1, 1, 1, 1, 1, 1, 1, // d8 - df
	1, 1, 1, 1, 1, 1, 1, 1, // e0 - e7
	1, 1, 1, 1, 1, 1, 1, 1, // e8 - ef
	1, 1, 1, 1, 1, 1, 1, 1, // f0 - f7
	1, 1, 1, 1, 1, 1, 1, 1, // f8 - ff
}

var HZ_ST = []int{
	START, ERROR, 3, START, START, START, ERROR, ERROR, // 00-07
	ERROR, ERROR, ERROR, ERROR, ITS_ME, ITS_ME, ITS_ME, ITS_ME, // 08-0f
	ITS_ME, ITS_ME, ERROR, ERROR, START, START, 4, ERROR, // 10-17
	5, ERROR, 6, ERROR, 5, 5, 4, ERROR, // 18-1f
	4, ERROR, 4, 4, 4, ERROR, 4, ERROR, // 20-27
	4, ITS_ME, START, START, START, START, START, START, // 28-2f
}

// fmt: on

var HZ_CHAR_LEN_TABLE = []int{0, 0, 0, 0, 0, 0}

var HZ_SM_MODEL = CodingStateMachineDict{
	classTable:   HZ_CLS,
	classFactor:  6,
	stateTable:   HZ_ST,
	charLenTable: HZ_CHAR_LEN_TABLE,
	name:         "HZ-GB-2312",
	language:     "Chinese",
}

// fmt: off
var ISO2022CN_CLS = []int{
	2, 0, 0, 0, 0, 0, 0, 0, // 00 - 07
	0, 0, 0, 0, 0, 0, 0, 0, // 08 - 0f
	0, 0, 0, 0, 0, 0, 0, 0, // 10 - 17
	0, 0, 0, 1, 0, 0, 0, 0, // 18 - 1f
	0, 0, 0, 0, 0, 0, 0, 0, // 20 - 27
	0, 3, 0, 0, 0, 0, 0, 0, // 28 - 2f
	0, 0, 0, 0, 0, 0, 0, 0, // 30 - 37
	0, 0, 0, 0, 0, 0, 0, 0, // 38 - 3f
	0, 0, 0, 4, 0, 0, 0, 0, // 40 - 47
	0, 0, 0, 0, 0, 0, 0, 0, // 48 - 4f
	0, 0, 0, 0, 0, 0, 0, 0, // 50 - 57
	0, 0, 0, 0, 0, 0, 0, 0, // 58 - 5f
	0, 0, 0, 0, 0, 0, 0, 0, // 60 - 67
	0, 0, 0, 0, 0, 0, 0, 0, // 68 - 6f
	0, 0, 0, 0, 0, 0, 0, 0, // 70 - 77
	0, 0, 0, 0, 0, 0, 0, 0, // 78 - 7f
	2, 2, 2, 2, 2, 2, 2, 2, // 80 - 87
	2, 2, 2, 2, 2, 2, 2, 2, // 88 - 8f
	2, 2, 2, 2, 2, 2, 2, 2, // 90 - 97
	2, 2, 2, 2, 2, 2, 2, 2, // 98 - 9f
	2, 2, 2, 2, 2, 2, 2, 2, // a0 - a7
	2, 2, 2, 2, 2, 2, 2, 2, // a8 - af
	2, 2, 2, 2, 2, 2, 2, 2, // b0 - b7
	2, 2, 2, 2, 2, 2, 2, 2, // b8 - bf
	2, 2, 2, 2, 2, 2, 2, 2, // c0 - c7
	2, 2, 2, 2, 2, 2, 2, 2, // c8 - cf
	2, 2, 2, 2, 2, 2, 2, 2, // d0 - d7
	2, 2, 2, 2, 2, 2, 2, 2, // d8 - df
	2, 2, 2, 2, 2, 2, 2, 2, // e0 - e7
	2, 2, 2, 2, 2, 2, 2, 2, // e8 - ef
	2, 2, 2, 2, 2, 2, 2, 2, // f0 - f7
	2, 2, 2, 2, 2, 2, 2, 2, // f8 - ff
}

var ISO2022CN_ST = []int{
	START, 3, ERROR, START, START, START, START, START, // 00-07
	START, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, // 08-0f
	ERROR, ERROR, ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, // 10-17
	ITS_ME, ITS_ME, ITS_ME, ERROR, ERROR, ERROR, 4, ERROR, // 18-1f
	ERROR, ERROR, ERROR, ITS_ME, ERROR, ERROR, ERROR, ERROR, // 20-27
	5, 6, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, // 28-2f
	ERROR, ERROR, ERROR, ITS_ME, ERROR, ERROR, ERROR, ERROR, // 30-37
	ERROR, ERROR, ERROR, ERROR, ERROR, ITS_ME, ERROR, START, // 38-3f
}

// fmt: on

var ISO2022CN_CHAR_LEN_TABLE = []int{0, 0, 0, 0, 0, 0, 0, 0, 0}

var ISO2022CN_SM_MODEL = CodingStateMachineDict{
	classTable:   ISO2022CN_CLS,
	classFactor:  9,
	stateTable:   ISO2022CN_ST,
	charLenTable: ISO2022CN_CHAR_LEN_TABLE,
	name:         "ISO-2022-CN",
	language:     "Chinese",
}

// fmt: off
var ISO2022JP_CLS = []int{
	2, 0, 0, 0, 0, 0, 0, 0, // 00 - 07
	0, 0, 0, 0, 0, 0, 2, 2, // 08 - 0f
	0, 0, 0, 0, 0, 0, 0, 0, // 10 - 17
	0, 0, 0, 1, 0, 0, 0, 0, // 18 - 1f
	0, 0, 0, 0, 7, 0, 0, 0, // 20 - 27
	3, 0, 0, 0, 0, 0, 0, 0, // 28 - 2f
	0, 0, 0, 0, 0, 0, 0, 0, // 30 - 37
	0, 0, 0, 0, 0, 0, 0, 0, // 38 - 3f
	6, 0, 4, 0, 8, 0, 0, 0, // 40 - 47
	0, 9, 5, 0, 0, 0, 0, 0, // 48 - 4f
	0, 0, 0, 0, 0, 0, 0, 0, // 50 - 57
	0, 0, 0, 0, 0, 0, 0, 0, // 58 - 5f
	0, 0, 0, 0, 0, 0, 0, 0, // 60 - 67
	0, 0, 0, 0, 0, 0, 0, 0, // 68 - 6f
	0, 0, 0, 0, 0, 0, 0, 0, // 70 - 77
	0, 0, 0, 0, 0, 0, 0, 0, // 78 - 7f
	2, 2, 2, 2, 2, 2, 2, 2, // 80 - 87
	2, 2, 2, 2, 2, 2, 2, 2, // 88 - 8f
	2, 2, 2, 2, 2, 2, 2, 2, // 90 - 97
	2, 2, 2, 2, 2, 2, 2, 2, // 98 - 9f
	2, 2, 2, 2, 2, 2, 2, 2, // a0 - a7
	2, 2, 2, 2, 2, 2, 2, 2, // a8 - af
	2, 2, 2, 2, 2, 2, 2, 2, // b0 - b7
	2, 2, 2, 2, 2, 2, 2, 2, // b8 - bf
	2, 2, 2, 2, 2, 2, 2, 2, // c0 - c7
	2, 2, 2, 2, 2, 2, 2, 2, // c8 - cf
	2, 2, 2, 2, 2, 2, 2, 2, // d0 - d7
	2, 2, 2, 2, 2, 2, 2, 2, // d8 - df
	2, 2, 2, 2, 2, 2, 2, 2, // e0 - e7
	2, 2, 2, 2, 2, 2, 2, 2, // e8 - ef
	2, 2, 2, 2, 2, 2, 2, 2, // f0 - f7
	2, 2, 2, 2, 2, 2, 2, 2, // f8 - ff
}

var ISO2022JP_ST = []int{
	START, 3, ERROR, START, START, START, START, START, // 00-07
	START, START, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, // 08-0f
	ERROR, ERROR, ERROR, ERROR, ITS_ME, ITS_ME, ITS_ME, ITS_ME, // 10-17
	ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, ERROR, ERROR, // 18-1f
	ERROR, 5, ERROR, ERROR, ERROR, 4, ERROR, ERROR, // 20-27
	ERROR, ERROR, ERROR, 6, ITS_ME, ERROR, ITS_ME, ERROR, // 28-2f
	ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ITS_ME, ITS_ME, // 30-37
	ERROR, ERROR, ERROR, ITS_ME, ERROR, ERROR, ERROR, ERROR, // 38-3f
	ERROR, ERROR, ERROR, ERROR, ITS_ME, ERROR, START, START, // 40-47
}

// fmt: on

var ISO2022JP_CHAR_LEN_TABLE = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

var ISO2022JP_SM_MODEL = CodingStateMachineDict{
	classTable:   ISO2022JP_CLS,
	classFactor:  10,
	stateTable:   ISO2022JP_ST,
	charLenTable: ISO2022JP_CHAR_LEN_TABLE,
	name:         "ISO-2022-JP",
	language:     "Japanese",
}

// fmt: off
var ISO2022KR_CLS = []int{
	2, 0, 0, 0, 0, 0, 0, 0, // 00 - 07
	0, 0, 0, 0, 0, 0, 0, 0, // 08 - 0f
	0, 0, 0, 0, 0, 0, 0, 0, // 10 - 17
	0, 0, 0, 1, 0, 0, 0, 0, // 18 - 1f
	0, 0, 0, 0, 3, 0, 0, 0, // 20 - 27
	0, 4, 0, 0, 0, 0, 0, 0, // 28 - 2f
	0, 0, 0, 0, 0, 0, 0, 0, // 30 - 37
	0, 0, 0, 0, 0, 0, 0, 0, // 38 - 3f
	0, 0, 0, 5, 0, 0, 0, 0, // 40 - 47
	0, 0, 0, 0, 0, 0, 0, 0, // 48 - 4f
	0, 0, 0, 0, 0, 0, 0, 0, // 50 - 57
	0, 0, 0, 0, 0, 0, 0, 0, // 58 - 5f
	0, 0, 0, 0, 0, 0, 0, 0, // 60 - 67
	0, 0, 0, 0, 0, 0, 0, 0, // 68 - 6f
	0, 0, 0, 0, 0, 0, 0, 0, // 70 - 77
	0, 0, 0, 0, 0, 0, 0, 0, // 78 - 7f
	2, 2, 2, 2, 2, 2, 2, 2, // 80 - 87
	2, 2, 2, 2, 2, 2, 2, 2, // 88 - 8f
	2, 2, 2, 2, 2, 2, 2, 2, // 90 - 97
	2, 2, 2, 2, 2, 2, 2, 2, // 98 - 9f
	2, 2, 2, 2, 2, 2, 2, 2, // a0 - a7
	2, 2, 2, 2, 2, 2, 2, 2, // a8 - af
	2, 2, 2, 2, 2, 2, 2, 2, // b0 - b7
	2, 2, 2, 2, 2, 2, 2, 2, // b8 - bf
	2, 2, 2, 2, 2, 2, 2, 2, // c0 - c7
	2, 2, 2, 2, 2, 2, 2, 2, // c8 - cf
	2, 2, 2, 2, 2, 2, 2, 2, // d0 - d7
	2, 2, 2, 2, 2, 2, 2, 2, // d8 - df
	2, 2, 2, 2, 2, 2, 2, 2, // e0 - e7
	2, 2, 2, 2, 2, 2, 2, 2, // e8 - ef
	2, 2, 2, 2, 2, 2, 2, 2, // f0 - f7
	2, 2, 2, 2, 2, 2, 2, 2, // f8 - ff
}

var ISO2022KR_ST = []int{
	START, 3, ERROR, START, START, START, ERROR, ERROR, // 00-07
	ERROR, ERROR, ERROR, ERROR, ITS_ME, ITS_ME, ITS_ME, ITS_ME, // 08-0f
	ITS_ME, ITS_ME, ERROR, ERROR, ERROR, 4, ERROR, ERROR, // 10-17
	ERROR, ERROR, ERROR, ERROR, 5, ERROR, ERROR, ERROR, // 18-1f
	ERROR, ERROR, ERROR, ITS_ME, START, START, START, START, // 20-27
}

// fmt: on

var ISO2022KR_CHAR_LEN_TABLE = []int{0, 0, 0, 0, 0, 0}

var ISO2022KR_SM_MODEL = CodingStateMachineDict{
	classTable:   ISO2022KR_CLS,
	classFactor:  6,
	stateTable:   ISO2022KR_ST,
	charLenTable: ISO2022KR_CHAR_LEN_TABLE,
	name:         "ISO-2022-KR",
	language:     "Korean",
}
