package main

import (
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	want := math.Sqrt2
	got := sqrt(2, 10)

	if math.Abs(want-got) >= 1e-10 {
		t.Errorf("want:%.10f, got:%.10f", want, got)
	}
}
