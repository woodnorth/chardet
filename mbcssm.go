package chardet

// fmt: off
var BIG5_CLS = []int{
	1, 1, 1, 1, 1, 1, 1, 1, // 00 - 07    //allow 0x00 as legal value
	1, 1, 1, 1, 1, 1, 0, 0, // 08 - 0f
	1, 1, 1, 1, 1, 1, 1, 1, // 10 - 17
	1, 1, 1, 0, 1, 1, 1, 1, // 18 - 1f
	1, 1, 1, 1, 1, 1, 1, 1, // 20 - 27
	1, 1, 1, 1, 1, 1, 1, 1, // 28 - 2f
	1, 1, 1, 1, 1, 1, 1, 1, // 30 - 37
	1, 1, 1, 1, 1, 1, 1, 1, // 38 - 3f
	2, 2, 2, 2, 2, 2, 2, 2, // 40 - 47
	2, 2, 2, 2, 2, 2, 2, 2, // 48 - 4f
	2, 2, 2, 2, 2, 2, 2, 2, // 50 - 57
	2, 2, 2, 2, 2, 2, 2, 2, // 58 - 5f
	2, 2, 2, 2, 2, 2, 2, 2, // 60 - 67
	2, 2, 2, 2, 2, 2, 2, 2, // 68 - 6f
	2, 2, 2, 2, 2, 2, 2, 2, // 70 - 77
	2, 2, 2, 2, 2, 2, 2, 1, // 78 - 7f
	4, 4, 4, 4, 4, 4, 4, 4, // 80 - 87
	4, 4, 4, 4, 4, 4, 4, 4, // 88 - 8f
	4, 4, 4, 4, 4, 4, 4, 4, // 90 - 97
	4, 4, 4, 4, 4, 4, 4, 4, // 98 - 9f
	4, 3, 3, 3, 3, 3, 3, 3, // a0 - a7
	3, 3, 3, 3, 3, 3, 3, 3, // a8 - af
	3, 3, 3, 3, 3, 3, 3, 3, // b0 - b7
	3, 3, 3, 3, 3, 3, 3, 3, // b8 - bf
	3, 3, 3, 3, 3, 3, 3, 3, // c0 - c7
	3, 3, 3, 3, 3, 3, 3, 3, // c8 - cf
	3, 3, 3, 3, 3, 3, 3, 3, // d0 - d7
	3, 3, 3, 3, 3, 3, 3, 3, // d8 - df
	3, 3, 3, 3, 3, 3, 3, 3, // e0 - e7
	3, 3, 3, 3, 3, 3, 3, 3, // e8 - ef
	3, 3, 3, 3, 3, 3, 3, 3, // f0 - f7
	3, 3, 3, 3, 3, 3, 3, 0, // f8 - ff
}

var BIG5_ST = []int{
	ERROR, START, START, 3, ERROR, ERROR, ERROR, ERROR, //00-07
	ERROR, ERROR, ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, ERROR, //08-0f
	ERROR, START, START, START, START, START, START, START, //10-17
}

// fmt: on

var BIG5_CHAR_LEN_TABLE = []int{0, 1, 1, 2, 0}

var BIG5_SM_MODEL = CodingStateMachineDict{
	classTable:   BIG5_CLS,
	classFactor:  5,
	stateTable:   BIG5_ST,
	charLenTable: BIG5_CHAR_LEN_TABLE,
	name:         "Big5",
}

var GB2312_CLS = []int{
	1, 1, 1, 1, 1, 1, 1, 1, // 00 - 07
	1, 1, 1, 1, 1, 1, 0, 0, // 08 - 0f
	1, 1, 1, 1, 1, 1, 1, 1, // 10 - 17
	1, 1, 1, 0, 1, 1, 1, 1, // 18 - 1f
	1, 1, 1, 1, 1, 1, 1, 1, // 20 - 27
	1, 1, 1, 1, 1, 1, 1, 1, // 28 - 2f
	3, 3, 3, 3, 3, 3, 3, 3, // 30 - 37
	3, 3, 1, 1, 1, 1, 1, 1, // 38 - 3f
	2, 2, 2, 2, 2, 2, 2, 2, // 40 - 47
	2, 2, 2, 2, 2, 2, 2, 2, // 48 - 4f
	2, 2, 2, 2, 2, 2, 2, 2, // 50 - 57
	2, 2, 2, 2, 2, 2, 2, 2, // 58 - 5f
	2, 2, 2, 2, 2, 2, 2, 2, // 60 - 67
	2, 2, 2, 2, 2, 2, 2, 2, // 68 - 6f
	2, 2, 2, 2, 2, 2, 2, 2, // 70 - 77
	2, 2, 2, 2, 2, 2, 2, 4, // 78 - 7f
	5, 6, 6, 6, 6, 6, 6, 6, // 80 - 87
	6, 6, 6, 6, 6, 6, 6, 6, // 88 - 8f
	6, 6, 6, 6, 6, 6, 6, 6, // 90 - 97
	6, 6, 6, 6, 6, 6, 6, 6, // 98 - 9f
	6, 6, 6, 6, 6, 6, 6, 6, // a0 - a7
	6, 6, 6, 6, 6, 6, 6, 6, // a8 - af
	6, 6, 6, 6, 6, 6, 6, 6, // b0 - b7
	6, 6, 6, 6, 6, 6, 6, 6, // b8 - bf
	6, 6, 6, 6, 6, 6, 6, 6, // c0 - c7
	6, 6, 6, 6, 6, 6, 6, 6, // c8 - cf
	6, 6, 6, 6, 6, 6, 6, 6, // d0 - d7
	6, 6, 6, 6, 6, 6, 6, 6, // d8 - df
	6, 6, 6, 6, 6, 6, 6, 6, // e0 - e7
	6, 6, 6, 6, 6, 6, 6, 6, // e8 - ef
	6, 6, 6, 6, 6, 6, 6, 6, // f0 - f7
	6, 6, 6, 6, 6, 6, 6, 0, // f8 - ff
}

