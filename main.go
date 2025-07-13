package main

import (
	"awesomeProject/calc"
	"fmt"
	"strconv"
)

func main() {
	var number1, operator, number2 string
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&number1)

	fmt.Print("Введите орифметический оператор (+, -, *, /):")
	fmt.Scanln(&operator)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&number2)

	num1, _ := strconv.ParseFloat(number1, 64)
	num2, _ := strconv.ParseFloat(number2, 64)

	calculator := calc.NewCalculator()
	result := calculator.Calculate(num1, num2, operator)
	fmt.Printf("Результат: %.2f\n", result)
}
