package main

import "testing"

// Running benchmark test for function performance comparison
// go test -bench=. -benchmem -memprofile mem.prof -cpuprofile cpu.prof -count 10 > 1.bench
// benchstat 1.bench 2.bench
func BenchmarkGenerateLargeString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateLargeString(1000)
	}
}
