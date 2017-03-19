package comb

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanner(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		runes := []rune("Hello, 世界")
		s := NewScanner(runes)

		var r rune
		next := s
		var err error

		assert.False(t, s.EOF())

		for i := 0; i < len(runes)-1; i++ {
			r, next, err = next.Next()
			assert.Equal(t, r, runes[i])
			assert.False(t, next.EOF())
			assert.Nil(t, err)
		}

		r, next, err = next.Next()
		assert.Equal(t, '界', r)
		assert.True(t, next.EOF())
		assert.Nil(t, err)

		_, _, err = next.Next()
		assert.Equal(t, io.EOF, err)

		rng := s.Between(next)
		assert.Equal(t, runes, rng)
	})

	t.Run("EOF", func(t *testing.T) {
		s := NewStringScanner("ab\nc")

		var r rune
		var err error

		assert.False(t, s.EOF())

		assert.Equal(t, 1, s.Line())
		assert.Equal(t, 1, s.Col())
		r, s, err = s.Next()
		assert.Equal(t, 'a', r)
		assert.False(t, s.EOF())
		assert.Nil(t, err)

		assert.Equal(t, 1, s.Line())
		assert.Equal(t, 2, s.Col())
		r, s, err = s.Next()
		assert.Equal(t, 'b', r)
		assert.False(t, s.EOF())
		assert.Nil(t, err)

		assert.Equal(t, 1, s.Line())
		assert.Equal(t, 3, s.Col())
		r, s, err = s.Next()
		assert.Equal(t, '\n', r)
		assert.False(t, s.EOF())
		assert.Nil(t, err)

		assert.Equal(t, 2, s.Line())
		assert.Equal(t, 1, s.Col())
		r, s, err = s.Next()
		assert.Equal(t, 'c', r)
		assert.True(t, s.EOF())
		assert.Nil(t, err)

		assert.Equal(t, 2, s.Line())
		assert.Equal(t, 2, s.Col())
		r, s, err = s.Next()
		assert.True(t, s.EOF())
		assert.Equal(t, io.EOF, err)
	})
}

func BenchmarkScanner(b *testing.B) {
	b.ReportAllocs()

	next := NewStringScanner("Hello, 世界")
	var err error

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for err != nil {
			_, next, err = next.Next()
		}
	}
}
