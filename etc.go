// Copyright 2015 The PL0 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pl0

import (
	"fmt"
	"go/token"

	"github.com/cznic/xc"
)

var (
	b2i = map[bool]int{true: 1}
)

func position(pos token.Pos) token.Position { return xc.FileSet.Position(pos) }

// Binding is the value bound to a name at Pos.
type Binding struct {
	Node Node
	Pos  token.Pos
}

// Bindings capture name -> Binding projections in a scope.
type Bindings struct {
	Map      map[int]*Binding
	Parent   *Bindings
	addr     int
	ifEnd    *JmpZero
	lnl      int // Lexical nesting level
	nvar     int
	while0   int
	whileEnd *JmpZero
}

func newBindings(parent *Bindings) *Bindings {
	lnl := 0
	if parent != nil {
		lnl = parent.lnl + 1
	}
	return &Bindings{
		Map:    map[int]*Binding{},
		Parent: parent,
		lnl:    lnl,
	}
}

func (b *Bindings) lookup(nm int) (*Binding, int) {
	f := 0
	for b != nil {
		if x := b.Map[nm]; x != nil {
			return x, f
		}

		b = b.Parent
		f++
	}
	return nil, -1
}

// Bind attempts to bind the name in t to node. Errors are reported to report.
func (b *Bindings) Bind(t xc.Token, node Node, report *xc.Report) {
	nm := t.Val
	if x := b.Map[nm]; x != nil {
		ok := false
		switch e := x.Node.(type) {
		case *ConstSpec, *Variable:
		case *ProcSpec:
			ok = e == nil
		default:
			panic("internal error")
		}

		if !ok {
			report.ErrTok(t, "redeclaration of %s at %s", xc.Dict.S(nm), position(x.Pos))
			return
		}
	}

	b.Map[nm] = &Binding{node, t.Pos()}
}

func (n *VarsOpt) nvars() int {
	if n == nil {
		return 0
	}

	return n.Vars.nvars
}

func phelp(t xc.Token) string {
	if len(t.S()) != 0 {
		return fmt.Sprintf(" %s", t.S())
	}

	return ""
}
