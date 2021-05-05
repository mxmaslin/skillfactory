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

	QuickSort(arr)

	fmt.Println(arr)
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