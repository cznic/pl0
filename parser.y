// CAUTION: Generated by yy - DO NOT EDIT.

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

%union {
	node  Node
	Token xc.Token
}

%token	<Token>
	'!'
	'#'
	'('
	')'
	'*'
	'+'
	','
	'-'
	'.'
	'/'
	';'
	'<'
	'='
	'>'
	'?'
	ASSIGN     ":="
	BEGIN      "BEGIN"
	CALL       "CALL"
	CONST      "CONST"
	DO         "DO"
	END        "END"
	GEQ        ">="
	IDENT      "identifier"
	IF         "IF"
	LEQ        "<="
	NUMBER     "integer literal"
	ODD        "ODD"
	PROCEDURE  "PROCEDURE"
	THEN       "THEN"
	VAR        "VAR"
	WHILE      "WHILE"

%type	<node>
	Block          "block"
	Condition      "conditional expression"
	ConstSpec      "constant specification"
	ConstSpecList  "list of constant declarations"
	Consts         "constant declarations"
	ConstsOpt      "optional constant declarations"
	Expression     "expression"
	Factor         "expression factor"
	Number         "number"
	ProcList       "procedure declarations"
	ProcListOpt    "optional procedure declarations"
	ProcSpec       "procedure defintion"
	Program        "program"
	Statement      "statement"
	StatementList  "statement list"
	Term           "expression term"
	Variable       "variable name"
	VariableList   "variables list"
	Vars           "variable declarations"
	VarsOpt        "optional variable declarations"

%start Program

%%

