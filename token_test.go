package comb

import (
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func matchingToken(t *testing.T, p Parser, str string) {
	s := NewStringScanner(str)

	r, _ := p.Parse(s)

	expected := Result{
		Runes: []rune(str),
	}

	assert.True(t, r.Matched())
	assert.Equal(t, expected, r)
}

func notMatchingToken(t *testing.T, p Parser, str, prefix string) {
	s := NewStringScanner(str)

	r, _ := p.Parse(s)

	assert.False(t, r.Matched())

	if prefix != "" {
		errStr := fmt.Sprintf("'%s' is not a prefix of any token", prefix)
		assert.EqualError(t, r.Err, errStr)
	} else {
		assert.Equal(t, r.Err, io.EOF)
	}
}

func TestBadToken(t *testing.T) {
	assert.Panics(t, func() {
		Token()
	})
	assert.Panics(t, func() {
		Token()
	})
	assert.Panics(t, func() {
		TokenRunes([]rune(""))
	})
	assert.Panics(t, func() {
		Token("")
	})

	assert.Panics(t, func() {
		TokenRunes([]rune("foobar"), []rune(""), []rune("asdf"))
	})
	assert.Panics(t, func() {
		Token("foobar", "", "asdf")
	})
}

func TestSingleToken(t *testing.T) {
	p := Token("foobar")

	t.Run("easy", func(t *testing.T) {
		tests := map[string]string{
			"foobar": "foobar", // parses whole string
			"foo":    "",       // ends with EOF
			"abcd":   "a",      // stops at 'a'
		}

		for s, prefix := range tests {
			if s == prefix {
				matchingToken(t, p, s)
			} else {
				notMatchingToken(t, p, s, prefix)
			}
		}
	})

	t.Run("substring", func(t *testing.T) {
		s := NewStringScanner("foobarbaz")

		r, next := p.Parse(s)

		expected := Result{
			Runes: []rune("foobar"),
		}

		assert.True(t, r.Matched())
		assert.Equal(t, expected, r)
		assert.False(t, next.EOF())
	})
}

func BenchmarkSingleToken(b *testing.B) {
	b.ReportAllocs()

	p := Token("foobar")
	s := NewStringScanner("foobarbaz")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Parse(s)
	}
}

func TestManyTokens(t *testing.T) {
	p := Token("foobar", "foolol", "hello")

	t.Run("easy", func(t *testing.T) {
		tests := map[string]string{
			"foobar": "foobar",
			"foo":    "",
			"abcd":   "a",
			"hello":  "hello",
			"foolol": "foolol",
		}

		for s, prefix := range tests {
			if s == prefix {
				matchingToken(t, p, s)
			} else {
				notMatchingToken(t, p, s, prefix)
			}
		}
	})

	t.Run("substring", func(t *testing.T) {
		s := NewStringScanner("foobarbaz")

		r, next := p.Parse(s)

		expected := Result{
			Runes: []rune("foobar"),
		}

		assert.True(t, r.Matched())
		assert.Equal(t, expected, r)
		assert.False(t, next.EOF())
	})
}

func BenchmarkManyTokens(b *testing.B) {
	b.ReportAllocs()

	p := Token("foobar", "foolol", "hello")
	s := NewStringScanner("foobarbaz")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Parse(s)
	}
}
