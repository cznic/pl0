// CAUTION: Generated file - DO NOT EDIT.

// Copyright 2015 The PL0 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on the EBNF grammar found in [0].

package pl0

func (l *lexer) scan() (r int) {
	c := l.Enter()

yystate0:
	yyrule := -1
	_ = yyrule
	c = l.Rule0()

	goto yystart1

	goto yystate0 // silence unused label error
	goto yyAction // silence unused label error
yyAction:
	switch yyrule {
	case 1:
		goto yyrule1
	case 2:
		goto yyrule2
	case 3:
		goto yyrule3
	case 4:
		goto yyrule4
	case 5:
		goto yyrule5
	case 6:
		goto yyrule6
	case 7:
		goto yyrule7
	case 8:
		goto yyrule8
	case 9:
		goto yyrule9
	case 10:
		goto yyrule10
	case 11:
		goto yyrule11
	case 12:
		goto yyrule12
	case 13:
		goto yyrule13
	case 14:
		goto yyrule14
	case 15:
		goto yyrule15
	case 16:
		goto yyrule16
	case 17:
		goto yyrule17
	case 18:
		goto yyrule18
	}
	goto yystate1 // silence unused label error
yystate1:
	c = l.Next()
yystart1:
	switch {
	default:
		goto yyabort
	case c == '/':
		goto yystate3
	case c == ':':
		goto yystate6
	case c == '<':
		goto yystate8
	case c == '>':
		goto yystate10
	case c == 'A' || c >= 'F' && c <= 'H' || c >= 'J' && c <= 'N' || c >= 'Q' && c <= 'S' || c == 'U' || c >= 'X' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	case c == 'B':
		goto yystate13
	case c == 'C':
		goto yystate18
	case c == 'D':
		goto yystate26
	case c == 'E':
		goto yystate28
	case c == 'I':
		goto yystate31
	case c == 'O':
		goto yystate33
	case c == 'P':
		goto yystate36
	case c == 'T':
		goto yystate45
	case c == 'V':
		goto yystate49
	case c == 'W':
		goto yystate52
	case c >= '0' && c <= '9':
		goto yystate5
	case c >= '\t' && c <= '\f' || c == ' ':
		goto yystate2
	}

yystate2:
	c = l.Next()
	yyrule = 1
	l.Mark()
	switch {
	default:
		goto yyrule1
	case c >= '\t' && c <= '\f' || c == ' ':
		goto yystate2
	}

yystate3:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '/':
		goto yystate4
	}

yystate4:
	c = l.Next()
	yyrule = 2
	l.Mark()
	switch {
	default:
		goto yyrule2
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate4
	}

yystate5:
	c = l.Next()
	yyrule = 18
	l.Mark()
	switch {
	default:
		goto yyrule18
	case c >= '0' && c <= '9':
		goto yystate5
	}

yystate6:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate7
	}

yystate7:
	c = l.Next()
	yyrule = 3
	l.Mark()
	goto yyrule3

yystate8:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate9
	}

yystate9:
	c = l.Next()
	yyrule = 4
	l.Mark()
	goto yyrule4

yystate10:
	c = l.Next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate11
	}

yystate11:
	c = l.Next()
	yyrule = 5
	l.Mark()
	goto yyrule5

