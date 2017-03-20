package main

import (
	"testing"

	"github.com/jakebailey/comb"
)

func BenchmarkCalculator(b *testing.B) {
	b.ReportAllocs()

	test := "(1 + 2 * 3 + 9) * 2 + 1"
	s := comb.NewStringScanner(test)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		expr.Parse(s)
	}
}
