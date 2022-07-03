package sort

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func init() {
	// set seed first to generate same slices for all benchmarks
	rand.Seed(time.Now().Unix())
}

func slice(length int) []int {
	var slice []int
	for i := 0; i < length; i++ {
		slice = append(slice, rand.Int())
	}
	return slice
}

func BenchmarkBubble100(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bubble(slice(100))
	}
}

func BenchmarkMerge100(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mergesort(slice(100))
	}
}

func BenchmarkDefault100(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Ints(slice(100))
	}
}

func BenchmarkBubble100000(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bubble(slice(100000))
	}
}

func BenchmarkMerge100000(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mergesort(slice(100000))
	}
}

func BenchmarkDefault100000(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Ints(slice(100000))
	}
}
