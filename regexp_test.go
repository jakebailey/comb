package comb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegexp(t *testing.T) {
	p := Regexp(`[^\.]+`)

	t.Run("match", func(t *testing.T) {
		s := NewStringScanner("Hello, 世界.")

		r, next := p.Parse(s)

		expected := Result{
			Runes: []rune("Hello, 世界"),
		}

		assert.True(t, r.Matched())
		assert.Equal(t, expected, r)
		assert.False(t, next.EOF())
	})

	t.Run("no match", func(t *testing.T) {
		s := NewStringScanner("... Hello, 世界.")

		r, _ := p.Parse(s)

		assert.False(t, r.Matched())
		assert.EqualError(t, r.Err, `regexp "[^\\.]+" did not match`)
	})

	t.Run("empty", func(t *testing.T) {
		p := Regexp("")
		s := NewStringScanner("Hello, 世界.")

		r, next := p.Parse(s)

		expected := Result{
			Runes: []rune{},
		}

		assert.True(t, r.Matched())
		assert.Equal(t, expected, r)
		assert.False(t, next.EOF())
	})
}

func BenchmarkRegexp(b *testing.B) {
	b.ReportAllocs()

	p := Regexp(`[^\.]+`)
	s := NewStringScanner("Hello, 世界.")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Parse(s)
	}
}

func BenchmarkRegexpManyTokens(b *testing.B) {
	b.ReportAllocs()

	p := Regexp(`(foobar|foolol|hello)`)
	s := NewStringScanner("foobarbaz")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Parse(s)
	}
}
