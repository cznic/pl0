// Copyright 2015 The PL0 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pl0

import (
	"flag"
	"fmt"
	"go/scanner"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"
)

func caller(s string, va ...interface{}) {
	if s == "" {
		s = strings.Repeat("%v ", len(va))
	}
	_, fn, fl, _ := runtime.Caller(2)
	fmt.Fprintf(os.Stderr, "caller: %s:%d: ", path.Base(fn), fl)
	fmt.Fprintf(os.Stderr, s, va...)
	fmt.Fprintln(os.Stderr)
	_, fn, fl, _ = runtime.Caller(1)
	fmt.Fprintf(os.Stderr, "\tcallee: %s:%d: ", path.Base(fn), fl)
	fmt.Fprintln(os.Stderr)
	os.Stderr.Sync()
}

func dbg(s string, va ...interface{}) {
	if s == "" {
		s = strings.Repeat("%v ", len(va))
	}
	_, fn, fl, _ := runtime.Caller(1)
	fmt.Fprintf(os.Stderr, "dbg %s:%d: ", path.Base(fn), fl)
	fmt.Fprintf(os.Stderr, s, va...)
	fmt.Fprintln(os.Stderr)
	os.Stderr.Sync()
}

func TODO(...interface{}) string {
	_, fn, fl, _ := runtime.Caller(1)
	return fmt.Sprintf("TODO: %s:%d:\n", path.Base(fn), fl)
}

func use(...interface{}) {}

// ============================================================================

func init() {
	flag.IntVar(&yyDebug, "yydebug", 0, "")
}

func Test(t *testing.T) {
	_, err := Parse("empty.pl0")
	if err == nil {
		t.Error("expected error")
	}
	t.Log(err)
}

func Example_redeclaration() {
	_, err := ParseString("",
		`
	CONST a = 42, b = 24, a = 314;
	
	VAR b, c, d;
	
	PROCEDURE c;
	BEGIN
	END;
	
	PROCEDURE d;
	BEGIN
	END;
	
	PROCEDURE e;
	BEGIN
	END;
	
	PROCEDURE e;
	BEGIN
	END;
	
	BEGIN
	END.
`)

	scanner.PrintError(os.Stdout, err)
	// Output:
	// 2:24: redeclaration of a at 2:8
	// 4:6: redeclaration of b at 2:16
	// 6:12: redeclaration of c at 4:9
	// 10:12: redeclaration of d at 4:12
	// 18:12: redeclaration of e at 14:12
}

func Example_undeclared() {
	_, err := ParseString("",
		`
	CONST a = 42;
	
	VAR b;
	
	PROCEDURE c;
	BEGIN
		CALL aa;
		CALL c;
	END;
	
	BEGIN
		b := a;
		b := d;
		e := 10;
		CALL c;
		CALL f;
		? b;
		? g;
	END.
`)

	scanner.PrintError(os.Stdout, err)
	// Output:
	// 8:8: undeclared identifier aa
	// 14:8: undeclared identifier d
	// 15:3: undeclared identifier e
	// 17:8: undeclared identifier f
	// 19:5: undeclared identifier g
}

func Example_illegalAssignment() {
	_, err := ParseString("",
		`
	CONST a = 42;
	
	VAR b;
	
	PROCEDURE c;
	BEGIN
	END;
	
	BEGIN
		a := 1;
		b := 1;
		c := 1;
	END.
`)

	scanner.PrintError(os.Stdout, err)
	// Output:
	// 11:3: a is not a variable
	// 13:3: c is not a variable
}

func Example_illegalCall() {
	_, err := ParseString("",
		`
	CONST a = 42;
	
	VAR b;
	
	PROCEDURE c;
	BEGIN
	END;
	
	BEGIN
		CALL a;
		CALL b;
		CALL c;
	END.
`)

	scanner.PrintError(os.Stdout, err)
	// Output:
	// 11:8: a is not a procedure
	// 12:8: b is not a procedure
}

func Example_illegalRead() {
	_, err := ParseString("",
		`
	CONST a = 42;
	
	VAR b;
	
	PROCEDURE c;
	BEGIN
	END;
	
	BEGIN
		? a;
		? b;
		? c;
	END.
`)

	scanner.PrintError(os.Stdout, err)
	// Output:
	// 11:5: a is not a variable
	// 13:5: c is not a variable
}

func Example_illegalFactor() {
	_, err := ParseString("",
		`
	CONST a = 42;
	
	VAR b;
	
	PROCEDURE c;
	BEGIN
	END;
	
	BEGIN
		b := a;
		b := b;
		b := c;
	END.
`)

	scanner.PrintError(os.Stdout, err)
	// Output:
	// 13:8: c is not a constant or a variable
}

