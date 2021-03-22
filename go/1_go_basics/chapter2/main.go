package main

import (
	"fmt"
)

func main() {
	var number1 float32
	var number2 float32
	var operation string
	var result float32
	var output string = "Результат %s: %v"
	
	fmt.Scanln(&number1)
	fmt.Scanln(&operation)
	fmt.Scanln(&number2)
	
	switch operation {
		case "+":
			result = number1 + number2
			output = fmt.Sprintf(output, "сложения", result)
			fmt.Println(output)

		case "-":
			result = number1 - number2
			output = fmt.Sprintf(output, "вычитания", result)
			fmt.Println(output)

		case "*":
			result = number1 * number2
			output = fmt.Sprintf(output, "умножения", result)
			fmt.Println(output)

		case "/":
			if number2 == 0 {
				fmt.Println("Ошибка: деление на ноль")
				break
			}
			result = number1 / number2
			output = fmt.Sprintf(output, "деления", result)
			fmt.Println(output)								
	}
}