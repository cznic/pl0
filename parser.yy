%{
// Copyright 2015 The PL0 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pl0

import (
	"strconv"

	"github.com/cznic/xc"
)
%}

/*

Based on the EBNF grammar found in [0]:

program = block "." .

block = [ "const" ident "=" number { "," ident "=" number} ";" ]
        [ "var" ident { "," ident } ";" ]
        { "procedure" ident ";" block ";" } statement .

statement = [ ident ":=" expression | "call" ident 
              | "?" ident | "!" expression 
              | "begin" statement { ";" statement } "end" 
              | "if" condition "then" statement 
              | "while" condition "do" statement ] .

condition = "odd" expression |
            expression ( "=" | "#" | "<" | "<=" | ">" | ">=" ) expression .

expression = [ "+" | "-" ] term { ( "+" | "-" ) term } .

term = factor { ( "*" | "/" ) factor } .

factor = ident | number | "(" expression ")".

*/

%union	{
	node	Node
	Token	xc.Token
}

%token
	/*yy:token "%c" */	IDENT	"identifier"
	/*yy:token "%d" */	NUMBER	"integer literal"

	ASSIGN		":="
	BEGIN		"BEGIN"
	CALL		"CALL"
	CONST		"CONST"
	DO		"DO"
	END		"END"
	GEQ		">="
	IF		"IF"
	LEQ		"<="
	ODD		"ODD"
	PROCEDURE	"PROCEDURE"
	THEN		"THEN"
	VAR		"VAR"
	WHILE		"WHILE"

%type	<node>
	Block		"block"
	Condition	"conditional expression"
	ConstSpec	"constant specification"
	ConstSpecList	"list of constant declarations"
	Consts		"constant declarations"
	ConstsOpt	"optional constant declarations"
	Expression	"expression"
	Factor		"expression factor"
	Number		"number"
	ProcList	"procedure declarations"
	ProcListOpt	"optional procedure declarations"
	ProcSpec	"procedure defintion"
	Program		"program"
	Statement	"statement"
	StatementList	"statement list"
	Term		"expression term"
	Variable	"variable name"
	VariableList	"variables list"
	Vars		"variable declarations"
	VarsOpt		"optional variable declarations"

%%

