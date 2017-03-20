package comb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSequence(t *testing.T) {
	t.Run("match", func(t *testing.T) {
		p := Sequence(
			nil,
			Char('a'),
			Char('b'),
			Char('c'),
		)
		s := NewStringScanner("abcd")

		r, next := p.Parse(s)

		expected := Result{
			Interface: []Result{
				{Runes: []rune("a")},
				{Runes: []rune("b")},
				{Runes: []rune("c")},
			},
		}

		assert.True(t, r.Matched())
		assert.Equal(t, expected, r)
		assert.False(t, next.EOF())
	})

	t.Run("no match", func(t *testing.T) {
		p := Sequence(
			nil,
			Char('a'),
			Char('b'),
			Char('c'),
		)
		s := NewStringScanner("aaaa")

		r, _ := p.Parse(s)

		assert.False(t, r.Matched())
		assert.EqualError(t, r.Err, "unexpected character 'a'")
	})

	t.Run("ignored", func(t *testing.T) {
		p := Sequence(
			nil,
			Char('a'),
			Ignore(Char('b')),
			Char('c'),
		)
		s := NewStringScanner("abcd")

		r, next := p.Parse(s)

		expected := Result{
			Interface: []Result{
				{Runes: []rune("a")},
				{Runes: []rune("c")},
			},
		}

		assert.True(t, r.Matched())
		assert.Equal(t, expected, r)
		assert.False(t, next.EOF())
	})
}

func BenchmarkSequence(b *testing.B) {
	b.ReportAllocs()

	p := Sequence(
		nil,
		Char('a'),
		Char('b'),
		Char('c'),
	)
	s := NewStringScanner("abcd")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Parse(s)
	}
}

func TestTextSequence(t *testing.T) {
	t.Run("match", func(t *testing.T) {
		p := TextSequence(
			Char('a'),
			Char('b'),
			Char('c'),
		)
		s := NewStringScanner("abcd")

		r, next := p.Parse(s)

		expected := Result{
			Runes: []rune("abc"),
		}

		assert.True(t, r.Matched())
		assert.Equal(t, expected, r)
		assert.False(t, next.EOF())
	})

	t.Run("no match", func(t *testing.T) {
		p := TextSequence(
			Char('a'),
			Char('b'),
			Char('c'),
		)
		s := NewStringScanner("aaa")

		r, _ := p.Parse(s)

		assert.False(t, r.Matched())
		assert.EqualError(t, r.Err, "unexpected character 'a'")
	})
}

func BenchmarkTextSequence(b *testing.B) {
	b.ReportAllocs()

	p := TextSequence(
		Char('a'),
		Char('b'),
		Char('c'),
	)
	s := NewStringScanner("abcd")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p.Parse(s)
	}
}

func TestSurround(t *testing.T) {
	p := Surround(
		Char('('),
		StringToken("foo", "bar", "baz"),
		Char(')'),
	)

	t.Run("simple", func(t *testing.T) {
		s := NewStringScanner("(foo)aaa")

		r, next := p.Parse(s)

		expected := Result{
			Runes: []rune("foo"),
		}

		assert.True(t, r.Matched())
		assert.Equal(t, expected, r)
		assert.False(t, next.EOF())
	})

	t.Run("no match", func(t *testing.T) {
		s := NewStringScanner("(asd)aaa")

		r, _ := p.Parse(s)

		assert.False(t, r.Matched())
		assert.EqualError(t, r.Err, "unexpected character 'a'")
	})

	t.Run("nested", func(t *testing.T) {
		var p Parser
		p = Or(
			StringToken("foo", "bar", "baz"),
			Surround(
				Char('('),
				Reference(&p),
				Char(')'),
			),
		)

		s := NewStringScanner("((((baz))))aaa")

		r, next := p.Parse(s)

		expected := Result{
			Runes: []rune("baz"),
		}

		assert.True(t, r.Matched())
		assert.Equal(t, expected, r)
		assert.False(t, next.EOF())
	})
}
