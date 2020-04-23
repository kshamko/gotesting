package fib

import "testing"

func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(10)
	}
}

func BenchmarkFibImproved(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibImproved(10)
	}
}
