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

// Named sets the ParserName of a parser's result.
func Named(name string, p Parser) Parser {
	return ParserFunc(func(s Scanner) (Result, Scanner) {
		r, next := p.Parse(s)
		r.ParserName = name
		return r, next
	})
}

// Ignore sets the result of a Parser to be Ignored
func Ignore(p Parser) Parser {
	return ParserFunc(func(s Scanner) (Result, Scanner) {
		r, next := p.Parse(s)
		r.Ignore = true
		return r, next
	})
}

// AnyChar accepts any single character.
func AnyChar() Parser {
	return ParserFunc(func(s Scanner) (Result, Scanner) {
		_, next, err := s.Next()

		if err != nil {
			return Failed(err), next
		}

		return Result{
			Runes: s.Between(next),
		}, next
	})
}

// Char accepts a single given character.
func Char(chars ...rune) Parser {
	m := make(map[rune]struct{}, len(chars))
	for _, r := range chars {
		m[r] = struct{}{}
	}

	return ParserFunc(func(s Scanner) (Result, Scanner) {
		r, next, err := s.Next()
		if err != nil {
			return Failed(err), next
		}

		if _, ok := m[r]; ok {
			return Result{
				Runes: s.Between(next),
			}, next
		}

		return Failedf("unexpected character '%c'", r), s
	})
}

// Take accepts n characters and returns the runes captured.
func Take(n int) Parser {
	return ParserFunc(func(s Scanner) (Result, Scanner) {
		next := s
		var err error

		for i := 0; i < n; i++ {
			_, next, err = next.Next()

			if err != nil {
				return Failed(err), next
			}
		}

		return Result{
			Runes: s.Between(next),
		}, next
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
