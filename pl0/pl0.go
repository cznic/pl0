// Copyright 2015 The PL0 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command pl0 compiles and executes pl0[0] programs.
//
// Usage
//
//	$ pl0 [-l] [-t] program.pl0
//
//	-l	List assembly code and exit.
//	-s	Trace execution and stack.
//	-t	Trace execution.
//
// Links
//
// Referenced from elsewhere
//
//	[0]: https://en.wikipedia.org/wiki/PL/0
package main

import (
	"flag"
	"go/scanner"
	"log"
	"os"

	"github.com/cznic/tmp/pl0" //TODO -tmp
)

func main() {
	log.SetFlags(0)
	list := flag.Bool("l", false, "List assembly code and exit.")
	trace := flag.Bool("t", false, "Trace execution.")
	stack := flag.Bool("s", false, "Trace execution and stack.")
	flag.Parse()
	if n := flag.NArg(); n != 1 {
		log.Fatalf("expected 1 argument, got %v", n)
	}

	prog, err := pl0.Parse(flag.Arg(0))
	if err != nil {
		scanner.PrintError(os.Stderr, err)
		os.Stderr.Sync()
		os.Exit(1)
	}

	if *list {
		prog.List(os.Stdout)
		return
	}

	switch {
	case *stack:
		err = prog.Run(pl0.TraceStack())
	case *trace:
		err = prog.Run(pl0.Trace())
	default:
		err = prog.Run()
	}
	if err != nil {
		log.Fatal(err)
	}
}
