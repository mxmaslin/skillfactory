package algorithms


import (
	// "fmt"
	// "time"
	// int_helper "github.com/mxmaslin/skillfactory/go/2_algorithms/sorting_algorithms/int_helper"
	"math/rand"
	"testing"
	algorithms "github.com/mxmaslin/skillfactory/go/2_algorithms/sorting_algorithms/tests"
)


func generateSlice(max, size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(max*2) - max
	}

	return arr
}


func BenchmarkBubbleSort(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			arr := generateSlice(10, 10)
			b.StartTimer()
			algorithms.BubbleSort(arr)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			arr := generateSlice(100, 1000)
			b.StartTimer()
			algorithms.BubbleSort(arr)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			arr := generateSlice(10000, 100000)
			b.StartTimer()
			algorithms.BubbleSort(arr)
			b.StopTimer()
		}
	})
}
