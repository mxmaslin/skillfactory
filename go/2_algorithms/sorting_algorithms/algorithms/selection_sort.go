package algorithms


import (
	"fmt"
	"math/rand"
	"time"
	helpers "github.com/mxmaslin/skillfactory/go/2_algorithms/sorting_algorithms/helpers"
)


func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func RunSelectionSort() {
	arr := make([]int, 10)
	for i := range arr {
		arr[i] = rand.Intn(20) - 10 // ограничиваем случайно значение от [-100;100]
	}

	fmt.Println(arr)
	SelectionSortBidirectional(arr)
	fmt.Println(arr)
}


func SelectionSort(arr []int) {
	if len(arr) > 1 {
		for i := 0; i < len(arr); i++ {
			smallest := helpers.GetMaxInt()
			for j := i; j < len(arr); j++ {
				if arr[j] < smallest {
					smallest = arr[j]
					arr[i], arr[j] = arr[j], arr[i]				
				}
			}
		}

	}
}


func SelectionSortDesc(arr []int) {
	if len(arr) > 1 {
		for i := 0; i < len(arr); i++ {
			largest := helpers.GetMinInt()
			for j := i; j < len(arr); j++ {
				if arr[j] > largest {
					largest = arr[j]
					arr[i], arr[j] = arr[j], arr[i]				
				}
			}
		}

	}
}


func SelectionSortBidirectional(arr []int) {
	if len(arr) > 1 {
		for i := 0; i < len(arr); i++ {
			smallest := helpers.GetMaxInt()
			largest := helpers.GetMinInt()

			for j := i; j < len(arr); j++ {
				if arr[j] < smallest {
					smallest = arr[j]
					arr[i], arr[j] = arr[j], arr[i]
				}
				if arr[len(arr) - j - 1] > largest {
					largest = arr[len(arr) - j - 1]
					arr[len(arr) - i - 1], arr[len(arr) - j - 1] = arr[len(arr) - j - 1], arr[len(arr) - i - 1]
				}
			}
			if arr[i] >= arr[len(arr) / 2] {
				break
			}	
		}
	}
}
