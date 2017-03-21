/*
Package comb implements an efficient parser combinator framework.

comb offers various useful parsers, including nonterminals such as characters,
character ranges, tokens, and regular expressions. Also included are sequences,
repetitions, and optional parsers.

For example, here is a parser that recognizes the decimal, octal,
and hex integers accepted in Go.


    var integerParser = comb.SequenceRunes(
        comb.Maybe(
            comb.Char('-'),
        ),
        comb.Or(
            comb.SequenceRunes(
                comb.Token("0x", "0X"),
                comb.OnePlusRunes(
                    comb.Or(
                        comb.CharRange('a', 'z'),
                        comb.CharRange('a', 'z'),
                        Digit(),
                    ),
                ),
            ),
            Digits(),
        ),
    )

(Though, this is more succinctly expressed with a regular expression, giving a
moderate performance gain.)

comb uses a scanner which traverses a rune slice. All builtin parsers
return results that are simply slices of the original data, keeping copying
to a minimum.

In the _examples directory, you can find examples of comb in use, including
a recursive expression calculator.

The combext package offers other general-use parsers (such as alpha-numeric
characters, whitespace, etc) that may be frequently needed, though not always
used.
*/
package comb

//go:generate sh -c "godoc2md github.com/jakebailey/comb > README.md"
