package comb

// Many looks for a series of 0+ matches of a parser,
// then combines the results with a combiner. If combiner is nil,
// SliceCombiner is used.
//
// If you only need the runes captured by Many, use TextMany instead.
func Many(combiner ResultCombiner, parser Parser) Parser {
	if combiner == nil {
		combiner = SliceCombiner
	}

	return ParserFunc(func(s Scanner) (Result, Scanner) {
		var results []Result
		next := s

		for {
			r, maybeNext := parser.Parse(next)
			if !r.Matched() {
				break
			}

			next = maybeNext
			results = append(results, r)
		}

		return combiner(results, s, next), next
	})
}

// ManyRunes looks for a series of 0+ matches of a parser,
// then returns the runes captured.
func ManyRunes(parser Parser) Parser {
	return ParserFunc(func(s Scanner) (Result, Scanner) {
		next := s

		for {
			r, maybeNext := parser.Parse(next)
			if !r.Matched() {
				break
			}
			next = maybeNext
		}

		return Result{
			Runes: s.Between(next),
		}, next
	})
}

// OnePlus looks for a series of 1+ matches of a parser.
//
// If you only need the runes captured by OnePlus, use TextOnePlus instead.
func OnePlus(combiner ResultCombiner, parser Parser) Parser {
	if combiner == nil {
		combiner = SliceCombiner
	}

	return ParserFunc(func(s Scanner) (Result, Scanner) {
		r, next := parser.Parse(s)
		if !r.Matched() {
			return r, next
		}

		results := []Result{r}

		for {
			r, maybeNext := parser.Parse(next)
			if !r.Matched() {
				break
			}

			next = maybeNext
			results = append(results, r)
		}

		return combiner(results, s, next), next
	})
}

// OnePlusRunes looks for a series of 1+ matches of a parser,
// then returns the runes captured.
func OnePlusRunes(parser Parser) Parser {
	return ParserFunc(func(s Scanner) (Result, Scanner) {
		r, next := parser.Parse(s)
		if !r.Matched() {
			return r, next
		}

		for {
			r, maybeNext := parser.Parse(next)
			if !r.Matched() {
				break
			}
			next = maybeNext
		}

		return Result{
			Runes: s.Between(next),
		}, next
	})
}
