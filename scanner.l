%{
// Copyright 2015 The PL0 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Based on the EBNF grammar found in [0].
%}

%yyc c
%yyn c = l.Next()
%yym l.Mark()

%{
package pl0

func (l *lexer) scan() (r int) {
	c := l.Enter()
%}

ident-first	[a-zA-Z_]
ident-next	{ident-first}|{digit}
digit		[0-9]
ident		{ident-first}{ident-next}*
number		{digit}+

%%
		c = l.Rule0()

[ \n\t\f\v]+	|
"//".*

":="		return ASSIGN
"<="		return LEQ
">="		return GEQ

BEGIN		return BEGIN
CALL		return CALL
CONST		return CONST
DO		return DO
END		return END
IF		return IF
ODD		return ODD
PROCEDURE	return PROCEDURE
THEN		return THEN
VAR		return VAR
WHILE		return WHILE

{ident}		return IDENT
{number}	return NUMBER

%%
	if c, ok := l.Abort(); ok {
		return c
	}
        
	goto yyAction
}