var GB2312_ST = []int{
	ERROR, START, START, START, START, START, 3, ERROR, //00-07
	ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ITS_ME, ITS_ME, //08-0f
	ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, ERROR, ERROR, START, //10-17
	4, ERROR, START, START, ERROR, ERROR, ERROR, ERROR, //18-1f
	ERROR, ERROR, 5, ERROR, ERROR, ERROR, ITS_ME, ERROR, //20-27
	ERROR, ERROR, START, START, START, START, START, START, //28-2f
}

var GB2312_CHAR_LEN_TABLE = []int{0, 1, 1, 1, 1, 1, 2}

var GB2312_SM_MODEL = CodingStateMachineDict{
	classTable:   GB2312_CLS,
	classFactor:  7,
	stateTable:   GB2312_ST,
	charLenTable: GB2312_CHAR_LEN_TABLE,
	name:         "GB2312",
}

// UTF-8
// fmt: off
var UTF8_CLS = []int{
	1, 1, 1, 1, 1, 1, 1, 1, // 00 - 07  //allow 0x00 as a legal value
	1, 1, 1, 1, 1, 1, 0, 0, // 08 - 0f
	1, 1, 1, 1, 1, 1, 1, 1, // 10 - 17
	1, 1, 1, 0, 1, 1, 1, 1, // 18 - 1f
	1, 1, 1, 1, 1, 1, 1, 1, // 20 - 27
	1, 1, 1, 1, 1, 1, 1, 1, // 28 - 2f
	1, 1, 1, 1, 1, 1, 1, 1, // 30 - 37
	1, 1, 1, 1, 1, 1, 1, 1, // 38 - 3f
	1, 1, 1, 1, 1, 1, 1, 1, // 40 - 47
	1, 1, 1, 1, 1, 1, 1, 1, // 48 - 4f
	1, 1, 1, 1, 1, 1, 1, 1, // 50 - 57
	1, 1, 1, 1, 1, 1, 1, 1, // 58 - 5f
	1, 1, 1, 1, 1, 1, 1, 1, // 60 - 67
	1, 1, 1, 1, 1, 1, 1, 1, // 68 - 6f
	1, 1, 1, 1, 1, 1, 1, 1, // 70 - 77
	1, 1, 1, 1, 1, 1, 1, 1, // 78 - 7f
	2, 2, 2, 2, 3, 3, 3, 3, // 80 - 87
	4, 4, 4, 4, 4, 4, 4, 4, // 88 - 8f
	4, 4, 4, 4, 4, 4, 4, 4, // 90 - 97
	4, 4, 4, 4, 4, 4, 4, 4, // 98 - 9f
	5, 5, 5, 5, 5, 5, 5, 5, // a0 - a7
	5, 5, 5, 5, 5, 5, 5, 5, // a8 - af
	5, 5, 5, 5, 5, 5, 5, 5, // b0 - b7
	5, 5, 5, 5, 5, 5, 5, 5, // b8 - bf
	0, 0, 6, 6, 6, 6, 6, 6, // c0 - c7
	6, 6, 6, 6, 6, 6, 6, 6, // c8 - cf
	6, 6, 6, 6, 6, 6, 6, 6, // d0 - d7
	6, 6, 6, 6, 6, 6, 6, 6, // d8 - df
	7, 8, 8, 8, 8, 8, 8, 8, // e0 - e7
	8, 8, 8, 8, 8, 9, 8, 8, // e8 - ef
	10, 11, 11, 11, 11, 11, 11, 11, // f0 - f7
	12, 13, 13, 13, 14, 15, 0, 0, // f8 - ff
}

var UTF8_ST = []int{
	ERROR, START, ERROR, ERROR, ERROR, ERROR, 12, 10, //00-07
	9, 11, 8, 7, 6, 5, 4, 3, //08-0f
	ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, //10-17
	ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, //18-1f
	ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, //20-27
	ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, //28-2f
	ERROR, ERROR, 5, 5, 5, 5, ERROR, ERROR, //30-37
	ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, //38-3f
	ERROR, ERROR, ERROR, 5, 5, 5, ERROR, ERROR, //40-47
	ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, //48-4f
	ERROR, ERROR, 7, 7, 7, 7, ERROR, ERROR, //50-57
	ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, //58-5f
	ERROR, ERROR, ERROR, ERROR, 7, 7, ERROR, ERROR, //60-67
	ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, //68-6f
	ERROR, ERROR, 9, 9, 9, 9, ERROR, ERROR, //70-77
	ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, //78-7f
	ERROR, ERROR, ERROR, ERROR, ERROR, 9, ERROR, ERROR, //80-87
	ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, //88-8f
	ERROR, ERROR, 12, 12, 12, 12, ERROR, ERROR, //90-97
	ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, //98-9f
	ERROR, ERROR, ERROR, ERROR, ERROR, 12, ERROR, ERROR, //a0-a7
	ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, //a8-af
	ERROR, ERROR, 12, 12, 12, ERROR, ERROR, ERROR, //b0-b7
	ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, //b8-bf
	ERROR, ERROR, START, START, START, START, ERROR, ERROR, //c0-c7
	ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, //c8-cf
}

