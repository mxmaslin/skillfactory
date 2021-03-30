package main

import (
	"fmt"
	calc "github.com/mxmaslin/skillfactory/go/1_go_basics/chapter3/calc"
)

func main() {

	var number1 float64
	var number2 float64
	var operation string
	var calculated string = "Результат: %v\n"
	
	fmt.Scanln(&number1)
	fmt.Scanln(&operation)
	fmt.Scanln(&number2)


	c := calc.NewCalculator(number1, number2, operation)
	result := c.Calculate(number1, number2, operation)

	fmt.Printf(calculated, result)

	// fmt.Printf(output, calculated)
}