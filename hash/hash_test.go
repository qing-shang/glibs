package hash

import (
	"testing"
)

func Benchmark_Sum128(b *testing.B) {
	data := "csa15d"
	b.ResetTimer()
	b.StartTimer()

	b.Run("Run.N", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			Sum128String(data)
		}
	})
}

func Benchmark_Sum1281(b *testing.B) {
	data := "csa15d"
	b.ResetTimer()
	b.StartTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Sum128String(data)
	}
}

func Benchmark_Sum32(b *testing.B) {
	data := "csa15d"
	b.ResetTimer()
	b.StartTimer()

	b.Run("Run.N", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			Sum32String(data)
		}
	})
}