// fmt: on

var UTF8_CHAR_LEN_TABLE = []int{0, 1, 0, 0, 0, 0, 2, 3, 3, 3, 4, 4, 5, 5, 6, 6}

var UTF8_SM_MODEL = CodingStateMachineDict{
	classTable:   UTF8_CLS,
	classFactor:  16,
	stateTable:   UTF8_ST,
	charLenTable: UTF8_CHAR_LEN_TABLE,
	name:         "UTF-8",
}

// Shift_JIS
// fmt: off
var SJIS_CLS = []int{
	1, 1, 1, 1, 1, 1, 1, 1, // 00 - 07
	1, 1, 1, 1, 1, 1, 0, 0, // 08 - 0f
	1, 1, 1, 1, 1, 1, 1, 1, // 10 - 17
	1, 1, 1, 0, 1, 1, 1, 1, // 18 - 1f
	1, 1, 1, 1, 1, 1, 1, 1, // 20 - 27
	1, 1, 1, 1, 1, 1, 1, 1, // 28 - 2f
	1, 1, 1, 1, 1, 1, 1, 1, // 30 - 37
	1, 1, 1, 1, 1, 1, 1, 1, // 38 - 3f
	2, 2, 2, 2, 2, 2, 2, 2, // 40 - 47
	2, 2, 2, 2, 2, 2, 2, 2, // 48 - 4f
	2, 2, 2, 2, 2, 2, 2, 2, // 50 - 57
	2, 2, 2, 2, 2, 2, 2, 2, // 58 - 5f
	2, 2, 2, 2, 2, 2, 2, 2, // 60 - 67
	2, 2, 2, 2, 2, 2, 2, 2, // 68 - 6f
	2, 2, 2, 2, 2, 2, 2, 2, // 70 - 77
	2, 2, 2, 2, 2, 2, 2, 1, // 78 - 7f
	3, 3, 3, 3, 3, 2, 2, 3, // 80 - 87
	3, 3, 3, 3, 3, 3, 3, 3, // 88 - 8f
	3, 3, 3, 3, 3, 3, 3, 3, // 90 - 97
	3, 3, 3, 3, 3, 3, 3, 3, // 98 - 9f
	//0xa0 is illegal in sjis encoding, but some pages does
	//contain such byte. We need to be more error forgiven.
	2, 2, 2, 2, 2, 2, 2, 2, // a0 - a7
	2, 2, 2, 2, 2, 2, 2, 2, // a8 - af
	2, 2, 2, 2, 2, 2, 2, 2, // b0 - b7
	2, 2, 2, 2, 2, 2, 2, 2, // b8 - bf
	2, 2, 2, 2, 2, 2, 2, 2, // c0 - c7
	2, 2, 2, 2, 2, 2, 2, 2, // c8 - cf
	2, 2, 2, 2, 2, 2, 2, 2, // d0 - d7
	2, 2, 2, 2, 2, 2, 2, 2, // d8 - df
	3, 3, 3, 3, 3, 3, 3, 3, // e0 - e7
	3, 3, 3, 3, 3, 4, 4, 4, // e8 - ef
	3, 3, 3, 3, 3, 3, 3, 3, // f0 - f7
	3, 3, 3, 3, 3, 0, 0, 0, // f8 - ff
}

var SJIS_ST = []int{
	ERROR, START, START, 3, ERROR, ERROR, ERROR, ERROR, //00-07
	ERROR, ERROR, ERROR, ERROR, ITS_ME, ITS_ME, ITS_ME, ITS_ME, //08-0f
	ITS_ME, ITS_ME, ERROR, ERROR, START, START, START, START, //10-17
}

// fmt: on

var SJIS_CHAR_LEN_TABLE = []int{0, 1, 1, 2, 0, 0}

var SJIS_SM_MODEL = CodingStateMachineDict{
	classTable:   SJIS_CLS,
	classFactor:  6,
	stateTable:   SJIS_ST,
	charLenTable: SJIS_CHAR_LEN_TABLE,
	name:         "Shift_JIS",
}

