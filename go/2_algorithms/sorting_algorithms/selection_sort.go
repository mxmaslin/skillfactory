package main

import (
	"fmt"
	"math/rand"
	"time"
	int_helper "github.com/mxmaslin/skillfactory/go/2_algorithms/sorting_algorithms/int_helper"
)


func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	arr := make([]int, 10)
	for i := range arr {
		arr[i] = rand.Intn(20) - 10 // ограничиваем случайно значение от [-100;100]
	}

	fmt.Println(arr)
	selectionSortBidirectional(arr)
	fmt.Println(arr)
}


func selectionSort(arr []int) {
	if len(arr) > 1 {
		for i := 0; i < len(arr); i++ {
			smallest := int_helper.GetMaxInt()
			for j := i; j < len(arr); j++ {
				if arr[j] < smallest {
					smallest = arr[j]
					arr[i], arr[j] = arr[j], arr[i]				
				}
			}
		}

	}
}


func selectionSortDesc(arr []int) {
	if len(arr) > 1 {
		for i := 0; i < len(arr); i++ {
			largest := int_helper.GetMinInt()
			for j := i; j < len(arr); j++ {
				if arr[j] > largest {
					largest = arr[j]
					arr[i], arr[j] = arr[j], arr[i]				
				}
			}
		}

	}
}


func selectionSortBidirectional(arr []int) {
	if len(arr) > 1 {
		for i := 0; i < len(arr); i++ {
			smallest := int_helper.GetMaxInt()
			largest := int_helper.GetMinInt()

			for j := i; j < len(arr); j++ {
				if arr[j] < smallest {
					smallest = arr[j]
					arr[i], arr[j] = arr[j], arr[i]
				}
				if arr[len(arr) - j - 1] > largest {
					largest = arr[len(arr) - j - 1]
					arr[len(arr) - i - 1], arr[len(arr) - j - 1] = arr[len(arr) - j - 1], arr[len(arr) - i - 1]
				}
				fmt.Println(i, j)
			}
			if arr[i] >= arr[len(arr) / 2] {
				break
			}	
		}
	}
}