Program:
	Block '.'
	{
		lx := yylex.(*lexer)
		lhs := &Program{
			Block:  $1.(*Block),
			Token:  $2,
		}
		$$ = lhs
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

Block:
	ConstsOpt VarsOpt ProcListOpt
	{
		lx := yylex.(*lexer)
		lx.scope.addr = len(lx.code)
	}
	Statement
	{
		lx := yylex.(*lexer)
		lhs := &Block{
			ConstsOpt:    $1.(*ConstsOpt),
			VarsOpt:      $2.(*VarsOpt),
			ProcListOpt:  $3.(*ProcListOpt),
			Statement:    $5.(*Statement),
		}
		$$ = lhs
		lhs.addr = lx.scope.addr
	}

ConstsOpt:
	/* empty */
	{
		$$ = (*ConstsOpt)(nil)
		}
|	Consts
	{
		$$ = &ConstsOpt{
			Consts:  $1.(*Consts),
		}
	}

Consts:
	"CONST" ConstSpecList ';'
	{
		$$ = &Consts{
			Token:          $1,
			ConstSpecList:  $2.(*ConstSpecList).reverse(),
			Token2:         $3,
		}
	}

ConstSpecList:
	ConstSpec
	{
		$$ = &ConstSpecList{
			ConstSpec:  $1.(*ConstSpec),
		}
	}
|	ConstSpecList ',' ConstSpec
	{
		$$ = &ConstSpecList{
			Case:           1,
			ConstSpecList:  $1.(*ConstSpecList),
			Token:          $2,
			ConstSpec:      $3.(*ConstSpec),
		}
	}

ConstSpec:
	IDENT '=' Number
	{
		lx := yylex.(*lexer)
		lhs := &ConstSpec{
			Token:   $1,
			Token2:  $2,
			Number:  $3.(*Number),
		}
		$$ = lhs
		lx.bind(lhs.Token, lhs)
	}

Number:
	NUMBER
	{
		lx := yylex.(*lexer)
		lhs := &Number{
			Token:  $1,
		}
		$$ = lhs
		t := lhs.Token
		n, err := strconv.ParseUint(string(t.S()), 10, 31)
		if err != nil {
			lx.report.ErrTok(t, "%s", err)
		}
		lhs.Value = int(n)
	}

VarsOpt:
	/* empty */
	{
		$$ = (*VarsOpt)(nil)
		}
|	Vars
	{
		$$ = &VarsOpt{
			Vars:  $1.(*Vars),
		}
	}

Vars:
	"VAR" VariableList ';'
	{
		lx := yylex.(*lexer)
		lhs := &Vars{
			Token:         $1,
			VariableList:  $2.(*VariableList).reverse(),
			Token2:        $3,
		}
		$$ = lhs
		lhs.nvars = lx.scope.nvar
	}

Variable:
	IDENT
	{
		lx := yylex.(*lexer)
		lhs := &Variable{
			Token:  $1,
		}
		$$ = lhs
		lhs.index = lx.scope.nvar
		lx.scope.nvar++
		lx.bind(lhs.Token, lhs)
	}

VariableList:
	Variable
	{
		$$ = &VariableList{
			Variable:  $1.(*Variable),
		}
	}
|	VariableList ',' Variable
	{
		$$ = &VariableList{
			Case:          1,
			VariableList:  $1.(*VariableList),
			Token:         $2,
			Variable:      $3.(*Variable),
		}
	}

ProcListOpt:
	/* empty */
	{
		$$ = (*ProcListOpt)(nil)
		}
|	ProcList
	{
		$$ = &ProcListOpt{
			ProcList:  $1.(*ProcList).reverse(),
		}
	}

ProcList:
	ProcSpec
	{
		$$ = &ProcList{
			ProcSpec:  $1.(*ProcSpec),
		}
	}
|	ProcList ProcSpec
	{
		$$ = &ProcList{
			Case:      1,
			ProcList:  $1.(*ProcList),
			ProcSpec:  $2.(*ProcSpec),
		}
	}

ProcSpec:
	"PROCEDURE" IDENT
	{
		lx := yylex.(*lexer)
		enter := &Enter{Token: $2, NVars: -1}
		jmp := &Jmp{Target: -1}
		lx.bind($2, &ProcSpec{addr: len(lx.code), enter: enter, jmp: jmp}) // Declare early.
		lx.emit(enter)
		lx.emit(jmp)
		enter.LNL = lx.newScope().lnl
	}
	';' Block ';'
	{
		lx := yylex.(*lexer)
		lhs := &ProcSpec{
			Token:   $1,
			Token2:  $2,
			Token3:  $4,
			Block:   $5.(*Block),
			Token4:  $6,
		}
		$$ = lhs
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
	/* empty */
	{
		$$ = (*Statement)(nil)
		}
|	IDENT ":=" Expression
	{
		lx := yylex.(*lexer)
		lhs := &Statement{
			Case:        1,
			Token:       $1,
			Token2:      $2,
			Expression:  $3.(*Expression),
		}
		$$ = lhs
		n, f := lx.mustVar(lhs.Token)
		if n == nil {
			break
		}

		x := n.(*Variable)
		lx.emit(&PopVar{Index: x.index, Frames: f, Token: x.Token})
	}
|	"CALL" IDENT
	{
		lx := yylex.(*lexer)
		lhs := &Statement{
			Case:    2,
			Token:   $1,
			Token2:  $2,
		}
		$$ = lhs
		n, _ := lx.mustProc(lhs.Token2)
		if n == nil {
			break
		}

		x := n.(*ProcSpec)
		lx.emit(&Call{Target: x.addr, Token: lhs.Token})
	}
|	'?' IDENT
	{
		lx := yylex.(*lexer)
		lhs := &Statement{
			Case:    3,
			Token:   $1,
			Token2:  $2,
		}
		$$ = lhs
		n, f := lx.mustVar(lhs.Token2)
		if n == nil {
			break
		}

		x := n.(*Variable)
		lx.emit(&Read{Index: x.index, Frames: f, Token: x.Token})
	}
|	'!' Expression
	{
		lx := yylex.(*lexer)
		lhs := &Statement{
			Case:        4,
			Token:       $1,
			Expression:  $2.(*Expression),
		}
		$$ = lhs
		lx.emit(&Write{Token: lhs.Token})
	}
|	"BEGIN" StatementList "END"
	{
		$$ = &Statement{
			Case:           5,
			Token:          $1,
			StatementList:  $2.(*StatementList).reverse(),
			Token2:         $3,
		}
	}
|	"IF" Condition
	{
		lx := yylex.(*lexer)
		ifEnd := &JmpZero{}
		lx.scope.ifEnd = ifEnd
		lx.emit(ifEnd)
	}
	"THEN" Statement
	{
		lx := yylex.(*lexer)
		$$ = &Statement{
			Case:       6,
			Token:      $1,
			Condition:  $2.(*Condition),
			Token2:     $4,
			Statement:  $5.(*Statement),
		}
		lx.scope.ifEnd.Target = len(lx.code)
	}
|	"WHILE"
	{
		lx := yylex.(*lexer)
		lx.scope.while0 = len(lx.code)
	}
	Condition
	{
		lx := yylex.(*lexer)
		lx.scope.whileEnd = &JmpZero{}
		lx.emit(lx.scope.whileEnd)
	}
	"DO" Statement
	{
		lx := yylex.(*lexer)
		$$ = &Statement{
			Case:       7,
			Token:      $1,
			Condition:  $3.(*Condition),
			Token2:     $5,
			Statement:  $6.(*Statement),
		}
		lx.emit(&Jmp{Target: lx.scope.while0})
		lx.scope.whileEnd.Target = len(lx.code)
	}

StatementList:
	Statement
	{
		$$ = &StatementList{
			Statement:  $1.(*Statement),
		}
	}
|	StatementList ';' Statement
	{
		$$ = &StatementList{
			Case:           1,
			StatementList:  $1.(*StatementList),
			Token:          $2,
			Statement:      $3.(*Statement),
		}
	}

Condition:
	"ODD" Expression
	{
		lx := yylex.(*lexer)
		lhs := &Condition{
			Token:       $1,
			Expression:  $2.(*Expression),
		}
		$$ = lhs
		lx.emit(&TestOdd{Token: lhs.Token})
	}
|	Expression '=' Expression
	{
		lx := yylex.(*lexer)
		lhs := &Condition{
			Case:         1,
			Expression:   $1.(*Expression),
			Token:        $2,
			Expression2:  $3.(*Expression),
		}
		$$ = lhs
		lx.emit(&TestEQ{Token: lhs.Token})
	}
|	Expression '#' Expression
	{
		lx := yylex.(*lexer)
		lhs := &Condition{
			Case:         2,
			Expression:   $1.(*Expression),
			Token:        $2,
			Expression2:  $3.(*Expression),
		}
		$$ = lhs
		lx.emit(&TestMod{Token: lhs.Token})
	}
|	Expression '<' Expression
	{
		lx := yylex.(*lexer)
		lhs := &Condition{
			Case:         3,
			Expression:   $1.(*Expression),
			Token:        $2,
			Expression2:  $3.(*Expression),
		}
		$$ = lhs
		lx.emit(&TestLT{Token: lhs.Token})
	}
|	Expression "<=" Expression
	{
		lx := yylex.(*lexer)
		lhs := &Condition{
			Case:         4,
			Expression:   $1.(*Expression),
			Token:        $2,
			Expression2:  $3.(*Expression),
		}
		$$ = lhs
		lx.emit(&TestLEQ{Token: lhs.Token})
	}
|	Expression '>' Expression
	{
		lx := yylex.(*lexer)
		lhs := &Condition{
			Case:         5,
			Expression:   $1.(*Expression),
			Token:        $2,
			Expression2:  $3.(*Expression),
		}
		$$ = lhs
		lx.emit(&TestGT{Token: lhs.Token})
	}
|	Expression ">=" Expression
	{
		lx := yylex.(*lexer)
		lhs := &Condition{
			Case:         6,
			Expression:   $1.(*Expression),
			Token:        $2,
			Expression2:  $3.(*Expression),
		}
		$$ = lhs
		lx.emit(&TestGEQ{Token: lhs.Token})
	}

Expression:
	'+' Term
	{
		$$ = &Expression{
			Token:  $1,
			Term:   $2.(*Term),
		}
	}
|	'-' Term
	{
		lx := yylex.(*lexer)
		lhs := &Expression{
			Case:   1,
			Token:  $1,
			Term:   $2.(*Term),
		}
		$$ = lhs
		lx.emit(&Neg{Token: lhs.Token})
	}
|	Term
	{
		$$ = &Expression{
			Case:  2,
			Term:  $1.(*Term),
		}
	}
|	Expression '+' Term
	{
		lx := yylex.(*lexer)
		lhs := &Expression{
			Case:        3,
			Expression:  $1.(*Expression),
			Token:       $2,
			Term:        $3.(*Term),
		}
		$$ = lhs
		lx.emit(&Add{Token: lhs.Token})
	}
|	Expression '-' Term
	{
		lx := yylex.(*lexer)
		lhs := &Expression{
			Case:        4,
			Expression:  $1.(*Expression),
			Token:       $2,
			Term:        $3.(*Term),
		}
		$$ = lhs
		lx.emit(&Sub{Token: lhs.Token})
	}

Term:
	Factor
	{
		$$ = &Term{
			Factor:  $1.(*Factor),
		}
	}
|	Term '*' Factor
	{
		lx := yylex.(*lexer)
		lhs := &Term{
			Case:    1,
			Term:    $1.(*Term),
			Token:   $2,
			Factor:  $3.(*Factor),
		}
		$$ = lhs
		lx.emit(&Mul{Token: lhs.Token})
	}
|	Term '/' Factor
	{
		lx := yylex.(*lexer)
		lhs := &Term{
			Case:    2,
			Term:    $1.(*Term),
			Token:   $2,
			Factor:  $3.(*Factor),
		}
		$$ = lhs
		lx.emit(&Div{Token: lhs.Token})
	}

Factor:
	IDENT
	{
		lx := yylex.(*lexer)
		lhs := &Factor{
			Token:  $1,
		}
		$$ = lhs
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
		lx := yylex.(*lexer)
		lhs := &Factor{
			Case:    1,
			Number:  $1.(*Number),
		}
		$$ = lhs
		lx.emit(&PushConst{Value: lhs.Number.Value, Token: lhs.Number.Token})
	}
|	'(' Expression ')'
	{
		$$ = &Factor{
			Case:        2,
			Token:       $1,
			Expression:  $2.(*Expression),
			Token2:      $3,
		}
	}