// UCS2-BE
// fmt: off
var UCS2BE_CLS = []int{
	0, 0, 0, 0, 0, 0, 0, 0, // 00 - 07
	0, 0, 1, 0, 0, 2, 0, 0, // 08 - 0f
	0, 0, 0, 0, 0, 0, 0, 0, // 10 - 17
	0, 0, 0, 3, 0, 0, 0, 0, // 18 - 1f
	0, 0, 0, 0, 0, 0, 0, 0, // 20 - 27
	0, 3, 3, 3, 3, 3, 0, 0, // 28 - 2f
	0, 0, 0, 0, 0, 0, 0, 0, // 30 - 37
	0, 0, 0, 0, 0, 0, 0, 0, // 38 - 3f
	0, 0, 0, 0, 0, 0, 0, 0, // 40 - 47
	0, 0, 0, 0, 0, 0, 0, 0, // 48 - 4f
	0, 0, 0, 0, 0, 0, 0, 0, // 50 - 57
	0, 0, 0, 0, 0, 0, 0, 0, // 58 - 5f
	0, 0, 0, 0, 0, 0, 0, 0, // 60 - 67
	0, 0, 0, 0, 0, 0, 0, 0, // 68 - 6f
	0, 0, 0, 0, 0, 0, 0, 0, // 70 - 77
	0, 0, 0, 0, 0, 0, 0, 0, // 78 - 7f
	0, 0, 0, 0, 0, 0, 0, 0, // 80 - 87
	0, 0, 0, 0, 0, 0, 0, 0, // 88 - 8f
	0, 0, 0, 0, 0, 0, 0, 0, // 90 - 97
	0, 0, 0, 0, 0, 0, 0, 0, // 98 - 9f
	0, 0, 0, 0, 0, 0, 0, 0, // a0 - a7
	0, 0, 0, 0, 0, 0, 0, 0, // a8 - af
	0, 0, 0, 0, 0, 0, 0, 0, // b0 - b7
	0, 0, 0, 0, 0, 0, 0, 0, // b8 - bf
	0, 0, 0, 0, 0, 0, 0, 0, // c0 - c7
	0, 0, 0, 0, 0, 0, 0, 0, // c8 - cf
	0, 0, 0, 0, 0, 0, 0, 0, // d0 - d7
	0, 0, 0, 0, 0, 0, 0, 0, // d8 - df
	0, 0, 0, 0, 0, 0, 0, 0, // e0 - e7
	0, 0, 0, 0, 0, 0, 0, 0, // e8 - ef
	0, 0, 0, 0, 0, 0, 0, 0, // f0 - f7
	0, 0, 0, 0, 0, 0, 4, 5, // f8 - ff
}

var UCS2BE_ST = []int{
	5, 7, 7, ERROR, 4, 3, ERROR, ERROR, //00-07
	ERROR, ERROR, ERROR, ERROR, ITS_ME, ITS_ME, ITS_ME, ITS_ME, //08-0f
	ITS_ME, ITS_ME, 6, 6, 6, 6, ERROR, ERROR, //10-17
	6, 6, 6, 6, 6, ITS_ME, 6, 6, //18-1f
	6, 6, 6, 6, 5, 7, 7, ERROR, //20-27
	5, 8, 6, 6, ERROR, 6, 6, 6, //28-2f
	6, 6, 6, 6, ERROR, ERROR, START, START, //30-37
}

// fmt: on

var UCS2BE_CHAR_LEN_TABLE = []int{2, 2, 2, 0, 2, 2}

var UCS2BE_SM_MODEL = CodingStateMachineDict{
	classTable:   UCS2BE_CLS,
	classFactor:  6,
	stateTable:   UCS2BE_ST,
	charLenTable: UCS2BE_CHAR_LEN_TABLE,
	name:         "UTF-16BE",
}

// UCS2-LE
// fmt: off
var UCS2LE_CLS = []int{
	0, 0, 0, 0, 0, 0, 0, 0, // 00 - 07
	0, 0, 1, 0, 0, 2, 0, 0, // 08 - 0f
	0, 0, 0, 0, 0, 0, 0, 0, // 10 - 17
	0, 0, 0, 3, 0, 0, 0, 0, // 18 - 1f
	0, 0, 0, 0, 0, 0, 0, 0, // 20 - 27
	0, 3, 3, 3, 3, 3, 0, 0, // 28 - 2f
	0, 0, 0, 0, 0, 0, 0, 0, // 30 - 37
	0, 0, 0, 0, 0, 0, 0, 0, // 38 - 3f
	0, 0, 0, 0, 0, 0, 0, 0, // 40 - 47
	0, 0, 0, 0, 0, 0, 0, 0, // 48 - 4f
	0, 0, 0, 0, 0, 0, 0, 0, // 50 - 57
	0, 0, 0, 0, 0, 0, 0, 0, // 58 - 5f
	0, 0, 0, 0, 0, 0, 0, 0, // 60 - 67
	0, 0, 0, 0, 0, 0, 0, 0, // 68 - 6f
	0, 0, 0, 0, 0, 0, 0, 0, // 70 - 77
	0, 0, 0, 0, 0, 0, 0, 0, // 78 - 7f
	0, 0, 0, 0, 0, 0, 0, 0, // 80 - 87
	0, 0, 0, 0, 0, 0, 0, 0, // 88 - 8f
	0, 0, 0, 0, 0, 0, 0, 0, // 90 - 97
	0, 0, 0, 0, 0, 0, 0, 0, // 98 - 9f
	0, 0, 0, 0, 0, 0, 0, 0, // a0 - a7
	0, 0, 0, 0, 0, 0, 0, 0, // a8 - af
	0, 0, 0, 0, 0, 0, 0, 0, // b0 - b7
	0, 0, 0, 0, 0, 0, 0, 0, // b8 - bf
	0, 0, 0, 0, 0, 0, 0, 0, // c0 - c7
	0, 0, 0, 0, 0, 0, 0, 0, // c8 - cf
	0, 0, 0, 0, 0, 0, 0, 0, // d0 - d7
	0, 0, 0, 0, 0, 0, 0, 0, // d8 - df
	0, 0, 0, 0, 0, 0, 0, 0, // e0 - e7
	0, 0, 0, 0, 0, 0, 0, 0, // e8 - ef
	0, 0, 0, 0, 0, 0, 0, 0, // f0 - f7
	0, 0, 0, 0, 0, 0, 4, 5, // f8 - ff
}

