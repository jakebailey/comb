package comb

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChar(t *testing.T) {
	t.Run("unexpected rune", func(t *testing.T) {
		p := Char('a', 'b', 'c')
		s := NewStringScanner("abcd")
		var r Result

		r, s = p.Parse(s)
		assert.True(t, r.Matched())
		assert.Equal(t, 'a', r.Runes[0])

		r, s = p.Parse(s)
		assert.True(t, r.Matched())
		assert.Equal(t, 'b', r.Runes[0])

		r, s = p.Parse(s)
		assert.True(t, r.Matched())
		assert.Equal(t, 'c', r.Runes[0])

		r, s = p.Parse(s)
		assert.False(t, r.Matched())
		assert.EqualError(t, r.Err, "unexpected character 'd'")
	})

	t.Run("EOF", func(t *testing.T) {
		p := Char('a', 'b', 'c')
		s := NewStringScanner("a")
		var r Result

		r, s = p.Parse(s)
		assert.True(t, r.Matched())
		assert.Equal(t, 'a', r.Runes[0])

		r, s = p.Parse(s)
		assert.False(t, r.Matched())
		assert.Equal(t, io.EOF, r.Err)
	})
}

func BenchmarkChar(b *testing.B) {
	b.ReportAllocs()

	p := Char('a', 'b', 'c')
	s := NewStringScanner("abc")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Parse(s)
	}
}
