package main

import (
	"fmt"

	"github.com/jakebailey/comb"
	"github.com/jakebailey/comb/combext"
)

var (
	integer = whitespaceAround(combext.Integer())
	addOp   = whitespaceAround(comb.Char('+', '-'))
	mulOp   = whitespaceAround(comb.Char('*', '/'))
	lParen  = whitespaceAround(comb.Char('('))
	rParen  = whitespaceAround(comb.Char(')'))
)

var (
	expr   comb.Parser
	term   comb.Parser
	factor comb.Parser
)

func init() {
	expr = comb.Sequence(
		func(results []comb.Result, start, end comb.Scanner) comb.Result {
			acc := results[0].Int64

			l := results[1].Interface.([]comb.Result)
			for i := 0; i < len(l); i += 2 {
				op := l[i].Runes[0]
				v := l[i+1].Int64

				switch op {
				case '+':
					acc += v
				case '-':
					acc -= v
				}
			}

			return comb.Result{
				Int64: acc,
			}
		},
		comb.Reference(&term),
		comb.Many(
			flatten,
			comb.Sequence(
				nil,
				addOp,
				comb.Reference(&term),
			),
		),
	)

	term = comb.Sequence(
		func(results []comb.Result, start, end comb.Scanner) comb.Result {
			acc := results[0].Int64

			l := results[1].Interface.([]comb.Result)
			for i := 0; i < len(l); i += 2 {
				op := l[i].Runes[0]
				v := l[i+1].Int64

				switch op {
				case '*':
					acc *= v
				case '/':
					acc /= v
				}
			}

			return comb.Result{
				Int64: acc,
			}
		},
		comb.Reference(&factor),
		comb.Many(
			flatten,
			comb.Sequence(
				nil,
				mulOp,
				comb.Reference(&term),
			),
		),
	)

	factor = comb.Or(
		integer,
		comb.Surround(
			lParen,
			comb.Reference(&expr),
			rParen,
		),
	)
}

func whitespaceAround(p comb.Parser) comb.Parser {
	return comb.Surround(
		combext.ManyWhitespace(),
		p,
		combext.ManyWhitespace(),
	)
}

func flatten(results []comb.Result, start, end comb.Scanner) comb.Result {
	out := make([]comb.Result, 0, len(results)*2)

	for _, r := range results {
		sub := r.Interface.([]comb.Result)
		out = append(out, sub...)
	}

	return comb.Result{
		Interface: out,
	}
}

func main() {
	test := "(1 + 2 * 3 + 9) * 2 + 1"

	s := comb.NewStringScanner(test)

	r, _ := expr.Parse(s)

	if r.Matched() {
		fmt.Printf("%v = %v\n", test, r.Int64)
	} else {
		fmt.Println(r.Err)
	}
}
