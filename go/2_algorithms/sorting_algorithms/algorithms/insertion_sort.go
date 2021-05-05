package algorithms

import (
	"fmt"
	"math/rand"
	"time"
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
	InsertionSort(arr)
	fmt.Println(arr)
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
