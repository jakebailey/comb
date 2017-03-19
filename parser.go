package comb

// Parser describes comb parsers, which take a scanner,
// scan some amount of text, then return the next scanner
// and a result.
type Parser interface {
	Parse(s Scanner) (next Scanner, r Result)
}
