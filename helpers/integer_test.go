package helpers

import (
	"fmt"
	"testing"

	"github.com/jakebailey/comb"
	"github.com/stretchr/testify/assert"
)

func intShouldMatch(t *testing.T, i int64, str string) {
	s := comb.NewStringScanner(str)
	r, _ := IntegerParser().Parse(s)

	assert.Nil(t, r.Err)
	assert.Equal(t, i, r.Int64)
}

func intCheck(t *testing.T, i int64) {
	ui := uint64(i)

	tests := []string{
		fmt.Sprintf("%v", i),
		fmt.Sprintf("%v", ui),
		fmt.Sprintf("%#x", i),
		fmt.Sprintf("%#x", ui),
		fmt.Sprintf("%#o", i),
		fmt.Sprintf("%#o", ui),
	}

	for _, s := range tests {
		intShouldMatch(t, i, s)
	}
}

func TestIntegerParser(t *testing.T) {
	for _, v := range []int64{
		0, 1234, 0xDEADBEEF, -1,
	} {
		intCheck(t, v)
	}
}

func BenchmarkIntegerParser(b *testing.B) {
	b.ReportAllocs()
	p := IntegerParser()

	tests := []string{
		"0",
		"-12345",
		"0xDEADBEEF",
		"0777",
		"0x1234123412341234",
		"-0xFFFF0000",
		"0Xaaaa",
	}

	for _, str := range tests {
		b.Run(str, func(b *testing.B) {
			s := comb.NewStringScanner(str)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				p.Parse(s)
			}
		})
	}
}
