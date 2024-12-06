package calc

import "fmt"

// Add two numbers
func Add(a, b float64) float64 {
	return a + b
}

// Subtract second number from first
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply two numbers
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide first number by second
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}