var UCS2LE_ST = []int{
	6, 6, 7, 6, 4, 3, ERROR, ERROR, //00-07
	ERROR, ERROR, ERROR, ERROR, ITS_ME, ITS_ME, ITS_ME, ITS_ME, //08-0f
	ITS_ME, ITS_ME, 5, 5, 5, ERROR, ITS_ME, ERROR, //10-17
	5, 5, 5, ERROR, 5, ERROR, 6, 6, //18-1f
	7, 6, 8, 8, 5, 5, 5, ERROR, //20-27
	5, 5, 5, ERROR, ERROR, ERROR, 5, 5, //28-2f
	5, 5, 5, ERROR, 5, ERROR, START, START, //30-37
}

// fmt: on

var UCS2LE_CHAR_LEN_TABLE = []int{2, 2, 2, 2, 2, 2}

var UCS2LE_SM_MODEL = CodingStateMachineDict{
	classTable:   UCS2LE_CLS,
	classFactor:  6,
	stateTable:   UCS2LE_ST,
	charLenTable: UCS2LE_CHAR_LEN_TABLE,
	name:         "UTF-16LE",
}

// CP949
// fmt: off
var CP949_CLS = []int{
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, // 00 - 0f
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, // 10 - 1f
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, // 20 - 2f
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, // 30 - 3f
	1, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, // 40 - 4f
	4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 1, 1, 1, 1, 1, // 50 - 5f
	1, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, // 60 - 6f
	5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 1, 1, 1, 1, 1, // 70 - 7f
	0, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, // 80 - 8f
	6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, // 90 - 9f
	6, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 8, 8, 8, // a0 - af
	7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, // b0 - bf
	7, 7, 7, 7, 7, 7, 9, 2, 2, 3, 2, 2, 2, 2, 2, 2, // c0 - cf
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, // d0 - df
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, // e0 - ef
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 0, // f0 - ff
}

var CP949_ST = []int{
	//cls=    0      1      2      3      4      5      6      7      8      9  // previous state =
	ERROR, START, 3, ERROR, START, START, 4, 5, ERROR, 6, // START
	ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, // ERROR
	ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, // ITS_ME
	ERROR, ERROR, START, START, ERROR, ERROR, ERROR, START, START, START, // 3
	ERROR, ERROR, START, START, START, START, START, START, START, START, // 4
	ERROR, START, START, START, START, START, START, START, START, START, // 5
	ERROR, START, START, START, START, ERROR, ERROR, START, START, START, // 6
}

// fmt: on

var CP949_CHAR_LEN_TABLE = []int{0, 1, 2, 0, 1, 1, 2, 2, 0, 2}

var CP949_SM_MODEL = CodingStateMachineDict{
	classTable:   CP949_CLS,
	classFactor:  10,
	stateTable:   CP949_ST,
	charLenTable: CP949_CHAR_LEN_TABLE,
	name:         "CP949",
}

// EUC-JP
// fmt: off
var EUCJP_CLS = []int{
	4, 4, 4, 4, 4, 4, 4, 4, // 00 - 07
	4, 4, 4, 4, 4, 4, 5, 5, // 08 - 0f
	4, 4, 4, 4, 4, 4, 4, 4, // 10 - 17
	4, 4, 4, 5, 4, 4, 4, 4, // 18 - 1f
	4, 4, 4, 4, 4, 4, 4, 4, // 20 - 27
	4, 4, 4, 4, 4, 4, 4, 4, // 28 - 2f
	4, 4, 4, 4, 4, 4, 4, 4, // 30 - 37
	4, 4, 4, 4, 4, 4, 4, 4, // 38 - 3f
	4, 4, 4, 4, 4, 4, 4, 4, // 40 - 47
	4, 4, 4, 4, 4, 4, 4, 4, // 48 - 4f
	4, 4, 4, 4, 4, 4, 4, 4, // 50 - 57
	4, 4, 4, 4, 4, 4, 4, 4, // 58 - 5f
	4, 4, 4, 4, 4, 4, 4, 4, // 60 - 67
	4, 4, 4, 4, 4, 4, 4, 4, // 68 - 6f
	4, 4, 4, 4, 4, 4, 4, 4, // 70 - 77
	4, 4, 4, 4, 4, 4, 4, 4, // 78 - 7f
	5, 5, 5, 5, 5, 5, 5, 5, // 80 - 87
	5, 5, 5, 5, 5, 5, 1, 3, // 88 - 8f
	5, 5, 5, 5, 5, 5, 5, 5, // 90 - 97
	5, 5, 5, 5, 5, 5, 5, 5, // 98 - 9f
	5, 2, 2, 2, 2, 2, 2, 2, // a0 - a7
	2, 2, 2, 2, 2, 2, 2, 2, // a8 - af
	2, 2, 2, 2, 2, 2, 2, 2, // b0 - b7
	2, 2, 2, 2, 2, 2, 2, 2, // b8 - bf
	2, 2, 2, 2, 2, 2, 2, 2, // c0 - c7
	2, 2, 2, 2, 2, 2, 2, 2, // c8 - cf
	2, 2, 2, 2, 2, 2, 2, 2, // d0 - d7
	2, 2, 2, 2, 2, 2, 2, 2, // d8 - df
	0, 0, 0, 0, 0, 0, 0, 0, // e0 - e7
	0, 0, 0, 0, 0, 0, 0, 0, // e8 - ef
	0, 0, 0, 0, 0, 0, 0, 0, // f0 - f7
	0, 0, 0, 0, 0, 0, 0, 5, // f8 - ff
}

