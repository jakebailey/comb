package combext

import "github.com/jakebailey/comb"

// Digit accepts a single digit from 0 to 9.
func Digit() comb.Parser {
	return comb.CharRange('0', '9')
}

// Digits accepts 1+ digits.
func Digits() comb.Parser {
	return comb.OnePlusRunes(Digit())
}

// AlphaLower accepts characters from a to z.
func AlphaLower() comb.Parser {
	return comb.CharRange('a', 'z')
}

// AlphaUpper accepts characters from A to Z.
func AlphaUpper() comb.Parser {
	return comb.CharRange('A', 'Z')
}

// Alpha accepts characters from either a to z or A to Z.
func Alpha() comb.Parser {
	return comb.Or(
		AlphaLower(),
		AlphaUpper(),
	)
}

// AlphaDigit accepts characters that match Alpha or Digit.
func AlphaDigit() comb.Parser {
	return comb.Or(
		Alpha(),
		Digit(),
	)
}

// Whitespace accepts any of ' ', '\t', '\n', or '\r'.
func Whitespace() comb.Parser {
	return comb.Char(' ', '\t', '\n', '\r')
}

// ManyWhitespace accepts 0+ instances of whitespace.
func ManyWhitespace() comb.Parser {
	return comb.ManyRunes(Whitespace())
}

// OnePlusWhitespace accepts 1+ instances of whitespace.
func OnePlusWhitespace() comb.Parser {
	return comb.ManyRunes(Whitespace())
}
