package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		word := fmt.Sprint(i)
		if i % 15 == 0 {
			word = "FizzBuzz"
		} else if i % 3 == 0 {
			word = "Fizz"
		} else if i % 5 == 0 {
			word = "Buzz"
		}
		fmt.Println(word)
	}
}
