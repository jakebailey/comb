package comb

// Sequence runs multiple parsers in a sequence, combining results
// with a combiner function. If combiner is nil, then SliceCombiner is used.
// Sequence must allocate a slice of results the same length as the number of parsers required.
//
// If you only need the runes captured by Sequence, use TextSequence instead.
func Sequence(combiner ResultCombiner, parsers ...Parser) Parser {
	if combiner == nil {
		combiner = SliceCombiner
	}

	return ParserFunc(func(s Scanner) (Result, Scanner) {
		results := make([]Result, len(parsers))

		var r Result
		next := s

		for i, p := range parsers {
			r, next = p.Parse(next)

			if !r.Matched() {
				return r, next
			}

			results[i] = r
		}

		return combiner(results, s, next), next
	})
}

// ResultCombiner is a function that takes a slice of results
// and surrounding scanners and combines them into a single result.
type ResultCombiner func(results []Result, begin, end Scanner) Result

// SliceCombiner combines results by returning a Result with the
// slice in Interface.
func SliceCombiner(results []Result, begin, end Scanner) Result {
	return Result{
		Interface: results,
	}
}

// TextCombiner combines results by returning a Result with the runes
// between begin and end.
func TextCombiner(results []Result, begin, end Scanner) Result {
	return Result{
		Runes: begin.Between(end),
	}
}

// TextSequence is like Sequence, but does not capture all results,
// instead returning the runes between the start and end of the matching
// region. Unlike Sequence, this does not allocate anything.
func TextSequence(parsers ...Parser) Parser {
	return ParserFunc(func(s Scanner) (Result, Scanner) {
		var r Result
		next := s

		for _, p := range parsers {
			r, next = p.Parse(next)

			if !r.Matched() {
				return r, next
			}
		}

		return Result{
			Runes: s.Between(next),
		}, next
	})
}
