// Copyright 2015 The PL0 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pl0

import (
	"go/token"
	"io"

	"github.com/cznic/golex/lex"
	"github.com/cznic/xc"
)

type lexer struct {
	*lex.Lexer
	code        []Instruction
	enter       *Enter
	example     interface{}
	exampleRule int
	firstTok    xc.Token
	jmp         *Jmp
	lastPos     token.Pos
	nproc       int
	program     *Program
	report      *xc.Report
	scope       *Bindings
}

func newLexer(nm string, sz int, r io.RuneReader, report *xc.Report) (*lexer, error) {
	file := xc.FileSet.AddFile(nm, -1, sz+2)
	lx, err := lex.New(
		file,
		r,
		lex.ErrorFunc(func(pos token.Pos, msg string) {
			report.Err(pos, msg)
		}),
		lex.RuneClass(func(r rune) int { return int(r) }),
	)
	if err != nil {
		return nil, err
	}

	l := &lexer{
		Lexer:  lx,
		report: report,
	}
	l.emit(&Call{Target: 2})
	l.emit(&Halt{})
	l.enter = &Enter{NVars: -1}
	l.jmp = &Jmp{Target: -1}
	l.emit(l.enter)
	l.emit(l.jmp)
	return l, nil
}

func (l *lexer) newScope() *Bindings {
	l.scope = newBindings(l.scope)
	return l.scope
}

func (l *lexer) popScope(t xc.Token) *Bindings {
	if l.scope.Parent == nil {
		l.report.ErrTok(t, "cannot pop scope")
		return l.scope
	}

	l.scope = l.scope.Parent
	return l.scope
}

// Implements yyLexer.
func (l *lexer) Lex(lval *yySymType) int {
	r := rune(l.scan())
	if r == lex.RuneEOF {
		r = 0
	}

	pos := l.First.Pos()
	l.lastPos = pos
	c := lex.NewChar(pos, r)
	var val int
	switch r {
	case IDENT, NUMBER:
		val = xc.Dict.ID(l.TokenBytes(nil))
	}
	t := xc.Token{Char: c, Val: val}
	if !l.firstTok.Pos().IsValid() {
		l.firstTok = t
	}
	lval.Token = t
	return int(r)
}

// Implements yyLexer.
func (l *lexer) Error(e string) { l.report.Err(l.lastPos, e) }

// Reduced implements yyLexerEx
func (l *lexer) Reduced(rule, state int, lval *yySymType) (stop bool) {
	if n := l.exampleRule; n >= 0 && rule != n {
		return false
	}

	switch x := lval.node.(type) {
	case interface {
		fragment() interface{}
	}:
		l.example = x.fragment()
	default:
		l.example = x
	}
	return true
}

func (l *lexer) mustIdent(t xc.Token) (Node, int) {
	b, f := l.scope.lookup(t.Val)
	if b == nil {
		l.report.ErrTok(t, "undeclared identifier %s", t.S())
		return nil, -1
	}

	return b.Node, f
}

func (l *lexer) mustVarOrConst(t xc.Token) (Node, int) {
	n, f := l.mustIdent(t)
	if n == nil {
		return nil, -1
	}

	switch n.(type) {
	case *Variable, *ConstSpec:
		return n, f
	default:
		l.report.ErrTok(t, "%s is not a constant or a variable", t.S())
		return nil, -1
	}
}

func (l *lexer) mustVar(t xc.Token) (Node, int) {
	n, f := l.mustIdent(t)
	if n == nil {
		return nil, -1
	}

	if _, ok := n.(*Variable); !ok {
		l.report.ErrTok(t, "%s is not a variable", t.S())
		return nil, -1
	}

	return n, f
}

func (l *lexer) mustProc(t xc.Token) (Node, int) {
	n, f := l.mustIdent(t)
	if n == nil {
		return nil, -1
	}

	if _, ok := n.(*ProcSpec); !ok {
		l.report.ErrTok(t, "%s is not a procedure", t.S())
		return nil, -1
	}

	return n, f
}

func (l *lexer) bind(t xc.Token, n Node) {
	l.scope.Bind(t, n, l.report)
}

func (l *lexer) emit(i Instruction) int {
	addr := len(l.code)
	i.setAddr(addr)
	l.code = append(l.code, i)
	return addr
}