yystate12:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate13:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == 'E':
		goto yystate14
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate14:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == 'G':
		goto yystate15
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'H' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate15:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == 'I':
		goto yystate16
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate16:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == 'N':
		goto yystate17
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate17:
	c = l.Next()
	yyrule = 6
	l.Mark()
	switch {
	default:
		goto yyrule6
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate18:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == 'A':
		goto yystate19
	case c == 'O':
		goto yystate22
	case c >= '0' && c <= '9' || c >= 'B' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate19:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == 'L':
		goto yystate20
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate20:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == 'L':
		goto yystate21
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate21:
	c = l.Next()
	yyrule = 7
	l.Mark()
	switch {
	default:
		goto yyrule7
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate22:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == 'N':
		goto yystate23
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate23:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == 'S':
		goto yystate24
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'R' || c >= 'T' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate24:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == 'T':
		goto yystate25
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'S' || c >= 'U' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate25:
	c = l.Next()
	yyrule = 8
	l.Mark()
	switch {
	default:
		goto yyrule8
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate26:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == 'O':
		goto yystate27
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate27:
	c = l.Next()
	yyrule = 9
	l.Mark()
	switch {
	default:
		goto yyrule9
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate28:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == 'N':
		goto yystate29
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate29:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == 'D':
		goto yystate30
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate30:
	c = l.Next()
	yyrule = 10
	l.Mark()
	switch {
	default:
		goto yyrule10
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate31:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == 'F':
		goto yystate32
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'E' || c >= 'G' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate32:
	c = l.Next()
	yyrule = 11
	l.Mark()
	switch {
	default:
		goto yyrule11
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate33:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == 'D':
		goto yystate34
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate34:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == 'D':
		goto yystate35
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate35:
	c = l.Next()
	yyrule = 12
	l.Mark()
	switch {
	default:
		goto yyrule12
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate36:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == 'R':
		goto yystate37
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate37:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == 'O':
		goto yystate38
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'N' || c >= 'P' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate38:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == 'C':
		goto yystate39
	case c >= '0' && c <= '9' || c == 'A' || c == 'B' || c >= 'D' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate39:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == 'E':
		goto yystate40
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate40:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == 'D':
		goto yystate41
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'C' || c >= 'E' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate41:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == 'U':
		goto yystate42
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'T' || c >= 'V' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate42:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == 'R':
		goto yystate43
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate43:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == 'E':
		goto yystate44
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate44:
	c = l.Next()
	yyrule = 13
	l.Mark()
	switch {
	default:
		goto yyrule13
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate45:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == 'H':
		goto yystate46
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'G' || c >= 'I' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate46:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == 'E':
		goto yystate47
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate47:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == 'N':
		goto yystate48
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'M' || c >= 'O' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate48:
	c = l.Next()
	yyrule = 14
	l.Mark()
	switch {
	default:
		goto yyrule14
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate49:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == 'A':
		goto yystate50
	case c >= '0' && c <= '9' || c >= 'B' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate50:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == 'R':
		goto yystate51
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Q' || c >= 'S' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate51:
	c = l.Next()
	yyrule = 15
	l.Mark()
	switch {
	default:
		goto yyrule15
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate52:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == 'H':
		goto yystate53
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'G' || c >= 'I' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate53:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == 'I':
		goto yystate54
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'H' || c >= 'J' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate54:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == 'L':
		goto yystate55
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate55:
	c = l.Next()
	yyrule = 17
	l.Mark()
	switch {
	default:
		goto yyrule17
	case c == 'E':
		goto yystate56
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'D' || c >= 'F' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yystate56:
	c = l.Next()
	yyrule = 16
	l.Mark()
	switch {
	default:
		goto yyrule16
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate12
	}

yyrule1: // [ \n\t\f\v]+
yyrule2: // "//".*

	goto yystate0
yyrule3: // ":="
	{
		return ASSIGN
	}
yyrule4: // "<="
	{
		return LEQ
	}
yyrule5: // ">="
	{
		return GEQ
	}
yyrule6: // BEGIN
	{
		return BEGIN
	}
yyrule7: // CALL
	{
		return CALL
	}
yyrule8: // CONST
	{
		return CONST
	}
yyrule9: // DO
	{
		return DO
	}
yyrule10: // END
	{
		return END
	}
yyrule11: // IF
	{
		return IF
	}
yyrule12: // ODD
	{
		return ODD
	}
yyrule13: // PROCEDURE
	{
		return PROCEDURE
	}
yyrule14: // THEN
	{
		return THEN
	}
yyrule15: // VAR
	{
		return VAR
	}
yyrule16: // WHILE
	{
		return WHILE
	}
yyrule17: // {ident}
	{
		return IDENT
	}
yyrule18: // {number}
	{
		return NUMBER
	}
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized
	if c, ok := l.Abort(); ok {
		return c
	}

	goto yyAction
}
