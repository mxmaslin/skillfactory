package tests

import (
	"math/rand"
	helpers "github.com/mxmaslin/skillfactory/go/2_algorithms/sorting_algorithms/helpers"
)


func BubbleSort(arr []int) {
	arrLength := len(arr)
	for i := 0; i < arrLength - 1; i++ {
		for j := 0; j < arrLength - i - 1; j++ {
			if arr[j] > arr[j + 1] {
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
			}
		}
	}
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


func InsertionSort(arr []int) {
	j := 0
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j = i - 1
		for j >= 0 && arr[j] > key {
			arr[j + 1] = arr[j]
			j = j - 1
		}
		arr[j + 1] = key
	}
}


func merge(left, right []int) []int {
	merged := make([]int, len(left) + len(right))
	
	i := 0

	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			merged[i] = left[0]
			left = left[1:]
		} else {
			merged[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		merged[i] = left[j]
		i++
	}

	for j := 0; j < len(right); j++ {
		merged[i] = right[j]
		i++
	}

	return merged
}


func MergeSort(arr []int) []int {
	length := len(arr)

	if length == 1 {
		return arr
	}

    center := int(length / 2)

    left := make([]int, center)
    right := make([]int, length - center)

    for i := 0; i < length; i++ {
    	if i < center {
    		left[i] = arr[i]
    	} else {
    		right[i - center] = arr[i]
    	}
    }

    return merge(MergeSort(left), MergeSort(right))
}


func QuickSort(arr []int) []int {
    if len(arr) < 2 {
        return arr
    }
     
    left, right := 0, len(arr) - 1
     
    pivot := rand.Int() % len(arr)
     
    arr[pivot], arr[right] = arr[right], arr[pivot]
     
    for i, _ := range arr {
        if arr[i] < arr[right] {
            arr[left], arr[i] = arr[i], arr[left]
            left++
        }
    }
     
    arr[left], arr[right] = arr[right], arr[left]
     
    QuickSort(arr[:left])
    QuickSort(arr[left+1:])
     
    return arr
}