func Example_program1() {
	prog, err := ParseString("",
		`
	// Original source: https://en.wikipedia.org/wiki/PL/0#Examples
	
	VAR x, squ;
	
	PROCEDURE square;
	BEGIN
		squ := x * x
	END;
	
	BEGIN
		x := 1;
		WHILE x <= 10 DO
		BEGIN
			CALL square;
			! squ;
			x := x + 1
		END
	END.
`)

	if err != nil {
		panic(err)
	}

	if err := prog.Run(); err != nil {
		panic(err)
	}
	// Output:
	// 1
	// 4
	// 9
	// 16
	// 25
	// 36
	// 49
	// 64
	// 81
	// 100
}

func Example_program2() {
	prog, err := ParseString("",
		`
	// Original source: https://en.wikipedia.org/wiki/PL/0#Examples
	
	CONST
		m =  7,
		n = 85;
	
	VAR
	 	x, y, z, q, r;
	
	PROCEDURE multiply;
	VAR a, b;
	
	BEGIN
		a := x;
		b := y;
		z := 0;
		WHILE b > 0 DO BEGIN
			IF ODD b THEN z := z + a;
			a := 2 * a;
			b := b / 2
		END;
		! z
	END;
	
	PROCEDURE divide;
	VAR w;
	BEGIN
		r := x;
		q := 0;
		w := y;
		WHILE w <= r DO w := 2 * w;
		WHILE w > y DO BEGIN
			q := 2 * q;
			w := w / 2;
			IF w <= r THEN BEGIN
				r := r - w;
				q := q + 1
			END
		END;
		! q;
		! r
	END;
	
	PROCEDURE gcd;
	VAR f, g;
	BEGIN
		f := x;
		g := y;
		WHILE f # g DO BEGIN
			IF f < g THEN g := g - f;
			IF g < f THEN f := f - g
		END;
		z := f;
		! z;
	END;
	
	BEGIN
		x := m;
		y := n;
		CALL multiply;
		x := 25;
		y :=  3;
		CALL divide;
		x := 84;
		y := 36;
		CALL gcd
	END.
`)

	if err != nil {
		panic(err)
	}

	if err := prog.Run(); err != nil {
		panic(err)
	}
	// Output:
	// 595
	// 8
	// 1
	// 12
}

func Example_stackFrames() {
	prog, err := ParseString("",
		`
	VAR a1, recur;
	
	PROCEDURE p2;
	VAR a2;

		PROCEDURE p3;
		VAR a3;
		BEGIN // p3
			a3 := 3;
			! 333333;
			! a1;
			! a2;
			! a3;
		END;

		PROCEDURE p4;
		VAR a3;
		BEGIN // p4
			a3 := 4;
			CALL p3;
			! 444444;
			! a1;
			! a2;
			! a3;
			WHILE recur > 0 DO BEGIN
				recur := recur - 1;
				CALL p2;
			END;
		END;

	BEGIN // p2
		a2 := 2;
		! 222222;
		! a1;
		! a2;
		CALL p4;
	END;
	
	BEGIN // main
		a1 := 1;
		recur := 2;
		! 111111;
		! a1;
		CALL p2;
		! 999999;
	END.
`)

	if err != nil {
		panic(err)
	}

	if err := prog.Run(); err != nil {
		panic(err)
	}
	// Output:
	// 111111
	// 1
	// 222222
	// 1
	// 2
	// 333333
	// 1
	// 2
	// 3
	// 444444
	// 1
	// 2
	// 4
	// 222222
	// 1
	// 2
	// 333333
	// 1
	// 2
	// 3
	// 444444
	// 1
	// 2
	// 4
	// 222222
	// 1
	// 2
	// 333333
	// 1
	// 2
	// 3
	// 444444
	// 1
	// 2
	// 4
	// 999999
}

