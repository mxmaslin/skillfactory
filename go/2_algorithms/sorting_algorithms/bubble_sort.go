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
	arr := make([]int, 10)
	for i := range arr {
		arr[i] = rand.Intn(20) - 10 // ограничиваем случайно значение от [-100;100]
	}

	fmt.Println(arr)
	BubbleSort(arr)
	fmt.Println(arr)
}


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
