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