func Example_trace() {
	prog, err := ParseString("",
		`
	VAR v;
	
	PROCEDURE addOne;
	BEGIN
		v := v + 1;
	END;

	BEGIN // main
		v := 42;
		CALL addOne;
		! v;
	END.
`)

	if err != nil {
		panic(err)
	}

	if err := prog.Run(Trace()); err != nil {
		panic(err)
	}
	// Output:
	// T 0001,  FP 0000, SP 0000:	00000000:	call		00000002	// -
	// T 0002,  FP 0000, SP 0001:	00000002:	enter		1, 0		// 2:2:
	// T 0003,  FP 0003, SP 0004:	00000003:	jmp		00000011
	// T 0004,  FP 0003, SP 0004:	00000011:	pushConst	42		// 10:8
	// T 0005,  FP 0003, SP 0005:	00000012:	popVar		0		// 2:6: v
	// T 0006,  FP 0003, SP 0004:	00000013:	call		00000004	// 11:3
	// T 0007,  FP 0003, SP 0005:	00000004:	enter		0, 1		// 4:12: addOne
	// T 0008,  FP 0008, SP 0008:	00000005:	jmp		00000006
	// T 0009,  FP 0008, SP 0008:	00000006:	pushVar		1, 0		// 2:6: v
	// T 0010,  FP 0008, SP 0009:	00000007:	pushConst	1		// 6:12
	// T 0011,  FP 0008, SP 0010:	00000008:	add				// 6:10
	// T 0012,  FP 0008, SP 0009:	00000009:	popVar		1, 0		// 2:6: v
	// T 0013,  FP 0008, SP 0008:	00000010:	leave
	// T 0014,  FP 0003, SP 0004:	00000014:	pushVar		0		// 2:6: v
	// T 0015,  FP 0003, SP 0005:	00000015:	write				// 12:3
	// 43
	// T 0016,  FP 0003, SP 0004:	00000016:	leave
	// T 0017,  FP 0000, SP 0000:	00000001:	halt
}

func Example_stackTrace() {
	prog, err := ParseString("",
		`
	VAR v;
	
	PROCEDURE addOne;
	BEGIN
		v := v + 1;
	END;

	BEGIN // main
		v := 42;
		CALL addOne;
		! v;
	END.
`)

	if err != nil {
		panic(err)
	}

	if err := prog.Run(TraceStack()); err != nil {
		panic(err)
	}
	// Output:
	// T 0001,  FP 0000, SP 0000, []:	00000000:	call		00000002	// -
	// T 0002,  FP 0000, SP 0001, [1]:	00000002:	enter		1, 0		// 2:2:
	// T 0003,  FP 0003, SP 0004, [1 0 0 0]:	00000003:	jmp		00000011
	// T 0004,  FP 0003, SP 0004, [1 0 0 0]:	00000011:	pushConst	42		// 10:8
	// T 0005,  FP 0003, SP 0005, [1 0 0 0 42]:	00000012:	popVar		0		// 2:6: v
	// T 0006,  FP 0003, SP 0004, [1 0 0 42]:	00000013:	call		00000004	// 11:3
	// T 0007,  FP 0003, SP 0005, [1 0 0 42 14]:	00000004:	enter		0, 1		// 4:12: addOne
	// T 0008,  FP 0008, SP 0008, [1 0 0 42 14 3 3 1]:	00000005:	jmp		00000006
	// T 0009,  FP 0008, SP 0008, [1 0 0 42 14 3 3 1]:	00000006:	pushVar		1, 0		// 2:6: v
	// T 0010,  FP 0008, SP 0009, [1 0 0 42 14 3 3 1 42]:	00000007:	pushConst	1		// 6:12
	// T 0011,  FP 0008, SP 0010, [1 0 0 42 14 3 3 1 42 1]:	00000008:	add				// 6:10
	// T 0012,  FP 0008, SP 0009, [1 0 0 42 14 3 3 1 43]:	00000009:	popVar		1, 0		// 2:6: v
	// T 0013,  FP 0008, SP 0008, [1 0 0 43 14 3 3 1]:	00000010:	leave
	// T 0014,  FP 0003, SP 0004, [1 0 0 43]:	00000014:	pushVar		0		// 2:6: v
	// T 0015,  FP 0003, SP 0005, [1 0 0 43 43]:	00000015:	write				// 12:3
	// 43
	// T 0016,  FP 0003, SP 0004, [1 0 0 43]:	00000016:	leave
	// T 0017,  FP 0000, SP 0000, []:	00000001:	halt
}

func Example_list() {
	prog, err := ParseString("",
		`
	VAR v;
	
	PROCEDURE addOne;
	BEGIN
		v := v + 1;
	END;

	BEGIN // main
		v := 42;
		CALL addOne;
		! v;
	END.
`)

	if err != nil {
		panic(err)
	}

	prog.List(os.Stdout)
	// Output:
	// 00000000:	call		00000002	// -
	// 00000001:	halt
	//
	// 00000002:	enter		1, 0		// 2:2:
	// 00000003:	jmp		00000011
	//
	// 00000004:	enter		0, 1		// 4:12: addOne
	// 00000005:	jmp		00000006
	//
	// 00000006:	pushVar		1, 0		// 2:6: v
	// 00000007:	pushConst	1		// 6:12
	// 00000008:	add				// 6:10
	// 00000009:	popVar		1, 0		// 2:6: v
	// 00000010:	leave
	//
	// 00000011:	pushConst	42		// 10:8
	// 00000012:	popVar		0		// 2:6: v
	// 00000013:	call		00000004	// 11:3
	// 00000014:	pushVar		0		// 2:6: v
	// 00000015:	write				// 12:3
	// 00000016:	leave
}
