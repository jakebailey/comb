package comb

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnyChar(t *testing.T) {
	p := AnyChar()
	s := NewStringScanner("abc")
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
	assert.Equal(t, io.EOF, r.Err)
}

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

func TestTake(t *testing.T) {
	t.Run("match", func(t *testing.T) {
		p := Take(3)
		s := NewStringScanner("abcd")

		r, next := p.Parse(s)

		expected := Result{
			Runes: []rune("abc"),
		}

		assert.True(t, r.Matched())
		assert.Equal(t, expected, r)
		assert.False(t, next.EOF())
	})

	t.Run("EOF", func(t *testing.T) {
		p := Take(3)
		s := NewStringScanner("a")

		r, next := p.Parse(s)

		expected := Result{
			Err: io.EOF,
		}

		assert.False(t, r.Matched())
		assert.Equal(t, expected, r)
		assert.True(t, next.EOF())
	})
}

func TestNotChar(t *testing.T) {
	p := NotChar('a', 'b', 'c')

	t.Run("match", func(t *testing.T) {
		s := NewStringScanner("def")

		r, next := p.Parse(s)

		expected := Result{
			Runes: []rune("d"),
		}

		assert.True(t, r.Matched())
		assert.Equal(t, expected, r)
		assert.False(t, next.EOF())
	})

	t.Run("no match", func(t *testing.T) {
		s := NewStringScanner("abc")

		r, _ := p.Parse(s)

		assert.False(t, r.Matched())
		assert.EqualError(t, r.Err, "unexpected character 'a'")
	})

	t.Run("EOF", func(t *testing.T) {
		s := NewStringScanner("")

		r, _ := p.Parse(s)

		assert.False(t, r.Matched())
		assert.Equal(t, r.Err, io.EOF)
	})
}
