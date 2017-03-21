

# comb
`import "github.com/jakebailey/comb"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>
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




## <a name="pkg-index">Index</a>
* [type Parser](#Parser)
  * [func AnyChar() Parser](#AnyChar)
  * [func Char(chars ...rune) Parser](#Char)
  * [func CharRange(from, to rune) Parser](#CharRange)
  * [func CharsIn(s string) Parser](#CharsIn)
  * [func EOF() Parser](#EOF)
  * [func Ignore(parser Parser) Parser](#Ignore)
  * [func Many(combiner ResultCombiner, parser Parser) Parser](#Many)
  * [func ManyRunes(parser Parser) Parser](#ManyRunes)
  * [func Maybe(parser Parser) Parser](#Maybe)
  * [func NotChar(chars ...rune) Parser](#NotChar)
  * [func OnePlus(combiner ResultCombiner, parser Parser) Parser](#OnePlus)
  * [func OnePlusRunes(parser Parser) Parser](#OnePlusRunes)
  * [func Or(parsers ...Parser) Parser](#Or)
  * [func OrLongest(parsers ...Parser) Parser](#OrLongest)
  * [func ParserFunc(fn func(Scanner) (Result, Scanner)) Parser](#ParserFunc)
  * [func Reference(p *Parser) Parser](#Reference)
  * [func Regexp(pattern string) Parser](#Regexp)
  * [func Sequence(combiner ResultCombiner, parsers ...Parser) Parser](#Sequence)
  * [func SequenceRunes(parsers ...Parser) Parser](#SequenceRunes)
  * [func Surround(left, parser, right Parser) Parser](#Surround)
  * [func Tag(tag string, parser Parser) Parser](#Tag)
  * [func Take(n int) Parser](#Take)
  * [func Token(tokens ...string) Parser](#Token)
  * [func TokenRunes(tokens ...[]rune) Parser](#TokenRunes)
* [type Result](#Result)
  * [func Failed(err error) Result](#Failed)
  * [func Failedf(format string, a ...interface{}) Result](#Failedf)
  * [func SliceCombiner(results []Result, begin, end Scanner) Result](#SliceCombiner)
  * [func (r Result) Matched() bool](#Result.Matched)
* [type ResultCombiner](#ResultCombiner)
* [type Scanner](#Scanner)
  * [func NewScanner(s []rune) Scanner](#NewScanner)
  * [func NewStringScanner(s string) Scanner](#NewStringScanner)
  * [func (s Scanner) Between(other Scanner) []rune](#Scanner.Between)
  * [func (s Scanner) Col() int](#Scanner.Col)
  * [func (s Scanner) EOF() bool](#Scanner.EOF)
  * [func (s Scanner) Line() int](#Scanner.Line)
  * [func (s Scanner) Next() (rune, Scanner, error)](#Scanner.Next)


#### <a name="pkg-files">Package files</a>
[char.go](/src/github.com/jakebailey/comb/char.go) [doc.go](/src/github.com/jakebailey/comb/doc.go) [helpers.go](/src/github.com/jakebailey/comb/helpers.go) [many.go](/src/github.com/jakebailey/comb/many.go) [or.go](/src/github.com/jakebailey/comb/or.go) [parser.go](/src/github.com/jakebailey/comb/parser.go) [regexp.go](/src/github.com/jakebailey/comb/regexp.go) [result.go](/src/github.com/jakebailey/comb/result.go) [scanner.go](/src/github.com/jakebailey/comb/scanner.go) [sequence.go](/src/github.com/jakebailey/comb/sequence.go) [token.go](/src/github.com/jakebailey/comb/token.go) 






## <a name="Parser">type</a> [Parser](/src/target/parser.go?s=163:231#L1)
``` go
type Parser interface {
    Parse(s Scanner) (r Result, next Scanner)
}
```
Parser describes comb parsers, which take a scanner,
scan some amount of text, then return a result and
the next scanner.







### <a name="AnyChar">func</a> [AnyChar](/src/target/char.go?s=55:76#L1)
``` go
func AnyChar() Parser
```
AnyChar accepts any single character.


### <a name="Char">func</a> [Char](/src/target/char.go?s=316:347#L9)
``` go
func Char(chars ...rune) Parser
```
Char accepts a single given character.


### <a name="CharRange">func</a> [CharRange](/src/target/char.go?s=1568:1604#L75)
``` go
func CharRange(from, to rune) Parser
```
CharRange accepts chars in an inclusive range.


### <a name="CharsIn">func</a> [CharsIn](/src/target/char.go?s=1940:1969#L93)
``` go
func CharsIn(s string) Parser
```
CharsIn accepts any of the chars in a given string.


### <a name="EOF">func</a> [EOF](/src/target/parser.go?s=1188:1205#L44)
``` go
func EOF() Parser
```
EOF matches only at EOF.


### <a name="Ignore">func</a> [Ignore](/src/target/parser.go?s=998:1031#L35)
``` go
func Ignore(parser Parser) Parser
```
Ignore sets the result of a Parser to be Ignored.


### <a name="Many">func</a> [Many](/src/target/many.go?s=233:289#L1)
``` go
func Many(combiner ResultCombiner, parser Parser) Parser
```
Many looks for a series of 0+ matches of a parser,
then combines the results with a combiner. If combiner is nil,
SliceCombiner is used.

If you only need the runes captured by Many, use TextMany instead.


### <a name="ManyRunes">func</a> [ManyRunes](/src/target/many.go?s=720:756#L23)
``` go
func ManyRunes(parser Parser) Parser
```
ManyRunes looks for a series of 0+ matches of a parser,
then returns the runes captured.


### <a name="Maybe">func</a> [Maybe](/src/target/parser.go?s=1529:1561#L57)
``` go
func Maybe(parser Parser) Parser
```
Maybe tries a parser and returns its result if it matches,
otherwise, it returns an empty result and the original scanner.


### <a name="NotChar">func</a> [NotChar](/src/target/char.go?s=1111:1145#L52)
``` go
func NotChar(chars ...rune) Parser
```
NotChar only accepts a char not given.


### <a name="OnePlus">func</a> [OnePlus](/src/target/many.go?s=1131:1190#L44)
``` go
func OnePlus(combiner ResultCombiner, parser Parser) Parser
```
OnePlus looks for a series of 1+ matches of a parser.

If you only need the runes captured by OnePlus, use TextOnePlus instead.


### <a name="OnePlusRunes">func</a> [OnePlusRunes](/src/target/many.go?s=1686:1725#L73)
``` go
func OnePlusRunes(parser Parser) Parser
```
OnePlusRunes looks for a series of 1+ matches of a parser,
then returns the runes captured.


### <a name="Or">func</a> [Or](/src/target/or.go?s=72:105#L1)
``` go
func Or(parsers ...Parser) Parser
```
Or checks parsers in order, returning the first match.


### <a name="OrLongest">func</a> [OrLongest](/src/target/or.go?s=531:571#L12)
``` go
func OrLongest(parsers ...Parser) Parser
```
OrLongest is like Or, but returns the result of the parser
that captures the most text. Ties are broken by taking the
first result. In order to do this, *every* parser will be run,
so keep that in mind.


### <a name="ParserFunc">func</a> [ParserFunc](/src/target/parser.go?s=433:491#L13)
``` go
func ParserFunc(fn func(Scanner) (Result, Scanner)) Parser
```
ParserFunc turns a parser function into a Parser.


### <a name="Reference">func</a> [Reference](/src/target/parser.go?s=616:648#L19)
``` go
func Reference(p *Parser) Parser
```
Reference takes a pointer to a Parser, and only dereferences it
when Parse is called.


### <a name="Regexp">func</a> [Regexp](/src/target/regexp.go?s=206:240#L1)
``` go
func Regexp(pattern string) Parser
```
Regexp compiles a Go regexp into a parser. If the pattern does
not begin with ^, one will be added, as the parser must begin
with the next rune.


### <a name="Sequence">func</a> [Sequence](/src/target/sequence.go?s=339:403#L1)
``` go
func Sequence(combiner ResultCombiner, parsers ...Parser) Parser
```
Sequence runs multiple parsers in a sequence, combining results
with a combiner function. If combiner is nil, then SliceCombiner
is used. Sequence must allocate a slice of results the same length
as the number of parsers required.

If you only need the runes captured by Sequence, use SequenceRunes instead.


### <a name="SequenceRunes">func</a> [SequenceRunes](/src/target/sequence.go?s=1731:1775#L54)
``` go
func SequenceRunes(parsers ...Parser) Parser
```
SequenceRunes is like Sequence, but does not capture all results,
instead returning the runes between the start and end of the matching
region. Unlike Sequence, this does not allocate anything.
SequenceRunes does not read any results, just the returned scanners,
so cannot respect the Ignored option.


### <a name="Surround">func</a> [Surround](/src/target/sequence.go?s=2199:2247#L76)
``` go
func Surround(left, parser, right Parser) Parser
```
Surround surrounds a parser with two parsers, and returns
the surrounded value. This is equivalent to Sequence with a combiner
which returns the middle result.


### <a name="Tag">func</a> [Tag](/src/target/parser.go?s=778:820#L26)
``` go
func Tag(tag string, parser Parser) Parser
```
Tag sets the tag of a parser's result.


### <a name="Take">func</a> [Take](/src/target/char.go?s=782:805#L32)
``` go
func Take(n int) Parser
```
Take accepts n characters and returns the runes captured.


### <a name="Token">func</a> [Token](/src/target/token.go?s=192:227#L1)
``` go
func Token(tokens ...string) Parser
```
Token accepts the shortest given token. At least one token must
be provided. If more than one token is given, then a trie is used
to check for membership.


### <a name="TokenRunes">func</a> [TokenRunes](/src/target/token.go?s=422:462#L8)
``` go
func TokenRunes(tokens ...[]rune) Parser
```
TokenRunes is like Token, but takes multiple rune slices.





## <a name="Result">type</a> [Result](/src/target/result.go?s=366:516#L1)
``` go
type Result struct {
    Err       error
    Runes     []rune
    Int64     int64
    Float64   float64
    Interface interface{}
    Tag       string
    Ignore    bool
}
```
Result represents the result of a parser.
It supports a range of common values, including a rune slice,
integer and float types used in strconv, as well as an interface{}
for anything not included. Err will be set if a Result is failed.
If your result contains an error that is not a failure, then it should
be placed into Interface.







### <a name="Failed">func</a> [Failed](/src/target/result.go?s=672:701#L15)
``` go
func Failed(err error) Result
```
Failed returns a failed result with a given error.


### <a name="Failedf">func</a> [Failedf](/src/target/result.go?s=972:1024#L23)
``` go
func Failedf(format string, a ...interface{}) Result
```
Failedf returns a failed result in fmt.Errorf form.
fmt.Errorf will not be called until the error is read to prevent
unnecessary computation. This is important, as failed results
can be checked without ever generating an error.


### <a name="SliceCombiner">func</a> [SliceCombiner](/src/target/sequence.go?s=1135:1198#L31)
``` go
func SliceCombiner(results []Result, begin, end Scanner) Result
```
SliceCombiner combines results by returning a Result with the
slice in Interface. If a result is set to be ignored, the result
will not be in the new result slice.





### <a name="Result.Matched">func</a> (Result) [Matched](/src/target/result.go?s=561:591#L10)
``` go
func (r Result) Matched() bool
```
Matched returns true if Err is not nil.




## <a name="ResultCombiner">type</a> [ResultCombiner](/src/target/sequence.go?s=891:960#L26)
``` go
type ResultCombiner func(results []Result, begin, end Scanner) Result
```
ResultCombiner is a function that takes a slice of results
and surrounding scanners and combines them into a single result.










## <a name="Scanner">type</a> [Scanner](/src/target/scanner.go?s=92:162#L1)
``` go
type Scanner struct {
    // contains filtered or unexported fields
}
```
Scanner is an immutable struct which scans over a rune slice.







### <a name="NewScanner">func</a> [NewScanner](/src/target/scanner.go?s=219:252#L4)
``` go
func NewScanner(s []rune) Scanner
```
NewScanner creates a new Scanner from a rune slice.


### <a name="NewStringScanner">func</a> [NewStringScanner](/src/target/scanner.go?s=341:380#L9)
``` go
func NewStringScanner(s string) Scanner
```
NewStringScanner creates a new Scanner from a string.





### <a name="Scanner.Between">func</a> (Scanner) [Between](/src/target/scanner.go?s=1102:1148#L48)
``` go
func (s Scanner) Between(other Scanner) []rune
```
Between returns the slice between two scanners.
s1.Between(s2) returns a slice in the range [s1, s2).




### <a name="Scanner.Col">func</a> (Scanner) [Col](/src/target/scanner.go?s=1340:1366#L58)
``` go
func (s Scanner) Col() int
```
Col returns the current column number, 1 indexed.




### <a name="Scanner.EOF">func</a> (Scanner) [EOF](/src/target/scanner.go?s=933:960#L42)
``` go
func (s Scanner) EOF() bool
```
EOF returns true if the scanner is at EOF, i.e. a call to Next would
return EOF.




### <a name="Scanner.Line">func</a> (Scanner) [Line](/src/target/scanner.go?s=1235:1262#L53)
``` go
func (s Scanner) Line() int
```
Line returns the current line number, 1 indexed.




### <a name="Scanner.Next">func</a> (Scanner) [Next](/src/target/scanner.go?s=553:599#L15)
``` go
func (s Scanner) Next() (rune, Scanner, error)
```
Next scans for the next rune, returning the rune and the next Scanner.
If there are no more runes to scan, io.EOF is returned.








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
