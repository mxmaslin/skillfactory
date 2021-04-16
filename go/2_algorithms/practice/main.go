package main

import "fmt"


func checkSliceIsSorted(a []int) bool {
	sorted := true
	if len(a) == 0 {
		return sorted
	}
	current := a[0]

	for _, e := range a[1:] {
		if e > current {
			current = e
		} else {
			sorted = false
			break
		}
	}
	return sorted
}


func bubble_sort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := 0; j <= i; j++ {
			if a[j] > a[i] {
				a[i], a[j] = a[j], a[i]
			}
		}

	}
	return a
}


func main() {
	// a1 := []int{0, 1, 4, 100}
	// a2 := []int{0, 2, 1, 100}

	// fmt.Println(checkSliceIsSorted(a1))
	// fmt.Println(checkSliceIsSorted(a2))


	a1 := []int {0, 1, 2, 3, 4}
	fmt.Println(bubble_sort(a1))


	a2 := []int {1, 4, 2, 3, 1}
	fmt.Println(bubble_sort(a2))


	a3 := []int {9, 8, 7, 6, 5, 4, 100}
	fmt.Println(bubble_sort(a3))
}
