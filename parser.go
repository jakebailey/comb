package comb

import (
	"io"
)

// Parser describes comb parsers, which take a scanner,
// scan some amount of text, then return a result and
// the next scanner.
type Parser interface {
	Parse(s Scanner) (r Result, next Scanner)
}

type parserFunc struct {
	fn func(Scanner) (Result, Scanner)
}

func (p parserFunc) Parse(s Scanner) (r Result, next Scanner) {
	return p.fn(s)
}

// ParserFunc turns a parser function into a Parser.
func ParserFunc(fn func(Scanner) (Result, Scanner)) Parser {
	return parserFunc{fn: fn}
}

// Reference takes a pointer to a Parser, and only dereferences it
// when Parse is called.
func Reference(p *Parser) Parser {
	return ParserFunc(func(s Scanner) (Result, Scanner) {
		return (*p).Parse(s)
	})
}

// Tag sets the tag of a parser's result.
func Tag(tag string, parser Parser) Parser {
	return ParserFunc(func(s Scanner) (Result, Scanner) {
		r, next := parser.Parse(s)
		r.Tag = tag
		return r, next
	})
}

// Ignore sets the result of a Parser to be Ignored.
func Ignore(parser Parser) Parser {
	return ParserFunc(func(s Scanner) (Result, Scanner) {
		r, next := parser.Parse(s)
		r.Ignore = true
		return r, next
	})
}

// EOF matches only at EOF.
func EOF() Parser {
	return ParserFunc(func(s Scanner) (Result, Scanner) {
		r, next, err := s.Next()
		if err != io.EOF {
			return Failedf("expected EOF, got '%c'", r), next
		}

		return Result{}, next
	})
}

// Maybe tries a parser and returns its result if it matches,
// otherwise, it returns an empty result and the original scanner.
func Maybe(parser Parser) Parser {
	return ParserFunc(func(s Scanner) (Result, Scanner) {
		r, next := parser.Parse(s)
		if r.Matched() {
			return r, next
		}
		return Result{}, s
	})
}
