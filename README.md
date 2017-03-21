# comb parser combinator framework

comb is an efficient parser combinator framework.

comb offers various useful parsers, including nonterminals such as
characters, character ranges, tokens, and regular expressions. Also
included are sequences, repetitions, and optional parsers.

For example, here is a parser that recognizes the decimal, octal,
and hex integers accepted in Go.

```go
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
```

(Though, this is more succinctly expressed with a regular expression,
giving a moderate performance gain.)

comb uses a scanner which traverses a rune slice. All builtin parsers
return results that are simply slices of the original data, keeping copying
to a minimum.

The `combext` package offers other general-use parsers (such as alpha-numeric
characters, whitespace, etc) that may be frequently needed, though not always
used.

## Examples

In the `_examples` directory, you can find examples of comb in use, including
a recursive expression calculator.


## Other libraries

comb takes inspiration from the following Go parser combinator libraries:

- [jmikkola/parsego](https://github.com/jmikkola/parsego)
- [prataprc/goparsec](https://github.com/prataprc/goparsec)

I tried both of these while writing a parser for another project.
They both offer some amount of good usability and performance,
but not enough to my liking. I thought the changes I would
make would be too breaking to turn into a reasonable PR, so here we are.