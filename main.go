package main

import (
	"fmt"
	"go-basics/calculator"
)

func main() {
	sumResult := calculator.Sum(5, 3)
	fmt.Printf("Sum(5, 3) = %f\n", sumResult)

	multiplyResult := calculator.Multiply(5, 3)
	fmt.Printf("Multiply(5, 3) = %f\n", multiplyResult)
}