package main

import (
	"fmt"
	"goproject/calculator"
)

func main() {
	a, b := 10.0, 3.0

	fmt.Printf("%.2f + %.2f = %.2f\n", a, b, calculator.Add(a, b))
	fmt.Printf("%.2f - %.2f = %.2f\n", a, b, calculator.Subtract(a, b))
	fmt.Printf("%.2f * %.2f = %.2f\n", a, b, calculator.Multiply(a, b))

	result, err := calculator.Divide(a, b)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("%.2f / %.2f = %.2f\n", a, b, result)
	}

	_, err = calculator.Divide(a, 0)
	if err != nil {
		fmt.Printf("%.2f / 0 = Error: %v\n", a, err)
	}
}