var EUCJP_ST = []int{
	3, 4, 3, 5, START, ERROR, ERROR, ERROR, //00-07
	ERROR, ERROR, ERROR, ERROR, ITS_ME, ITS_ME, ITS_ME, ITS_ME, //08-0f
	ITS_ME, ITS_ME, START, ERROR, START, ERROR, ERROR, ERROR, //10-17
	ERROR, ERROR, START, ERROR, ERROR, ERROR, 3, ERROR, //18-1f
	3, ERROR, ERROR, ERROR, START, START, START, START, //20-27
}

// fmt: on

var EUCJP_CHAR_LEN_TABLE = []int{2, 2, 2, 3, 1, 0}

var EUCJP_SM_MODEL = CodingStateMachineDict{
	classTable:   EUCJP_CLS,
	classFactor:  6,
	stateTable:   EUCJP_ST,
	charLenTable: EUCJP_CHAR_LEN_TABLE,
	name:         "EUC-JP",
}

// EUC-KR
// fmt: off
var EUCKR_CLS = []int{
	1, 1, 1, 1, 1, 1, 1, 1, // 00 - 07
	1, 1, 1, 1, 1, 1, 0, 0, // 08 - 0f
	1, 1, 1, 1, 1, 1, 1, 1, // 10 - 17
	1, 1, 1, 0, 1, 1, 1, 1, // 18 - 1f
	1, 1, 1, 1, 1, 1, 1, 1, // 20 - 27
	1, 1, 1, 1, 1, 1, 1, 1, // 28 - 2f
	1, 1, 1, 1, 1, 1, 1, 1, // 30 - 37
	1, 1, 1, 1, 1, 1, 1, 1, // 38 - 3f
	1, 1, 1, 1, 1, 1, 1, 1, // 40 - 47
	1, 1, 1, 1, 1, 1, 1, 1, // 48 - 4f
	1, 1, 1, 1, 1, 1, 1, 1, // 50 - 57
	1, 1, 1, 1, 1, 1, 1, 1, // 58 - 5f
	1, 1, 1, 1, 1, 1, 1, 1, // 60 - 67
	1, 1, 1, 1, 1, 1, 1, 1, // 68 - 6f
	1, 1, 1, 1, 1, 1, 1, 1, // 70 - 77
	1, 1, 1, 1, 1, 1, 1, 1, // 78 - 7f
	0, 0, 0, 0, 0, 0, 0, 0, // 80 - 87
	0, 0, 0, 0, 0, 0, 0, 0, // 88 - 8f
	0, 0, 0, 0, 0, 0, 0, 0, // 90 - 97
	0, 0, 0, 0, 0, 0, 0, 0, // 98 - 9f
	0, 2, 2, 2, 2, 2, 2, 2, // a0 - a7
	2, 2, 2, 2, 2, 3, 3, 3, // a8 - af
	2, 2, 2, 2, 2, 2, 2, 2, // b0 - b7
	2, 2, 2, 2, 2, 2, 2, 2, // b8 - bf
	2, 2, 2, 2, 2, 2, 2, 2, // c0 - c7
	2, 3, 2, 2, 2, 2, 2, 2, // c8 - cf
	2, 2, 2, 2, 2, 2, 2, 2, // d0 - d7
	2, 2, 2, 2, 2, 2, 2, 2, // d8 - df
	2, 2, 2, 2, 2, 2, 2, 2, // e0 - e7
	2, 2, 2, 2, 2, 2, 2, 2, // e8 - ef
	2, 2, 2, 2, 2, 2, 2, 2, // f0 - f7
	2, 2, 2, 2, 2, 2, 2, 0, // f8 - ff
}

var EUCKR_ST = []int{
	ERROR, START, 3, ERROR, ERROR, ERROR, ERROR, ERROR, //00-07
	ITS_ME, ITS_ME, ITS_ME, ITS_ME, ERROR, ERROR, START, START, //08-0f
}

// fmt: on

var EUCKR_CHAR_LEN_TABLE = []int{0, 1, 2, 0}

var EUCKR_SM_MODEL = CodingStateMachineDict{
	classTable:   EUCKR_CLS,
	classFactor:  4,
	stateTable:   EUCKR_ST,
	charLenTable: EUCKR_CHAR_LEN_TABLE,
	name:         "EUC-KR",
}

