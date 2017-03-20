package comb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNamed(t *testing.T) {
	p := Named("foobar", Char('a'))
	s := NewStringScanner("abcd")

	r, next := p.Parse(s)

	expected := Result{
		Runes:      []rune("a"),
		ParserName: "foobar",
	}

	assert.True(t, r.Matched())
	assert.Equal(t, expected, r)
	assert.False(t, next.EOF())
}

func TestIgnore(t *testing.T) {
	p := Ignore(Char('a'))
	s := NewStringScanner("abcd")

	r, next := p.Parse(s)

	expected := Result{
		Runes:  []rune("a"),
		Ignore: true,
	}

	assert.True(t, r.Matched())
	assert.Equal(t, expected, r)
	assert.False(t, next.EOF())
}

func TestEOF(t *testing.T) {
	t.Run("EOF", func(t *testing.T) {
		p := EOF()
		s := NewStringScanner("")

		r, next := p.Parse(s)

		expected := Result{}

		assert.True(t, r.Matched())
		assert.Equal(t, expected, r)
		assert.True(t, next.EOF())
	})

	t.Run("not EOF", func(t *testing.T) {
		p := EOF()
		s := NewStringScanner("abcd")

		r, next := p.Parse(s)

		assert.False(t, r.Matched())
		assert.EqualError(t, r.Err, "expected EOF, got 'a'")
		assert.False(t, next.EOF())
	})
}
