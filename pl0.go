// Copyright 2015 The PL0 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run generate.go
//go:generate golex -o scanner.go scanner.l
//go:generate go run generate.go -2

// Package pl0 is a PL/0[0] compiler front end.
//
// For an example using the front end see the pl0 command in the pl0
// subdirectory.
//
// Purpose
//
// This is a showcase demonstrating use of:
//
//	fe	http://godoc.org/github.com/cznic/fe
//	golex	http://godoc.org/github.com/cznic/golex
//	goyacc	http://godoc.org/github.com/cznic/goyacc
//	lex	http://godoc.org/github.com/cznic/golex/lex
//	xc	http://godoc.org/github.com/cznic/xc
//	yy	http://godoc.org/github.com/cznic/yy
//
// Links
//
// Referenced from elsewhere
//
//	[0]: https://en.wikipedia.org/wiki/PL/0
package pl0

import (
	"bufio"
	"bytes"
	"fmt"
	"go/token"
	"io"
	"os"
	"reflect"
	"strconv"

	"github.com/cznic/golex/lex"
	"github.com/cznic/strutil"
	"github.com/cznic/xc"
)

// Node must be implemented by all AST nodes.
type Node interface {
	Pos() token.Pos
}

var (
	printHooks = strutil.PrettyPrintHooks{}
)

func init() {
	for k, v := range xc.PrintHooks {
		printHooks[k] = v
	}
	lcRT := reflect.TypeOf(lex.Char{})
	lcH := func(f strutil.Formatter, v interface{}, prefix, suffix string) {
		c := v.(lex.Char)
		r := c.Rune
		s := yySymName(int(r))
		if x := s[0]; x >= '0' && x <= '9' {
			s = strconv.QuoteRune(r)
		}
		f.Format("%s%v: %s"+suffix, prefix, xc.FileSet.Position(c.Pos()), s)
	}
	printHooks[lcRT] = lcH
	printHooks[reflect.TypeOf(xc.Token{})] = func(f strutil.Formatter, v interface{}, prefix, suffix string) {
		t := v.(xc.Token)
		if !t.Pos().IsValid() {
			return
		}

		lcH(f, t.Char, prefix, "")
		if s := xc.Dict.S(t.Val); len(s) != 0 {
			f.Format(" %q", s)
		}
		f.Format(suffix)
	}
}

// PrettyString returns a pretty formatted representation of AST nodes.
func PrettyString(v interface{}) string { return strutil.PrettyString(v, "", "", printHooks) }

func exampleAST(rule int, src string) interface{} {
	lx, _ := parse(fmt.Sprintf("example%v.pl0", rule), len(src), bytes.NewBufferString(src), rule)
	return lx.example
}

func parse(nm string, sz int, r io.Reader, exampleRule int) (*lexer, error) {
	defer func() {
		if c, ok := r.(io.ReadCloser); ok {
			c.Close()
		}
	}()

	rr, ok := r.(io.RuneReader)
	if !ok {
		rr = bufio.NewReader(r)
	}

	report := xc.NewReport()
	lx, err := newLexer(nm, sz, rr, report)
	if err != nil {
		report.Err(0, "%v", err)
		return lx, report.Errors(true)
	}

	lx.exampleRule = exampleRule
	lx.newScope()
	yyParse(lx)
	return lx, report.Errors(true)
}

// Parse parses PL/0 source in nm and returns a *Program, or an error, if any.
// The returned error is either nil or it is guaranteed to be of type
// go/scanner.ErrorList.
func Parse(nm string) (*Program, error) {
	f, err := os.Open(nm)
	if err != nil {
		return nil, err
	}

	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}

	lx, err := parse(nm, int(fi.Size()), f, 0)
	if err != nil {
		return nil, err
	}

	return lx.program, nil
}

// ParseString is like Parse except the source code comes from a string.
func ParseString(nm, src string) (*Program, error) {
	lx, err := parse(nm, len(src), bytes.NewBufferString(src), 0)
	if err != nil {
		return nil, err
	}

	return lx.program, nil
}