// JOHAB
// fmt: off
var JOHAB_CLS = []int{
	4, 4, 4, 4, 4, 4, 4, 4, // 00 - 07
	4, 4, 4, 4, 4, 4, 0, 0, // 08 - 0f
	4, 4, 4, 4, 4, 4, 4, 4, // 10 - 17
	4, 4, 4, 0, 4, 4, 4, 4, // 18 - 1f
	4, 4, 4, 4, 4, 4, 4, 4, // 20 - 27
	4, 4, 4, 4, 4, 4, 4, 4, // 28 - 2f
	4, 3, 3, 3, 3, 3, 3, 3, // 30 - 37
	3, 3, 3, 3, 3, 3, 3, 3, // 38 - 3f
	3, 1, 1, 1, 1, 1, 1, 1, // 40 - 47
	1, 1, 1, 1, 1, 1, 1, 1, // 48 - 4f
	1, 1, 1, 1, 1, 1, 1, 1, // 50 - 57
	1, 1, 1, 1, 1, 1, 1, 1, // 58 - 5f
	1, 1, 1, 1, 1, 1, 1, 1, // 60 - 67
	1, 1, 1, 1, 1, 1, 1, 1, // 68 - 6f
	1, 1, 1, 1, 1, 1, 1, 1, // 70 - 77
	1, 1, 1, 1, 1, 1, 1, 2, // 78 - 7f
	6, 6, 6, 6, 8, 8, 8, 8, // 80 - 87
	8, 8, 8, 8, 8, 8, 8, 8, // 88 - 8f
	8, 7, 7, 7, 7, 7, 7, 7, // 90 - 97
	7, 7, 7, 7, 7, 7, 7, 7, // 98 - 9f
	7, 7, 7, 7, 7, 7, 7, 7, // a0 - a7
	7, 7, 7, 7, 7, 7, 7, 7, // a8 - af
	7, 7, 7, 7, 7, 7, 7, 7, // b0 - b7
	7, 7, 7, 7, 7, 7, 7, 7, // b8 - bf
	7, 7, 7, 7, 7, 7, 7, 7, // c0 - c7
	7, 7, 7, 7, 7, 7, 7, 7, // c8 - cf
	7, 7, 7, 7, 5, 5, 5, 5, // d0 - d7
	5, 9, 9, 9, 9, 9, 9, 5, // d8 - df
	9, 9, 9, 9, 9, 9, 9, 9, // e0 - e7
	9, 9, 9, 9, 9, 9, 9, 9, // e8 - ef
	9, 9, 9, 9, 9, 9, 9, 9, // f0 - f7
	9, 9, 5, 5, 5, 5, 5, 0, // f8 - ff
}

var JOHAB_ST = []int{
	// cls = 0                   1                   2                   3                   4                   5                   6                   7                   8                   9
	ERROR, START, START, START, START, ERROR, ERROR, 3, 3, 4, // START
	ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, // ITS_ME
	ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, // ERROR
	ERROR, START, START, ERROR, ERROR, START, START, START, START, START, // 3
	ERROR, START, ERROR, START, ERROR, START, ERROR, START, ERROR, START, // 4
}

// fmt: on

var JOHAB_CHAR_LEN_TABLE = []int{0, 1, 1, 1, 1, 0, 0, 2, 2, 2}

var JOHAB_SM_MODEL = CodingStateMachineDict{
	classTable:   JOHAB_CLS,
	classFactor:  10,
	stateTable:   JOHAB_ST,
	charLenTable: JOHAB_CHAR_LEN_TABLE,
	name:         "Johab",
}

// EUC-TW
// fmt: off
var EUCTW_CLS = []int{
	2, 2, 2, 2, 2, 2, 2, 2, // 00 - 07
	2, 2, 2, 2, 2, 2, 0, 0, // 08 - 0f
	2, 2, 2, 2, 2, 2, 2, 2, // 10 - 17
	2, 2, 2, 0, 2, 2, 2, 2, // 18 - 1f
	2, 2, 2, 2, 2, 2, 2, 2, // 20 - 27
	2, 2, 2, 2, 2, 2, 2, 2, // 28 - 2f
	2, 2, 2, 2, 2, 2, 2, 2, // 30 - 37
	2, 2, 2, 2, 2, 2, 2, 2, // 38 - 3f
	2, 2, 2, 2, 2, 2, 2, 2, // 40 - 47
	2, 2, 2, 2, 2, 2, 2, 2, // 48 - 4f
	2, 2, 2, 2, 2, 2, 2, 2, // 50 - 57
	2, 2, 2, 2, 2, 2, 2, 2, // 58 - 5f
	2, 2, 2, 2, 2, 2, 2, 2, // 60 - 67
	2, 2, 2, 2, 2, 2, 2, 2, // 68 - 6f
	2, 2, 2, 2, 2, 2, 2, 2, // 70 - 77
	2, 2, 2, 2, 2, 2, 2, 2, // 78 - 7f
	0, 0, 0, 0, 0, 0, 0, 0, // 80 - 87
	0, 0, 0, 0, 0, 0, 6, 0, // 88 - 8f
	0, 0, 0, 0, 0, 0, 0, 0, // 90 - 97
	0, 0, 0, 0, 0, 0, 0, 0, // 98 - 9f
	0, 3, 4, 4, 4, 4, 4, 4, // a0 - a7
	5, 5, 1, 1, 1, 1, 1, 1, // a8 - af
	1, 1, 1, 1, 1, 1, 1, 1, // b0 - b7
	1, 1, 1, 1, 1, 1, 1, 1, // b8 - bf
	1, 1, 3, 1, 3, 3, 3, 3, // c0 - c7
	3, 3, 3, 3, 3, 3, 3, 3, // c8 - cf
	3, 3, 3, 3, 3, 3, 3, 3, // d0 - d7
	3, 3, 3, 3, 3, 3, 3, 3, // d8 - df
	3, 3, 3, 3, 3, 3, 3, 3, // e0 - e7
	3, 3, 3, 3, 3, 3, 3, 3, // e8 - ef
	3, 3, 3, 3, 3, 3, 3, 3, // f0 - f7
	3, 3, 3, 3, 3, 3, 3, 0, // f8 - ff
}

