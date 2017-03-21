package comb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOr(t *testing.T) {
	p := Or(
		Char('a'),
		Char('b'),
		Token("foobar", "fizzbuzz", "hello"),
		Char('f'),
	)

	t.Run("match first", func(t *testing.T) {
		s := NewStringScanner("abc")

		r, next := p.Parse(s)

		expected := Result{
			Runes: []rune("a"),
		}

		assert.True(t, r.Matched())
		assert.Equal(t, expected, r)
		assert.False(t, next.EOF())
	})

	t.Run("match second", func(t *testing.T) {
		s := NewStringScanner("baaaa")

		r, next := p.Parse(s)

		expected := Result{
			Runes: []rune("b"),
		}

		assert.True(t, r.Matched())
		assert.Equal(t, expected, r)
		assert.False(t, next.EOF())
	})

	t.Run("match last", func(t *testing.T) {
		s := NewStringScanner("foobarbaz")

		r, next := p.Parse(s)

		expected := Result{
			Runes: []rune("foobar"),
		}

		assert.True(t, r.Matched())
		assert.Equal(t, expected, r)
		assert.False(t, next.EOF())
	})

	t.Run("no match", func(t *testing.T) {
		s := NewStringScanner("1234")

		r, _ := p.Parse(s)

		assert.False(t, r.Matched())
		assert.EqualError(t, r.Err, "no parser matched")
	})
}

func BenchmarkOr(b *testing.B) {
	b.ReportAllocs()

	p := Or(
		Char('a'),
		Char('b'),
		Token("foobar", "fizzbuzz", "hello"),
		Char('f'),
	)
	s := NewStringScanner("foobar")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Parse(s)
	}
}

func TestLongestOr(t *testing.T) {
	p := LongestOr(
		Char('a'),
		Char('f'),
		Token("foobar", "fizzbuzz", "hello"),
		Char('f'),
	)

	t.Run("match", func(t *testing.T) {
		s := NewStringScanner("foobarZZZZ")

		r, next := p.Parse(s)

		expected := Result{
			Runes: []rune("foobar"),
		}

		assert.True(t, r.Matched())
		assert.Equal(t, expected, r)
		assert.False(t, next.EOF())
	})

	t.Run("no match", func(t *testing.T) {
		s := NewStringScanner("1234")

		r, _ := p.Parse(s)

		assert.False(t, r.Matched())
		assert.EqualError(t, r.Err, "no parser matched")
	})
}

func BenchmarkLongestOr(b *testing.B) {
	b.ReportAllocs()

	p := LongestOr(
		Char('a'),
		Char('f'),
		Token("foobar", "fizzbuzz", "hello"),
		Char('f'),
	)
	s := NewStringScanner("foobar")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Parse(s)
	}
}
