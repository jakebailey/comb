package comb

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

// Char accepts a single given character.
func Char(chars ...rune) Parser {
	m := make(map[rune]struct{}, len(chars))
	for _, r := range chars {
		m[r] = struct{}{}
	}

	return ParserFunc(func(s Scanner) (Result, Scanner) {
		r, next, err := s.Next()
		if err != nil {
			return Failed(err), s
		}

		if _, ok := m[r]; ok {
			return Result{
				Runes: s.Between(next),
			}, next
		}

		return Failedf("unexpected character '%c'", r), s
	})
}
