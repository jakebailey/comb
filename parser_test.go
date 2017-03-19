package comb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRune(t *testing.T) {
	p := Rune('a', 'b', 'c')
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
	assert.EqualError(t, r.Err, "unexpected rune 'd'")
}

func BenchmarkRune(b *testing.B) {
	b.ReportAllocs()

	p := Rune('a', 'b', 'c')
	s := NewStringScanner("abc")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Parse(s)
	}
}
