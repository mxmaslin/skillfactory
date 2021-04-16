package main

import "fmt"


func factorial_iterative(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func factorial_recursive(n int) int {
	if n < 2 {
		return n
	}
	return n * factorial_recursive(n - 1)
}


func main() {
	fmt.Println(factorial_iterative(5))
	fmt.Println(factorial_recursive(5))

}