package sorting

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	arr := make([]int, 50)
	for i := range arr {
		arr[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}

	fmt.Println(arr)
	arr = MergeSort(arr)
	fmt.Println(arr)
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
