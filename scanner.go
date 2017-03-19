package comb

import "io"

// Scanner is an immutable struct which scans over a rune slice.
type Scanner struct {
	runes []rune
	i     int
	line  int
	col   int
}

// NewScanner creates a new Scanner from a rune slice.
func NewScanner(s []rune) Scanner {
	return Scanner{runes: s}
}

// NewStringScanner creates a new Scanner from a string.
func NewStringScanner(s string) Scanner {
	return Scanner{runes: []rune(s)}
}

// Next scans for the next rune, returning the rune and the next Scanner.
// If there are no more runes to scan, io.EOF is returned.
func (s Scanner) Next() (rune, Scanner, error) {
	if s.EOF() {
		return 0, s, io.EOF
	}

	r := s.runes[s.i]

	col := s.col
	line := s.line

	if r == '\n' {
		line++
		col = 0
	} else {
		col++
	}

	return r, Scanner{
		runes: s.runes,
		i:     s.i + 1,
		line:  line,
		col:   col,
	}, nil
}

// EOF returns true if the scanner is at EOF.
func (s Scanner) EOF() bool {
	return s.i >= len(s.runes)
}

// Range returns the slice between two scanners.
// s1.Range(s2) returns a slice in the range [s1, s2).
func (s Scanner) Range(other Scanner) []rune {
	return s.runes[s.i:other.i]
}

// Line returns the current line number, 1 indexed.
func (s Scanner) Line() int {
	return s.line + 1
}

// Col returns the current column number, 1 indexed.
func (s Scanner) Col() int {
	return s.col + 1
}
