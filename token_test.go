package comb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func matchingToken(t *testing.T, p Parser, str string) {
	s := NewStringScanner(str)

	r, next := p.Parse(s)

	expected := Result{
		Runes: []rune(str),
	}

	assert.True(t, r.Matched())
	assert.Equal(t, expected, r)
	assert.True(t, next.EOF())
}

func notMatchingToken(t *testing.T, p Parser, str string) {
	s := NewStringScanner(str)

	r, _ := p.Parse(s)

	assert.False(t, r.Matched())
}

func TestSingleToken(t *testing.T) {
	p := StringToken("foobar")

	t.Run("easy", func(t *testing.T) {
		tests := map[string]bool{
			"foobar": true,
			"foo":    false,
			"abcd":   false,
		}

		for s, match := range tests {
			if match {
				matchingToken(t, p, s)
			} else {
				notMatchingToken(t, p, s)
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

	p := StringToken("foobar")
	s := NewStringScanner("foobarbaz")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Parse(s)
	}
}

func TestManyTokens(t *testing.T) {
	p := StringToken("foobar", "foolol", "hello")

	t.Run("easy", func(t *testing.T) {
		tests := map[string]bool{
			"foobar": true,
			"foo":    false,
			"abcd":   false,
			"hello":  true,
			"foolol": true,
		}

		for s, match := range tests {
			if match {
				matchingToken(t, p, s)
			} else {
				notMatchingToken(t, p, s)
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

	p := StringToken("foobar", "foolol", "hello")
	s := NewStringScanner("foobarbaz")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Parse(s)
	}
}
