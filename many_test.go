package comb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMany(t *testing.T) {
	p := Many(
		nil,
		Char('a'),
	)

	t.Run("match", func(t *testing.T) {
		s := NewStringScanner("aaaab")

		r, next := p.Parse(s)

		expected := Result{
			Interface: []Result{
				{Runes: []rune("a")},
				{Runes: []rune("a")},
				{Runes: []rune("a")},
				{Runes: []rune("a")},
			},
		}

		assert.True(t, r.Matched())
		assert.Equal(t, expected, r)
		assert.False(t, next.EOF())
	})

	t.Run("empty", func(t *testing.T) {
		s := NewStringScanner("bbbbbb")

		r, next := p.Parse(s)

		expected := Result{
			Interface: []Result(nil),
		}

		assert.True(t, r.Matched())
		assert.Equal(t, expected, r)
		assert.False(t, next.EOF())
	})
}

func TestManyRunes(t *testing.T) {
	p := ManyRunes(
		Char('a'),
	)

	t.Run("match", func(t *testing.T) {
		s := NewStringScanner("aaaab")

		r, next := p.Parse(s)

		expected := Result{
			Runes: []rune("aaaa"),
		}

		assert.True(t, r.Matched())
		assert.Equal(t, expected, r)
		assert.False(t, next.EOF())
	})

	t.Run("empty", func(t *testing.T) {
		s := NewStringScanner("bbbbbb")

		r, next := p.Parse(s)

		expected := Result{
			Runes: []rune(""),
		}

		assert.True(t, r.Matched())
		assert.Equal(t, expected, r)
		assert.False(t, next.EOF())
	})
}

func TestOnePlus(t *testing.T) {
	p := OnePlus(
		nil,
		Char('a'),
	)

	t.Run("match", func(t *testing.T) {
		s := NewStringScanner("aaaab")

		r, next := p.Parse(s)

		expected := Result{
			Interface: []Result{
				{Runes: []rune("a")},
				{Runes: []rune("a")},
				{Runes: []rune("a")},
				{Runes: []rune("a")},
			},
		}

		assert.True(t, r.Matched())
		assert.Equal(t, expected, r)
		assert.False(t, next.EOF())
	})

	t.Run("no match", func(t *testing.T) {
		s := NewStringScanner("bbbbbb")

		r, next := p.Parse(s)

		assert.False(t, r.Matched())
		assert.EqualError(t, r.Err, "unexpected character 'b'")
		assert.False(t, next.EOF())
	})
}

func TestOnePlusRunes(t *testing.T) {
	p := OnePlusRunes(
		Char('a'),
	)

	t.Run("match", func(t *testing.T) {
		s := NewStringScanner("aaaab")

		r, next := p.Parse(s)

		expected := Result{
			Runes: []rune("aaaa"),
		}

		assert.True(t, r.Matched())
		assert.Equal(t, expected, r)
		assert.False(t, next.EOF())
	})

	t.Run("no match", func(t *testing.T) {
		s := NewStringScanner("bbbbbb")

		r, next := p.Parse(s)

		assert.False(t, r.Matched())
		assert.EqualError(t, r.Err, "unexpected character 'b'")
		assert.False(t, next.EOF())
	})
}
