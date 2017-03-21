package helpers

import (
	"strconv"

	"github.com/jakebailey/comb"
)

// var integerParser = comb.SequenceRunes(
// 	comb.Maybe(
// 		comb.Char('-'),
// 	),
// 	comb.Or(
// 		comb.SequenceRunes(
// 			comb.Token("0x", "0X"),
// 			comb.OnePlusRunes(
// 				comb.Or(
// 					comb.CharRange('a', 'z'),
// 					comb.CharRange('a', 'z'),
// 					comb.Digit(),
// 				),
// 			),
// 		),
// 		comb.Digits(),
// 	),
// )

var integerParser = comb.Regexp(`-?0[xX][\da-fA-F]+|\d+`)

// IntegerParser parses an integer in base 8, 10, or 16 using strconv.
// It first applies ParseUint, then ParseInt, taking the first non-failing
// parsed int64. As an optimization, the string "0" will be immediately
// converted without any strconv calls.
func IntegerParser() comb.Parser {
	return comb.ParserFunc(func(s comb.Scanner) (comb.Result, comb.Scanner) {
		r, next := integerParser.Parse(s)
		if !r.Matched() {
			return r, next
		}

		var i int64

		rs := string(r.Runes)

		if rs == "0" {
			return comb.Result{}, next
		}

		ui, err := strconv.ParseUint(rs, 0, 64)
		if err == nil {
			i = int64(ui)
		} else {
			i, err = strconv.ParseInt(rs, 0, 64)
		}

		if err != nil {
			return comb.Failed(err), next
		}

		return comb.Result{
			Int64: i,
		}, next
	})
}
