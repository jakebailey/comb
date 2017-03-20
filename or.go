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