var EUCTW_ST = []int{
	ERROR, ERROR, START, 3, 3, 3, 4, ERROR, //00-07
	ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ITS_ME, ITS_ME, //08-0f
	ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, ERROR, START, ERROR, //10-17
	START, START, START, ERROR, ERROR, ERROR, ERROR, ERROR, //18-1f
	5, ERROR, ERROR, ERROR, START, ERROR, START, START, //20-27
	START, ERROR, START, START, START, START, START, START, //28-2f
}

// fmt: on

var EUCTW_CHAR_LEN_TABLE = []int{0, 0, 1, 2, 2, 2, 3}

var EUCTW_SM_MODEL = CodingStateMachineDict{
	classTable:   EUCTW_CLS,
	classFactor:  7,
	stateTable:   EUCTW_ST,
	charLenTable: EUCTW_CHAR_LEN_TABLE,
	name:         "x-euc-tw",
}

var UTF8_CLS_NEW = []int{
	1, 1, 1, 1, 1, 1, 1, 1, // 00 - 07  //allow 0x00 as a legal value
	1, 1, 1, 1, 1, 1, 1, 1, // 08 - 0f
	1, 1, 1, 1, 1, 1, 1, 1, // 10 - 17
	1, 1, 1, 1, 1, 1, 1, 1, // 18 - 1f
	1, 1, 1, 1, 1, 1, 1, 1, // 20 - 27
	1, 1, 1, 1, 1, 1, 1, 1, // 28 - 2f
	1, 1, 1, 1, 1, 1, 1, 1, // 30 - 37
	1, 1, 1, 1, 1, 1, 1, 1, // 38 - 3f
	1, 1, 1, 1, 1, 1, 1, 1, // 40 - 47
	1, 1, 1, 1, 1, 1, 1, 1, // 48 - 4f
	1, 1, 1, 1, 1, 1, 1, 1, // 50 - 57
	1, 1, 1, 1, 1, 1, 1, 1, // 58 - 5f
	1, 1, 1, 1, 1, 1, 1, 1, // 60 - 67
	1, 1, 1, 1, 1, 1, 1, 1, // 68 - 6f
	1, 1, 1, 1, 1, 1, 1, 1, // 70 - 77
	1, 1, 1, 1, 1, 1, 1, 1, // 78 - 7f
	2, 2, 2, 2, 2, 2, 2, 2, // 80 - 87
	2, 2, 2, 2, 2, 2, 2, 2, // 88 - 8f
	3, 3, 3, 3, 3, 3, 3, 3, // 90 - 97
	3, 3, 3, 3, 3, 3, 3, 3, // 98 - 9f
	4, 4, 4, 4, 4, 4, 4, 4, // a0 - a7
	4, 4, 4, 4, 4, 4, 4, 4, // a8 - af
	4, 4, 4, 4, 4, 4, 4, 4, // b0 - b7
	4, 4, 4, 4, 4, 4, 4, 4, // b8 - bf
	0, 0, 5, 5, 5, 5, 5, 5, // c0 - c7
	5, 5, 5, 5, 5, 5, 5, 5, // c8 - cf
	5, 5, 5, 5, 5, 5, 5, 5, // d0 - d7
	5, 5, 5, 5, 5, 5, 5, 5, // d8 - df
	6, 7, 7, 7, 7, 7, 7, 7, // e0 - e7
	7, 7, 7, 7, 7, 8, 7, 7, // e8 - ef
	9, 10, 10, 10, 11, 0, 0, 0, // f0 - f7
	0, 0, 0, 0, 0, 0, 0, 0, // f8 - ff
}

var UTF8_ST_NEW = []int{
	ERROR, START, ERROR, ERROR, ERROR, 3, 5, 4, 6, 8, 7, 9, //start
	ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR,
	ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME, ITS_ME,
	ERROR, ERROR, START, START, START, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, //a
	ERROR, ERROR, 3, 3, 3, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, //b
	ERROR, ERROR, ERROR, ERROR, 3, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, //c
	ERROR, ERROR, 3, 3, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, //d
	ERROR, ERROR, 4, 4, 4, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, //e
	ERROR, ERROR, ERROR, 4, 4, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, //f
	ERROR, ERROR, 4, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, ERROR, //g
}

// fmt: on

var UTF8_CHAR_LEN_TABLE_NEW = []int{0, 1, 0, 0, 0, 2, 3, 3, 3, 4, 4, 4}

var UTF8_SM_MODEL_NEW = CodingStateMachineDict{
	classTable:   UTF8_CLS_NEW,
	classFactor:  12,
	stateTable:   UTF8_ST_NEW,
	charLenTable: UTF8_CHAR_LEN_TABLE_NEW,
	name:         "UTF-8",
}
