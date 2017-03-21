package comb

// Sequence runs multiple parsers in a sequence, combining results
// with a combiner function. If combiner is nil, then SliceCombiner
// is used. Sequence must allocate a slice of results the same length
// as the number of parsers required.
//
// If you only need the runes captured by Sequence, use SequenceRunes instead.
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
// slice in Interface. If a result is set to be ignored, the result
// will not be in the new result slice.
func SliceCombiner(results []Result, begin, end Scanner) Result {
	ignored := 0

	for i, r := range results {
		if r.Ignore {
			ignored++
		} else {
			results[i-ignored] = results[i]
		}
	}

	results = results[:len(results)-ignored]

	return Result{
		Interface: results,
	}
}

// SequenceRunes is like Sequence, but does not capture all results,
// instead returning the runes between the start and end of the matching
// region. Unlike Sequence, this does not allocate anything.
// SequenceRunes does not read any results, just the returned scanners,
// so cannot respect the Ignored option.
func SequenceRunes(parsers ...Parser) Parser {
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

// Surround surrounds a parser with two parsers, and returns
// the surrounded value. This is equivalent to Sequence with a combiner
// which returns the middle result.
func Surround(left, parser, right Parser) Parser {
	return Sequence(
		func(results []Result, begin, end Scanner) Result {
			return results[1]
		},
		left,
		parser,
		right,
	)
}
