package main

import "testing"

func BenchmarkFibonacciRecursiveEnd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciRecursiveEnd(45)
	}
}

func BenchmarkFibonacciRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciRecursiveSlow(45)
	}
}

func BenchmarkFibonacciFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciSmart(45)
	}
}
