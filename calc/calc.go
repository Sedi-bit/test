package calc

type calculator struct{}

func NewCalculator() *calculator {
	return &calculator{}
}

func (c *calculator) add(a, b float64) float64 {
	return a + b
}

func (c *calculator) subtract(a, b float64) float64 {
	return a - b
}

func (c *calculator) multiply(a, b float64) float64 {
	return a * b
}

func (c *calculator) divide(a, b float64) float64 {
	if b == 0 {
		panic("Ошибка, деление на 0")
	}
	return a / b
}

func (c *calculator) Calculate(a, b float64, operator string) float64 {

	switch operator {
	case "+":
		return c.add(a, b)
	case "-":

		return c.subtract(a, b)
	case "*":
		return c.multiply(a, b)
	case "/":

		return c.divide(a, b)
	default:
		panic("Ошибка: неизвестный оператор")

	}
}
