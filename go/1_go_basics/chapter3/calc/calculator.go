package calc


type calculator struct {
	number1, number2 float64
	operation string
}


func NewCalculator(number1, number2 float64, operation string) calculator{
	return calculator{
		number1: number1,
		number2: number2,
		operation: operation,
	}
}


func (c calculator) add(number1, number2 float64) float64 {
	return number1 + number2
}


func (c calculator) subtract(number1, number2 float64) float64 {
	return number1 - number2
}


func (c calculator) multiply(number1, number2 float64) float64 {
	return number1 * number2
}


func (c calculator) divide(number1, number2 float64) float64 {
	if number2 == 0 {
		return 0
	}
	return number1 / number2
}


func (c calculator) Calculate(number1, number2 float64, operation string) float64 {
	var result float64
	switch operation{
		case "+":
			result = c.add(number1, number2)
		case "-":
			result = c.subtract(number1, number2)
		case "*":
			result = c.multiply(number1, number2)
		case "/":
			result = c.divide(number1, number2)
		default:
			result = 666
	}
	return result
}