//yy:field	Code	[]Instruction
Program:
	Block '.'
	{
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

//yy:example	"CALL x."
//yy:field	addr	int
Block:
	ConstsOpt VarsOpt ProcListOpt
	{
		lx.scope.addr = len(lx.code)
	}
	Statement
	{
		lhs.addr = lx.scope.addr
	}

ConstsOpt:
	/* empty */ {}
|	Consts

Consts:
	"CONST" ConstSpecList ';'

ConstSpecList:
	ConstSpec
|	ConstSpecList ',' ConstSpec

ConstSpec:
	IDENT '=' Number
	{
		lx.bind(lhs.Token, lhs)
	}

//yy:field	Value	int	// Numeric value of the token.
Number:
	NUMBER
	{
		t := lhs.Token
		n, err := strconv.ParseUint(string(t.S()), 10, 31)
		if err != nil {
			lx.report.ErrTok(t, "%s", err)
		}
		lhs.Value = int(n)
	}

VarsOpt:
	/* empty */ {}
|	Vars

//yy:field	nvars	int
Vars:
	"VAR" VariableList ';'
	{
		lhs.nvars = lx.scope.nvar
	}

//yy:field	index	int
Variable:
	IDENT
	{
		lhs.index = lx.scope.nvar
		lx.scope.nvar++
		lx.bind(lhs.Token, lhs)
	}

VariableList:
	Variable
|	VariableList ',' Variable

ProcListOpt:
	/* empty */ {}
|	ProcList

ProcList:
	ProcSpec
|	ProcList ProcSpec

//yy:field	addr	int
//yy:field	enter	*Enter
//yy:field	jmp	*Jmp
ProcSpec:
	"PROCEDURE" IDENT
	{
		enter := &Enter{Token: $2, NVars: -1}
		jmp := &Jmp{Target: -1}
		lx.bind($2, &ProcSpec{addr: len(lx.code), enter: enter, jmp: jmp}) // Declare early.
		lx.emit(enter)
		lx.emit(jmp)
		enter.LNL = lx.newScope().lnl
	}
	';' Block ';'
	{
		lx.popScope(lhs.Token4)
		nm := lhs.Token2
		if ps, ok := lx.scope.Map[nm.Val].Node.(*ProcSpec); ok {
			ps.enter.NVars = lhs.Block.VarsOpt.nvars()
			ps.jmp.Target = lhs.Block.addr
			lhs.addr = ps.addr
			lhs.enter = ps.enter
			lhs.jmp = ps.jmp
			*ps = *lhs
			$$ = ps
		}
		lx.emit(&Leave{})
	}

Statement:
	/* empty */ {}
|	IDENT ":=" Expression 
	{
		n, f := lx.mustVar(lhs.Token)
		if n == nil {
			break
		}

		x := n.(*Variable)
		lx.emit(&PopVar{Index: x.index, Frames: f, Token: x.Token})
	}
|	"CALL" IDENT 
	{
		n, _ := lx.mustProc(lhs.Token2)
		if n == nil {
			break
		}

		x := n.(*ProcSpec)
		lx.emit(&Call{Target: x.addr, Token: lhs.Token})
	}
|	'?' IDENT
	{
		n, f := lx.mustVar(lhs.Token2)
		if n == nil {
			break
		}

		x := n.(*Variable)
		lx.emit(&Read{Index: x.index, Frames: f, Token: x.Token})
	}
|	'!' Expression 
	{
		lx.emit(&Write{Token: lhs.Token})
	}
|	"BEGIN" StatementList "END"
|	"IF" Condition
	{
		ifEnd := &JmpZero{}
		lx.scope.ifEnd = ifEnd
		lx.emit(ifEnd)
	}
	"THEN" Statement 
	{
		lx.scope.ifEnd.Target = len(lx.code)
	}
|	"WHILE"
	{
		lx.scope.while0 = len(lx.code)
	}
	Condition
	{
		lx.scope.whileEnd = &JmpZero{}
		lx.emit(lx.scope.whileEnd)
	}
	"DO" Statement
	{
		lx.emit(&Jmp{Target: lx.scope.while0})
		lx.scope.whileEnd.Target = len(lx.code)
	}

//yy:example	"BEGIN i := 42 END"
StatementList:
	Statement
|	StatementList ';' Statement

Condition:
	"ODD" Expression
	{
		lx.emit(&TestOdd{Token: lhs.Token})
	}
|	Expression '=' Expression
	{
		lx.emit(&TestEQ{Token: lhs.Token})
	}
|	Expression '#' Expression
	{
		lx.emit(&TestMod{Token: lhs.Token})
	}
|	Expression '<' Expression
	{
		lx.emit(&TestLT{Token: lhs.Token})
	}
|	Expression "<=" Expression
	{
		lx.emit(&TestLEQ{Token: lhs.Token})
	}
|	Expression '>' Expression
	{
		lx.emit(&TestGT{Token: lhs.Token})
	}
|	Expression ">=" Expression
	{
		lx.emit(&TestGEQ{Token: lhs.Token})
	}

Expression:
	'+' Term
|	'-' Term
	{
		lx.emit(&Neg{Token: lhs.Token})
	}
|	Term
|	Expression '+' Term
	{
		lx.emit(&Add{Token: lhs.Token})
	}
|	Expression '-' Term
	{
		lx.emit(&Sub{Token: lhs.Token})
	}

Term:
	Factor
|	Term '*' Factor
	{
		lx.emit(&Mul{Token: lhs.Token})
	}
|	Term '/' Factor
	{
		lx.emit(&Div{Token: lhs.Token})
	}

Factor:
	IDENT
	{
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
|	Number
	{
		lx.emit(&PushConst{Value: lhs.Number.Value, Token: lhs.Number.Token})
	}
|	'(' Expression ')'
