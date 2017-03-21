package comb

// Or checks parsers in order, returning the first match.
func Or(parsers ...Parser) Parser {
	return ParserFunc(func(s Scanner) (Result, Scanner) {
		for _, p := range parsers {
			r, next := p.Parse(s)

			if r.Matched() {
				return r, next
			}
		}

		return Failedf("no parser matched"), s
	})
}

// OrLongest is like Or, but returns the result of the parser
// that captures the most text. Ties are broken by taking the
// first result. In order to do this, *every* parser will be run,
// so keep that in mind.
func OrLongest(parsers ...Parser) Parser {
	return ParserFunc(func(s Scanner) (Result, Scanner) {
		matched := false
		var maxResult Result
		var maxNext Scanner
		first := true

		for _, p := range parsers {
			r, next := p.Parse(s)

			if !r.Matched() {
				continue
			}

			matched = true
			if first || len(s.Between(next)) > len(s.Between(maxNext)) {
				first = false
				maxResult = r
				maxNext = next
			}
		}

		if !matched {
			return Failedf("no parser matched"), s
		}

		return maxResult, maxNext
	})
}
