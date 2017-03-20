package comb

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

		if _, ok := m[r]; !ok {
			return Failedf("unexpected character '%c'", r), s
		}

		return Result{
			Runes: s.Between(next),
		}, next
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

// NotChar only accepts a char not given.
func NotChar(chars ...rune) Parser {
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
			return Failedf("unexpected character '%c'", r), s
		}

		return Result{
			Runes: s.Between(next),
		}, next
	})
}

// CharRange accepts chars in an inclusive range.
func CharRange(from, to rune) Parser {
	return ParserFunc(func(s Scanner) (Result, Scanner) {
		r, next, err := s.Next()
		if err != nil {
			return Failed(err), next
		}

		if r < from || r > to {
			return Failedf("unexpected character '%c'", r), s
		}

		return Result{
			Runes: s.Between(next),
		}, next
	})
}

// CharsIn accepts any of the chars in a given string.
func CharsIn(s string) Parser {
	return Char([]rune(s)...)
}

// Digit accepts a single digit from 0 to 9.
func Digit() Parser {
	return CharRange('0', '9')
}

// Digits accepts 1+ digits.
func Digits() Parser {
	return TextOnePlus(Digit())
}

// AlphaLower accepts characters from a to z.
func AlphaLower() Parser {
	return CharRange('a', 'z')
}

// AlphaUpper accepts characters from A to Z.
func AlphaUpper() Parser {
	return CharRange('A', 'Z')
}

// Alpha accepts characters from either a to z or A to Z.
func Alpha() Parser {
	return Or(
		AlphaLower(),
		AlphaUpper(),
	)
}

// AlphaDigit accepts characters that match Alpha or Digit.
func AlphaDigit() Parser {
	return Or(
		Alpha(),
		Digit(),
	)
}

// Whitespace accepts any of ' ', '\t', '\n', or '\r'.
func Whitespace() Parser {
	return Char(' ', '\t', '\n', '\r')
}

// ManyWhitespace accepts 0+ instances of whitespace.
func ManyWhitespace() Parser {
	return TextMany(Whitespace())
}

// OnePlusWhitespace accepts 1+ instances of whitespace.
func OnePlusWhitespace() Parser {
	return TextOnePlus(Whitespace())
}
