package calc

import (
	"fmt"
)

type calculator struct {
	number1, number2 float64
	operation string
}


func(c *calculator) SetCalculator(number1, number2 float64, operation string) {
	c.number1 = number1
	c.number2 = number2
	c.operation = operation
}


func(c calculator) Calculate(number1, number2 float64, operation string) float64 {
	var result float64
	switch operation{
		case "+":
			result = number1 + number2
		case "-":
			result = number1 - number2
		case "*":
			result = number1 * number2
		case "/":
			if number2 == 0 {
				fmt.Println("Ошибка деления на ноль")
				break
			}
			result = number1 / number2
	}
	return result
}




// Добавьте неэкспортируемые (приватные) методы для каждой операции (сложения, вычитания, умножения, деления). Каждый приватный метод должен принимать на вход 2 числа типа float64 и возвращать значение типа float64. (В методе деления чисел должны быть проверка делителя на равенство 0).

// Метод Calculate должен в зависимости от переданного оператора вызывать один из приватных методов. (Если в качестве оператора передан +, то должен быть вызван приватный метод для сложения чисел).

// Создайте пакет main c функцией main(). Внутри функции main необходимо из консоли считать 2 числа и оператор (как вы это делали в задании предыдущего модуля). Затем создать экземпляр структуры calculator из пакета calc и вызвать метод Calculate, передав ему полученные из консоли значения.

// Полученный из метода Calculate результат нужно распечатать в консоль.
