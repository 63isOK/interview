package main

import (
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	check(t, math.Sqrt2, sqrt(2, 10), 1e-10)
	check(t, 0, sqrt(0, 10), 1e-10)
}

func check(t *testing.T, want, got, precision float64) {
	if math.Abs(want-got) >= precision {
		t.Errorf("want:%.10f, got:%.10f", want, got)
	}
}

func BenchmarkSqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sqrt(90, 10)
	}
}

func BenchmarkStdSqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Sqrt(90)
	}
}
