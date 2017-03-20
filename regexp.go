package comb

import (
	"regexp"
	"unicode/utf8"
)

// Regexp compiles a Go regexp into a parser. If the pattern does
// not begin with ^, one will be added, as the parser must begin
// with the next rune.
func Regexp(pattern string) Parser {
	realPattern := pattern
	if realPattern == "" || realPattern[0] != '^' {
		realPattern = "^" + realPattern
	}

	re := regexp.MustCompile(realPattern)

	return ParserFunc(func(s Scanner) (Result, Scanner) {
		sr := &scannerReader{s}

		match := re.FindReaderIndex(sr)
		if match == nil {
			return Failedf("regexp %q did not match", pattern), s
		}

		var r rune
		next := s
		var err error

		count := match[1]

		for count > 0 {
			r, next, err = next.Next()
			if err != nil {
				return Failed(err), next
			}

			count -= utf8.RuneLen(r)
		}

		if count < 0 {
			panic("bug: got more bytes than regexp match specified")
		}

		return Result{
			Runes: s.Between(next),
		}, next
	})
}

type scannerReader struct {
	next Scanner
}

func (s *scannerReader) ReadRune() (rune, int, error) {
	r, next, err := s.next.Next()
	if err != nil {
		return 0, 0, err
	}
	s.next = next
	return r, utf8.RuneLen(r), nil
}
