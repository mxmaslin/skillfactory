package main

import "fmt"


func isPrime(digit int) bool {
	top := digit / 2
	for i := 2; i <= top; i++ {
		remainder := digit % i
		if remainder == 0 {
			return false
		}
	}
	return true
}


func main() {
	primesAmount := 20
	startDigit := 2

	for i := 1; i <= primesAmount; {
		if isPrime(startDigit) == true {
			fmt.Println(startDigit)
			i++
		}
		startDigit++
	}
}