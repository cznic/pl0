// CAUTION: Generated file - DO NOT EDIT.

// Copyright 2015 The PL0 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pl0

import __yyfmt__ "fmt"

import (
	"strconv"

	"github.com/cznic/xc"
)

type yySymType struct {
	yys   int
	node  Node
	Token xc.Token
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault = 57367
	yyEofCode = 57344
	ASSIGN    = 57346
	BEGIN     = 57347
	CALL      = 57348
	CONST     = 57349
	DO        = 57350
	END       = 57351
	GEQ       = 57352
	IDENT     = 57353
	IF        = 57354
	LEQ       = 57355
	NUMBER    = 57356
	ODD       = 57357
	PROCEDURE = 57358
	THEN      = 57359
	VAR       = 57360
	WHILE     = 57361
	yyErrCode = 57345

	yyMaxDepth = 200
	yyTabOfs   = -54
)

var (
	yyXLAT = map[int]int{
		59:    0,  // ';' (51x)
		57353: 1,  // IDENT (44x)
		43:    2,  // '+' (36x)
		45:    3,  // '-' (36x)
		46:    4,  // '.' (36x)
		57351: 5,  // END (26x)
		57350: 6,  // DO (21x)
		57359: 7,  // THEN (21x)
		57356: 8,  // NUMBER (20x)
		40:    9,  // '(' (19x)
		57376: 10, // Number (19x)
		33:    11, // '!' (18x)
		63:    12, // '?' (18x)
		57347: 13, // BEGIN (18x)
		57348: 14, // CALL (18x)
		57375: 15, // Factor (18x)
		57354: 16, // IF (18x)
		57361: 17, // WHILE (18x)
		57383: 18, // Term (16x)
		61:    19, // '=' (14x)
		35:    20, // '#' (13x)
		41:    21, // ')' (13x)
		60:    22, // '<' (13x)
		62:    23, // '>' (13x)
		57352: 24, // GEQ (13x)
		57355: 25, // LEQ (13x)
		42:    26, // '*' (12x)
		47:    27, // '/' (12x)
		57374: 28, // Expression (12x)
		57358: 29, // PROCEDURE (12x)
		44:    30, // ',' (9x)
		57381: 31, // Statement (5x)
		57360: 32, // VAR (5x)
		57357: 33, // ODD (3x)
		57344: 34, // $end (2x)
		57368: 35, // Block (2x)
		57369: 36, // Condition (2x)
		57349: 37, // CONST (2x)
		57372: 38, // Consts (2x)
		57373: 39, // ConstsOpt (2x)
		57370: 40, // ConstSpec (2x)
		57379: 41, // ProcSpec (2x)
		57384: 42, // Variable (2x)
		57362: 43, // $@1 (1x)
		57363: 44, // $@2 (1x)
		57364: 45, // $@3 (1x)
		57365: 46, // $@4 (1x)
		57366: 47, // $@5 (1x)
		57346: 48, // ASSIGN (1x)
		57371: 49, // ConstSpecList (1x)
		57377: 50, // ProcList (1x)
		57378: 51, // ProcListOpt (1x)
		57380: 52, // Program (1x)
		57382: 53, // StatementList (1x)
		57385: 54, // VariableList (1x)
		57386: 55, // Vars (1x)
		57387: 56, // VarsOpt (1x)
		57367: 57, // $default (0x)
		57345: 58, // error (0x)
	}

	yySymNames = []string{
		"';'",
		"IDENT",
		"'+'",
		"'-'",
		"'.'",
		"END",
		"DO",
		"THEN",
		"NUMBER",
		"'('",
		"Number",
		"'!'",
		"'?'",
		"BEGIN",
		"CALL",
		"Factor",
		"IF",
		"WHILE",
		"Term",
		"'='",
		"'#'",
		"')'",
		"'<'",
		"'>'",
		"GEQ",
		"LEQ",
		"'*'",
		"'/'",
		"Expression",
		"PROCEDURE",
		"','",
		"Statement",
		"VAR",
		"ODD",
		"$end",
		"Block",
		"Condition",
		"CONST",
		"Consts",
		"ConstsOpt",
		"ConstSpec",
		"ProcSpec",
		"Variable",
		"$@1",
		"$@2",
		"$@3",
		"$@4",
		"$@5",
		"ASSIGN",
		"ConstSpecList",
		"ProcList",
		"ProcListOpt",
		"Program",
		"StatementList",
		"VariableList",
		"Vars",
		"VarsOpt",
		"$default",
		"error",
	}

	yyReductions = map[int]struct{ xsym, components int }{
		0:  {0, 1},
		1:  {52, 2},
		2:  {43, 0},
		3:  {35, 5},
		4:  {39, 0},
		5:  {39, 1},
		6:  {38, 3},
		7:  {49, 1},
		8:  {49, 3},
		9:  {40, 3},
		10: {10, 1},
		11: {56, 0},
		12: {56, 1},
		13: {55, 3},
		14: {42, 1},
		15: {54, 1},
		16: {54, 3},
		17: {51, 0},
		18: {51, 1},
		19: {50, 1},
		20: {50, 2},
		21: {44, 0},
		22: {41, 6},
		23: {31, 0},
		24: {31, 3},
		25: {31, 2},
		26: {31, 2},
		27: {31, 2},
		28: {31, 3},
		29: {45, 0},
		30: {31, 5},
		31: {46, 0},
		32: {47, 0},
		33: {31, 6},
		34: {53, 1},
		35: {53, 3},
		36: {36, 2},
		37: {36, 3},
		38: {36, 3},
		39: {36, 3},
		40: {36, 3},
		41: {36, 3},
		42: {36, 3},
		43: {28, 2},
		44: {28, 2},
		45: {28, 1},
		46: {28, 3},
		47: {28, 3},
		48: {18, 1},
		49: {18, 3},
		50: {18, 3},
		51: {15, 1},
		52: {15, 1},
		53: {15, 3},
	}

	yyXErrors = map[yyXError]string{
		yyXError{0, 34}:  "invalid empty input",
		yyXError{1, -1}:  "expected $end",
		yyXError{96, -1}: "expected $end",
		yyXError{2, -1}:  "expected '.'",
		yyXError{28, -1}: "expected ';'",
		yyXError{29, -1}: "expected ';'",
		yyXError{31, -1}: "expected ';'",
		yyXError{8, -1}:  "expected '='",
		yyXError{36, -1}: "expected :=",
		yyXError{44, -1}: "expected DO",
		yyXError{79, -1}: "expected DO",
		yyXError{82, -1}: "expected THEN",
		yyXError{83, -1}: "expected THEN",
		yyXError{30, -1}: "expected block or one of ['!', ';', '?', BEGIN, CALL, CONST, IF, PROCEDURE, VAR, WHILE, identifier]",
		yyXError{41, -1}: "expected conditional expression or one of ['(', '+', '-', ODD, identifier, integer literal]",
		yyXError{42, -1}: "expected conditional expression or one of ['(', '+', '-', ODD, identifier, integer literal]",
		yyXError{43, -1}: "expected conditional expression or one of ['(', '+', '-', ODD, identifier, integer literal]",
		yyXError{13, -1}: "expected constant specification or identifier",
		yyXError{59, -1}: "expected expression factor or one of ['(', identifier, integer literal]",
		yyXError{60, -1}: "expected expression factor or one of ['(', identifier, integer literal]",
		yyXError{39, -1}: "expected expression or one of ['(', '+', '-', identifier, integer literal]",
		yyXError{45, -1}: "expected expression or one of ['(', '+', '-', identifier, integer literal]",
		yyXError{53, -1}: "expected expression or one of ['(', '+', '-', identifier, integer literal]",
		yyXError{66, -1}: "expected expression or one of ['(', '+', '-', identifier, integer literal]",
		yyXError{67, -1}: "expected expression or one of ['(', '+', '-', identifier, integer literal]",
		yyXError{68, -1}: "expected expression or one of ['(', '+', '-', identifier, integer literal]",
		yyXError{69, -1}: "expected expression or one of ['(', '+', '-', identifier, integer literal]",
		yyXError{70, -1}: "expected expression or one of ['(', '+', '-', identifier, integer literal]",
		yyXError{71, -1}: "expected expression or one of ['(', '+', '-', identifier, integer literal]",
		yyXError{94, -1}: "expected expression or one of ['(', '+', '-', identifier, integer literal]",
		yyXError{47, -1}: "expected expression term or one of ['(', identifier, integer literal]",
		yyXError{48, -1}: "expected expression term or one of ['(', identifier, integer literal]",
		yyXError{55, -1}: "expected expression term or one of ['(', identifier, integer literal]",
		yyXError{56, -1}: "expected expression term or one of ['(', identifier, integer literal]",
		yyXError{27, -1}: "expected identifier",
		yyXError{37, -1}: "expected identifier",
		yyXError{38, -1}: "expected identifier",
		yyXError{5, -1}:  "expected list of constant declarations or identifier",
		yyXError{9, -1}:  "expected number or integer literal",
		yyXError{4, -1}:  "expected one of ['!', '.', ';', '?', BEGIN, CALL, IF, PROCEDURE, VAR, WHILE, identifier]",
		yyXError{12, -1}: "expected one of ['!', '.', ';', '?', BEGIN, CALL, IF, PROCEDURE, VAR, WHILE, identifier]",
		yyXError{16, -1}: "expected one of ['!', '.', ';', '?', BEGIN, CALL, IF, PROCEDURE, WHILE, identifier]",
		yyXError{21, -1}: "expected one of ['!', '.', ';', '?', BEGIN, CALL, IF, PROCEDURE, WHILE, identifier]",
		yyXError{26, -1}: "expected one of ['!', '.', ';', '?', BEGIN, CALL, IF, PROCEDURE, WHILE, identifier]",
		yyXError{32, -1}: "expected one of ['!', '.', ';', '?', BEGIN, CALL, IF, PROCEDURE, WHILE, identifier]",
		yyXError{33, -1}: "expected one of ['!', '.', ';', '?', BEGIN, CALL, IF, PROCEDURE, WHILE, identifier]",
		yyXError{11, -1}: "expected one of ['#', ')', '*', '+', ',', '-', '.', '/', ';', '<', '=', '>', <=, >=, DO, END, THEN]",
		yyXError{49, -1}: "expected one of ['#', ')', '*', '+', '-', '.', '/', ';', '<', '=', '>', <=, >=, DO, END, THEN]",
		yyXError{50, -1}: "expected one of ['#', ')', '*', '+', '-', '.', '/', ';', '<', '=', '>', <=, >=, DO, END, THEN]",
		yyXError{51, -1}: "expected one of ['#', ')', '*', '+', '-', '.', '/', ';', '<', '=', '>', <=, >=, DO, END, THEN]",
		yyXError{52, -1}: "expected one of ['#', ')', '*', '+', '-', '.', '/', ';', '<', '=', '>', <=, >=, DO, END, THEN]",
		yyXError{57, -1}: "expected one of ['#', ')', '*', '+', '-', '.', '/', ';', '<', '=', '>', <=, >=, DO, END, THEN]",
		yyXError{58, -1}: "expected one of ['#', ')', '*', '+', '-', '.', '/', ';', '<', '=', '>', <=, >=, DO, END, THEN]",
		yyXError{61, -1}: "expected one of ['#', ')', '*', '+', '-', '.', '/', ';', '<', '=', '>', <=, >=, DO, END, THEN]",
		yyXError{62, -1}: "expected one of ['#', ')', '*', '+', '-', '.', '/', ';', '<', '=', '>', <=, >=, DO, END, THEN]",
		yyXError{63, -1}: "expected one of ['#', ')', '*', '+', '-', '.', '/', ';', '<', '=', '>', <=, >=, DO, END, THEN]",
		yyXError{64, -1}: "expected one of ['#', ')', '*', '+', '-', '.', '/', ';', '<', '=', '>', <=, >=, DO, END, THEN]",
		yyXError{65, -1}: "expected one of ['#', ')', '*', '+', '-', '.', '/', ';', '<', '=', '>', <=, >=, DO, END, THEN]",
		yyXError{46, -1}: "expected one of ['#', '+', '-', '<', '=', '>', <=, >=]",
		yyXError{54, -1}: "expected one of [')', '+', '-']",
		yyXError{91, -1}: "expected one of ['+', '-', '.', ';', END]",
		yyXError{95, -1}: "expected one of ['+', '-', '.', ';', END]",
		yyXError{72, -1}: "expected one of ['+', '-', DO, THEN]",
		yyXError{73, -1}: "expected one of ['+', '-', DO, THEN]",
		yyXError{74, -1}: "expected one of ['+', '-', DO, THEN]",
		yyXError{75, -1}: "expected one of ['+', '-', DO, THEN]",
		yyXError{76, -1}: "expected one of ['+', '-', DO, THEN]",
		yyXError{77, -1}: "expected one of ['+', '-', DO, THEN]",
		yyXError{78, -1}: "expected one of ['+', '-', DO, THEN]",
		yyXError{6, -1}:  "expected one of [',', ';']",
		yyXError{7, -1}:  "expected one of [',', ';']",
		yyXError{10, -1}: "expected one of [',', ';']",
		yyXError{14, -1}: "expected one of [',', ';']",
		yyXError{18, -1}: "expected one of [',', ';']",
		yyXError{19, -1}: "expected one of [',', ';']",
		yyXError{20, -1}: "expected one of [',', ';']",
		yyXError{23, -1}: "expected one of [',', ';']",
		yyXError{81, -1}: "expected one of ['.', ';', END]",
		yyXError{85, -1}: "expected one of ['.', ';', END]",
		yyXError{88, -1}: "expected one of ['.', ';', END]",
		yyXError{92, -1}: "expected one of ['.', ';', END]",
		yyXError{93, -1}: "expected one of ['.', ';', END]",
		yyXError{35, -1}: "expected one of ['.', ';']",
		yyXError{86, -1}: "expected one of [';', END]",
		yyXError{87, -1}: "expected one of [';', END]",
		yyXError{90, -1}: "expected one of [';', END]",
		yyXError{3, -1}:  "expected optional procedure declarations or optional variable declarations or statement or one of ['!', '.', ';', '?', BEGIN, CALL, IF, PROCEDURE, VAR, WHILE, identifier]",
		yyXError{15, -1}: "expected optional procedure declarations or statement or one of ['!', '.', ';', '?', BEGIN, CALL, IF, PROCEDURE, WHILE, identifier]",
		yyXError{25, -1}: "expected procedure defintion or one of ['!', '.', ';', '?', BEGIN, CALL, IF, PROCEDURE, WHILE, identifier]",
		yyXError{0, -1}:  "expected program or one of ['!', '.', '?', BEGIN, CALL, CONST, IF, PROCEDURE, VAR, WHILE, identifier]",
		yyXError{40, -1}: "expected statement list or one of ['!', ';', '?', BEGIN, CALL, END, IF, WHILE, identifier]",
		yyXError{80, -1}: "expected statement or one of ['!', '.', ';', '?', BEGIN, CALL, END, IF, WHILE, identifier]",
		yyXError{84, -1}: "expected statement or one of ['!', '.', ';', '?', BEGIN, CALL, END, IF, WHILE, identifier]",
		yyXError{24, -1}: "expected statement or one of ['!', '.', ';', '?', BEGIN, CALL, IF, WHILE, identifier]",
		yyXError{34, -1}: "expected statement or one of ['!', '.', ';', '?', BEGIN, CALL, IF, WHILE, identifier]",
		yyXError{89, -1}: "expected statement or one of ['!', ';', '?', BEGIN, CALL, END, IF, WHILE, identifier]",
		yyXError{22, -1}: "expected variable name or identifier",
		yyXError{17, -1}: "expected variables list or identifier",
	}

	yyParseTab = [97][]uint16{
		// 0
		{1: 50, 4: 50, 11: 50, 50, 50, 50, 16: 50, 50, 29: 50, 32: 50, 35: 56, 37: 59, 58, 57, 52: 55},
		{34: 54},
		{4: 150},
		{43, 43, 4: 43, 11: 43, 43, 43, 43, 16: 43, 43, 29: 43, 32: 71, 55: 70, 69},
		{49, 49, 4: 49, 11: 49, 49, 49, 49, 16: 49, 49, 29: 49, 32: 49},
		// 5
		{1: 62, 40: 61, 49: 60},
		{66, 30: 67},
		{47, 30: 47},
		{19: 63},
		{8: 65, 10: 64},
		// 10
		{45, 30: 45},
		{44, 2: 44, 44, 44, 44, 44, 44, 19: 44, 44, 44, 44, 44, 44, 44, 44, 44, 30: 44},
		{48, 48, 4: 48, 11: 48, 48, 48, 48, 16: 48, 48, 29: 48, 32: 48},
		{1: 62, 40: 68},
		{46, 30: 46},
		// 15
		{37, 37, 4: 37, 11: 37, 37, 37, 37, 16: 37, 37, 29: 81, 41: 80, 50: 79, 78},
		{42, 42, 4: 42, 11: 42, 42, 42, 42, 16: 42, 42, 29: 42},
		{1: 73, 42: 74, 54: 72},
		{75, 30: 76},
		{40, 30: 40},
		// 20
		{39, 30: 39},
		{41, 41, 4: 41, 11: 41, 41, 41, 41, 16: 41, 41, 29: 41},
		{1: 73, 42: 77},
		{38, 30: 38},
		{52, 52, 4: 52, 11: 52, 52, 52, 52, 16: 52, 52, 43: 88},
		// 25
		{36, 36, 4: 36, 11: 36, 36, 36, 36, 16: 36, 36, 29: 81, 41: 87},
		{35, 35, 4: 35, 11: 35, 35, 35, 35, 16: 35, 35, 29: 35},
		{1: 82},
		{33, 44: 83},
		{84},
		// 30
		{50, 50, 11: 50, 50, 50, 50, 16: 50, 50, 29: 50, 32: 50, 35: 85, 37: 59, 58, 57},
		{86},
		{32, 32, 4: 32, 11: 32, 32, 32, 32, 16: 32, 32, 29: 32},
		{34, 34, 4: 34, 11: 34, 34, 34, 34, 16: 34, 34, 29: 34},
		{31, 90, 4: 31, 11: 93, 92, 94, 91, 16: 95, 96, 31: 89},
		// 35
		{51, 4: 51},
		{48: 148},
		{1: 147},
		{1: 146},
		{1: 105, 101, 102, 8: 65, 107, 106, 15: 104, 18: 103, 28: 145},
		// 40
		{31, 90, 5: 31, 11: 93, 92, 94, 91, 16: 95, 96, 31: 141, 53: 140},
		{1: 105, 101, 102, 8: 65, 107, 106, 15: 104, 18: 103, 28: 100, 33: 99, 36: 136},
		{1: 23, 23, 23, 8: 23, 23, 33: 23, 46: 97},
		{1: 105, 101, 102, 8: 65, 107, 106, 15: 104, 18: 103, 28: 100, 33: 99, 36: 98},
		{6: 22, 47: 133},
		// 45
		{1: 105, 101, 102, 8: 65, 107, 106, 15: 104, 18: 103, 28: 132},
		{2: 109, 110, 19: 120, 121, 22: 122, 124, 125, 123},
		{1: 105, 8: 65, 107, 106, 15: 104, 18: 119},
		{1: 105, 8: 65, 107, 106, 15: 104, 18: 118},
		{9, 2: 9, 9, 9, 9, 9, 9, 19: 9, 9, 9, 9, 9, 9, 9, 113, 114},
		// 50
		{6, 2: 6, 6, 6, 6, 6, 6, 19: 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{3, 2: 3, 3, 3, 3, 3, 3, 19: 3, 3, 3, 3, 3, 3, 3, 3, 3},
		{2, 2: 2, 2, 2, 2, 2, 2, 19: 2, 2, 2, 2, 2, 2, 2, 2, 2},
		{1: 105, 101, 102, 8: 65, 107, 106, 15: 104, 18: 103, 28: 108},
		{2: 109, 110, 21: 111},
		// 55
		{1: 105, 8: 65, 107, 106, 15: 104, 18: 117},
		{1: 105, 8: 65, 107, 106, 15: 104, 18: 112},
		{1, 2: 1, 1, 1, 1, 1, 1, 19: 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{7, 2: 7, 7, 7, 7, 7, 7, 19: 7, 7, 7, 7, 7, 7, 7, 113, 114},
		{1: 105, 8: 65, 107, 106, 15: 116},
		// 60
		{1: 105, 8: 65, 107, 106, 15: 115},
		{4, 2: 4, 4, 4, 4, 4, 4, 19: 4, 4, 4, 4, 4, 4, 4, 4, 4},
		{5, 2: 5, 5, 5, 5, 5, 5, 19: 5, 5, 5, 5, 5, 5, 5, 5, 5},
		{8, 2: 8, 8, 8, 8, 8, 8, 19: 8, 8, 8, 8, 8, 8, 8, 113, 114},
		{10, 2: 10, 10, 10, 10, 10, 10, 19: 10, 10, 10, 10, 10, 10, 10, 113, 114},
		// 65
		{11, 2: 11, 11, 11, 11, 11, 11, 19: 11, 11, 11, 11, 11, 11, 11, 113, 114},
		{1: 105, 101, 102, 8: 65, 107, 106, 15: 104, 18: 103, 28: 131},
		{1: 105, 101, 102, 8: 65, 107, 106, 15: 104, 18: 103, 28: 130},
		{1: 105, 101, 102, 8: 65, 107, 106, 15: 104, 18: 103, 28: 129},
		{1: 105, 101, 102, 8: 65, 107, 106, 15: 104, 18: 103, 28: 128},
		// 70
		{1: 105, 101, 102, 8: 65, 107, 106, 15: 104, 18: 103, 28: 127},
		{1: 105, 101, 102, 8: 65, 107, 106, 15: 104, 18: 103, 28: 126},
		{2: 109, 110, 6: 12, 12},
		{2: 109, 110, 6: 13, 13},
		{2: 109, 110, 6: 14, 14},
		// 75
		{2: 109, 110, 6: 15, 15},
		{2: 109, 110, 6: 16, 16},
		{2: 109, 110, 6: 17, 17},
		{2: 109, 110, 6: 18, 18},
		{6: 134},
		// 80
		{31, 90, 4: 31, 31, 11: 93, 92, 94, 91, 16: 95, 96, 31: 135},
		{21, 4: 21, 21},
		{7: 25, 45: 137},
		{7: 138},
		{31, 90, 4: 31, 31, 11: 93, 92, 94, 91, 16: 95, 96, 31: 139},
		// 85
		{24, 4: 24, 24},
		{143, 5: 142},
		{20, 5: 20},
		{26, 4: 26, 26},
		{31, 90, 5: 31, 11: 93, 92, 94, 91, 16: 95, 96, 31: 144},
		// 90
		{19, 5: 19},
		{27, 2: 109, 110, 27, 27},
		{28, 4: 28, 28},
		{29, 4: 29, 29},
		{1: 105, 101, 102, 8: 65, 107, 106, 15: 104, 18: 103, 28: 149},
		// 95
		{30, 2: 109, 110, 30, 30},
		{34: 53},
	}
)

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyLexerEx interface {
	yyLexer
	Reduced(rule, state int, lval *yySymType) bool
}

func yySymName(c int) (s string) {
	x, ok := yyXLAT[c]
	if ok {
		return yySymNames[x]
	}

	return __yyfmt__.Sprintf("%d", c)
}

func yylex1(yylex yyLexer, lval *yySymType) (n int) {
	n = yylex.Lex(lval)
	if n <= 0 {
		n = yyEofCode
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("\nlex %s(%#x %d), PrettyString(lval.Token): %v\n", yySymName(n), n, n, PrettyString(lval.Token))
	}
	return n
}

func yyParse(yylex yyLexer) int {
	const yyError = 58

	yyEx, _ := yylex.(yyLexerEx)
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, 200)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yyerrok := func() {
		if yyDebug >= 2 {
			__yyfmt__.Printf("yyerrok()\n")
		}
		Errflag = 0
	}
	_ = yyerrok
	yystate := 0
	yychar := -1
	var yyxchar int
	var yyshift int
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
		var ok bool
		if yyxchar, ok = yyXLAT[yychar]; !ok {
			yyxchar = len(yySymNames) // > tab width
		}
	}
	if yyDebug >= 4 {
		var a []int
		for _, v := range yyS[:yyp+1] {
			a = append(a, v.yys)
		}
		__yyfmt__.Printf("state stack %v\n", a)
	}
	row := yyParseTab[yystate]
	yyn = 0
	if yyxchar < len(row) {
		if yyn = int(row[yyxchar]); yyn != 0 {
			yyn += yyTabOfs
		}
	}
	switch {
	case yyn > 0: // shift
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		yyshift = yyn
		if yyDebug >= 2 {
			__yyfmt__.Printf("shift, and goto state %d\n", yystate)
		}
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	case yyn < 0: // reduce
	case yystate == 1: // accept
		if yyDebug >= 2 {
			__yyfmt__.Println("accept")
		}
		goto ret0
	}

	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			if yyDebug >= 1 {
				__yyfmt__.Printf("no action for %s in state %d\n", yySymName(yychar), yystate)
			}
			msg, ok := yyXErrors[yyXError{yystate, yyxchar}]
			if !ok {
				msg, ok = yyXErrors[yyXError{yystate, -1}]
			}
			if !ok && yyshift != 0 {
				msg, ok = yyXErrors[yyXError{yyshift, yyxchar}]
			}
			if !ok {
				msg, ok = yyXErrors[yyXError{yyshift, -1}]
			}
			if !ok || msg == "" {
				msg = "syntax error"
			}
			yylex.Error(msg)
			Nerrs++
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				row := yyParseTab[yyS[yyp].yys]
				if yyError < len(row) {
					yyn = int(row[yyError]) + yyTabOfs
					if yyn > 0 { // hit
						if yyDebug >= 2 {
							__yyfmt__.Printf("error recovery found error shift in state %d\n", yyS[yyp].yys)
						}
						yystate = yyn /* simulate a shift of "error" */
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery failed\n")
			}
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yySymName(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}

			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	r := -yyn
	x0 := yyReductions[r]
	x, n := x0.xsym, x0.components
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= n
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	exState := yystate
	yystate = int(yyParseTab[yyS[yyp].yys][x]) + yyTabOfs
	/* reduction by production r */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce using rule %v (%s), and goto state %d\n", r, yySymNames[x], yystate)
	}

	switch r {
	case 1:
		{
			lx := yylex.(*lexer)
			lhs := &Program{
				Block: yyS[yypt-1].node.(*Block),
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			if lx.scope.Parent != nil {
				panic("internal error")
			}

			switch addr := lhs.Block.addr; {
			case addr == 0 || addr == 2:
				lx.code = []Instruction{}
			default:
				lx.enter.NVars = lhs.Block.VarsOpt.nvars()
				lx.enter.Token = lx.firstTok
				lx.jmp.Target = addr
			}
			lx.emit(&Leave{})
			lhs.Code = lx.code
			lx.program = lhs
		}
	case 2:
		{
			lx := yylex.(*lexer)
			lx.scope.addr = len(lx.code)
		}
	case 3:
		{
			lx := yylex.(*lexer)
			lhs := &Block{
				ConstsOpt:   yyS[yypt-4].node.(*ConstsOpt),
				VarsOpt:     yyS[yypt-3].node.(*VarsOpt),
				ProcListOpt: yyS[yypt-2].node.(*ProcListOpt),
				Statement:   yyS[yypt-0].node.(*Statement),
			}
			yyVAL.node = lhs
			lhs.addr = lx.scope.addr
		}
	case 4:
		{
			yyVAL.node = (*ConstsOpt)(nil)
		}
	case 5:
		{
			yyVAL.node = &ConstsOpt{
				Consts: yyS[yypt-0].node.(*Consts),
			}
		}
	case 6:
		{
			yyVAL.node = &Consts{
				Token:         yyS[yypt-2].Token,
				ConstSpecList: yyS[yypt-1].node.(*ConstSpecList).reverse(),
				Token2:        yyS[yypt-0].Token,
			}
		}
	case 7:
		{
			yyVAL.node = &ConstSpecList{
				ConstSpec: yyS[yypt-0].node.(*ConstSpec),
			}
		}
	case 8:
		{
			yyVAL.node = &ConstSpecList{
				Case:          1,
				ConstSpecList: yyS[yypt-2].node.(*ConstSpecList),
				Token:         yyS[yypt-1].Token,
				ConstSpec:     yyS[yypt-0].node.(*ConstSpec),
			}
		}
	case 9:
		{
			lx := yylex.(*lexer)
			lhs := &ConstSpec{
				Token:  yyS[yypt-2].Token,
				Token2: yyS[yypt-1].Token,
				Number: yyS[yypt-0].node.(*Number),
			}
			yyVAL.node = lhs
			lx.bind(lhs.Token, lhs)
		}
	case 10:
		{
			lx := yylex.(*lexer)
			lhs := &Number{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			t := lhs.Token
			n, err := strconv.ParseUint(string(t.S()), 10, 31)
			if err != nil {
				lx.report.ErrTok(t, "%s", err)
			}
			lhs.Value = int(n)
		}
	case 11:
		{
			yyVAL.node = (*VarsOpt)(nil)
		}
	case 12:
		{
			yyVAL.node = &VarsOpt{
				Vars: yyS[yypt-0].node.(*Vars),
			}
		}
	case 13:
		{
			lx := yylex.(*lexer)
			lhs := &Vars{
				Token:        yyS[yypt-2].Token,
				VariableList: yyS[yypt-1].node.(*VariableList).reverse(),
				Token2:       yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.nvars = lx.scope.nvar
		}
	case 14:
		{
			lx := yylex.(*lexer)
			lhs := &Variable{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lhs.index = lx.scope.nvar
			lx.scope.nvar++
			lx.bind(lhs.Token, lhs)
		}
	case 15:
		{
			yyVAL.node = &VariableList{
				Variable: yyS[yypt-0].node.(*Variable),
			}
		}
	case 16:
		{
			yyVAL.node = &VariableList{
				Case:         1,
				VariableList: yyS[yypt-2].node.(*VariableList),
				Token:        yyS[yypt-1].Token,
				Variable:     yyS[yypt-0].node.(*Variable),
			}
		}
	case 17:
		{
			yyVAL.node = (*ProcListOpt)(nil)
		}
	case 18:
		{
			yyVAL.node = &ProcListOpt{
				ProcList: yyS[yypt-0].node.(*ProcList).reverse(),
			}
		}
	case 19:
		{
			yyVAL.node = &ProcList{
				ProcSpec: yyS[yypt-0].node.(*ProcSpec),
			}
		}
	case 20:
		{
			yyVAL.node = &ProcList{
				Case:     1,
				ProcList: yyS[yypt-1].node.(*ProcList),
				ProcSpec: yyS[yypt-0].node.(*ProcSpec),
			}
		}
	case 21:
		{
			lx := yylex.(*lexer)
			enter := &Enter{Token: yyS[yypt-0].Token, NVars: -1}
			jmp := &Jmp{Target: -1}
			lx.bind(yyS[yypt-0].Token, &ProcSpec{addr: len(lx.code), enter: enter, jmp: jmp}) // Declare early.
			lx.emit(enter)
			lx.emit(jmp)
			enter.LNL = lx.newScope().lnl
		}
	case 22:
		{
			lx := yylex.(*lexer)
			lhs := &ProcSpec{
				Token:  yyS[yypt-5].Token,
				Token2: yyS[yypt-4].Token,
				Token3: yyS[yypt-2].Token,
				Block:  yyS[yypt-1].node.(*Block),
				Token4: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			lx.popScope(lhs.Token4)
			nm := lhs.Token2
			if ps, ok := lx.scope.Map[nm.Val].Node.(*ProcSpec); ok {
				ps.enter.NVars = lhs.Block.VarsOpt.nvars()
				ps.jmp.Target = lhs.Block.addr
				lhs.addr = ps.addr
				lhs.enter = ps.enter
				lhs.jmp = ps.jmp
				*ps = *lhs
				yyVAL.node = ps
			}
			lx.emit(&Leave{})
		}
	case 23:
		{
			yyVAL.node = (*Statement)(nil)
		}
	case 24:
		{
			lx := yylex.(*lexer)
			lhs := &Statement{
				Case:       1,
				Token:      yyS[yypt-2].Token,
				Token2:     yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			n, f := lx.mustVar(lhs.Token)
			if n == nil {
				break
			}

			x := n.(*Variable)
			lx.emit(&PopVar{Index: x.index, Frames: f, Token: x.Token})
		}
	case 25:
		{
			lx := yylex.(*lexer)
			lhs := &Statement{
				Case:   2,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			n, _ := lx.mustProc(lhs.Token2)
			if n == nil {
				break
			}

			x := n.(*ProcSpec)
			lx.emit(&Call{Target: x.addr, Token: lhs.Token})
		}
	case 26:
		{
			lx := yylex.(*lexer)
			lhs := &Statement{
				Case:   3,
				Token:  yyS[yypt-1].Token,
				Token2: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			n, f := lx.mustVar(lhs.Token2)
			if n == nil {
				break
			}

			x := n.(*Variable)
			lx.emit(&Read{Index: x.index, Frames: f, Token: x.Token})
		}
	case 27:
		{
			lx := yylex.(*lexer)
			lhs := &Statement{
				Case:       4,
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			lx.emit(&Write{Token: lhs.Token})
		}
	case 28:
		{
			yyVAL.node = &Statement{
				Case:          5,
				Token:         yyS[yypt-2].Token,
				StatementList: yyS[yypt-1].node.(*StatementList).reverse(),
				Token2:        yyS[yypt-0].Token,
			}
		}
	case 29:
		{
			lx := yylex.(*lexer)
			ifEnd := &JmpZero{}
			lx.scope.ifEnd = ifEnd
			lx.emit(ifEnd)
		}
	case 30:
		{
			lx := yylex.(*lexer)
			yyVAL.node = &Statement{
				Case:      6,
				Token:     yyS[yypt-4].Token,
				Condition: yyS[yypt-3].node.(*Condition),
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].node.(*Statement),
			}
			lx.scope.ifEnd.Target = len(lx.code)
		}
	case 31:
		{
			lx := yylex.(*lexer)
			lx.scope.while0 = len(lx.code)
		}
	case 32:
		{
			lx := yylex.(*lexer)
			lx.scope.whileEnd = &JmpZero{}
			lx.emit(lx.scope.whileEnd)
		}
	case 33:
		{
			lx := yylex.(*lexer)
			yyVAL.node = &Statement{
				Case:      7,
				Token:     yyS[yypt-5].Token,
				Condition: yyS[yypt-3].node.(*Condition),
				Token2:    yyS[yypt-1].Token,
				Statement: yyS[yypt-0].node.(*Statement),
			}
			lx.emit(&Jmp{Target: lx.scope.while0})
			lx.scope.whileEnd.Target = len(lx.code)
		}
	case 34:
		{
			yyVAL.node = &StatementList{
				Statement: yyS[yypt-0].node.(*Statement),
			}
		}
	case 35:
		{
			yyVAL.node = &StatementList{
				Case:          1,
				StatementList: yyS[yypt-2].node.(*StatementList),
				Token:         yyS[yypt-1].Token,
				Statement:     yyS[yypt-0].node.(*Statement),
			}
		}
	case 36:
		{
			lx := yylex.(*lexer)
			lhs := &Condition{
				Token:      yyS[yypt-1].Token,
				Expression: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			lx.emit(&TestOdd{Token: lhs.Token})
		}
	case 37:
		{
			lx := yylex.(*lexer)
			lhs := &Condition{
				Case:        1,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			lx.emit(&TestEQ{Token: lhs.Token})
		}
	case 38:
		{
			lx := yylex.(*lexer)
			lhs := &Condition{
				Case:        2,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			lx.emit(&TestMod{Token: lhs.Token})
		}
	case 39:
		{
			lx := yylex.(*lexer)
			lhs := &Condition{
				Case:        3,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			lx.emit(&TestLT{Token: lhs.Token})
		}
	case 40:
		{
			lx := yylex.(*lexer)
			lhs := &Condition{
				Case:        4,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			lx.emit(&TestLEQ{Token: lhs.Token})
		}
	case 41:
		{
			lx := yylex.(*lexer)
			lhs := &Condition{
				Case:        5,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			lx.emit(&TestGT{Token: lhs.Token})
		}
	case 42:
		{
			lx := yylex.(*lexer)
			lhs := &Condition{
				Case:        6,
				Expression:  yyS[yypt-2].node.(*Expression),
				Token:       yyS[yypt-1].Token,
				Expression2: yyS[yypt-0].node.(*Expression),
			}
			yyVAL.node = lhs
			lx.emit(&TestGEQ{Token: lhs.Token})
		}
	case 43:
		{
			yyVAL.node = &Expression{
				Token: yyS[yypt-1].Token,
				Term:  yyS[yypt-0].node.(*Term),
			}
		}
	case 44:
		{
			lx := yylex.(*lexer)
			lhs := &Expression{
				Case:  1,
				Token: yyS[yypt-1].Token,
				Term:  yyS[yypt-0].node.(*Term),
			}
			yyVAL.node = lhs
			lx.emit(&Neg{Token: lhs.Token})
		}
	case 45:
		{
			yyVAL.node = &Expression{
				Case: 2,
				Term: yyS[yypt-0].node.(*Term),
			}
		}
	case 46:
		{
			lx := yylex.(*lexer)
			lhs := &Expression{
				Case:       3,
				Expression: yyS[yypt-2].node.(*Expression),
				Token:      yyS[yypt-1].Token,
				Term:       yyS[yypt-0].node.(*Term),
			}
			yyVAL.node = lhs
			lx.emit(&Add{Token: lhs.Token})
		}
	case 47:
		{
			lx := yylex.(*lexer)
			lhs := &Expression{
				Case:       4,
				Expression: yyS[yypt-2].node.(*Expression),
				Token:      yyS[yypt-1].Token,
				Term:       yyS[yypt-0].node.(*Term),
			}
			yyVAL.node = lhs
			lx.emit(&Sub{Token: lhs.Token})
		}
	case 48:
		{
			yyVAL.node = &Term{
				Factor: yyS[yypt-0].node.(*Factor),
			}
		}
	case 49:
		{
			lx := yylex.(*lexer)
			lhs := &Term{
				Case:   1,
				Term:   yyS[yypt-2].node.(*Term),
				Token:  yyS[yypt-1].Token,
				Factor: yyS[yypt-0].node.(*Factor),
			}
			yyVAL.node = lhs
			lx.emit(&Mul{Token: lhs.Token})
		}
	case 50:
		{
			lx := yylex.(*lexer)
			lhs := &Term{
				Case:   2,
				Term:   yyS[yypt-2].node.(*Term),
				Token:  yyS[yypt-1].Token,
				Factor: yyS[yypt-0].node.(*Factor),
			}
			yyVAL.node = lhs
			lx.emit(&Div{Token: lhs.Token})
		}
	case 51:
		{
			lx := yylex.(*lexer)
			lhs := &Factor{
				Token: yyS[yypt-0].Token,
			}
			yyVAL.node = lhs
			n, f := lx.mustVarOrConst(lhs.Token)
			if n == nil {
				break
			}

			switch x := n.(type) {
			case *ConstSpec:
				lx.emit(&PushConst{-1, x.Number.Value, x.Token})
			case *Variable:
				lx.emit(&PushVar{-1, x.index, f, x.Token})
			default:
				panic("internal error")
			}
		}
	case 52:
		{
			lx := yylex.(*lexer)
			lhs := &Factor{
				Case:   1,
				Number: yyS[yypt-0].node.(*Number),
			}
			yyVAL.node = lhs
			lx.emit(&PushConst{Value: lhs.Number.Value, Token: lhs.Number.Token})
		}
	case 53:
		{
			yyVAL.node = &Factor{
				Case:       2,
				Token:      yyS[yypt-2].Token,
				Expression: yyS[yypt-1].node.(*Expression),
				Token2:     yyS[yypt-0].Token,
			}
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}
