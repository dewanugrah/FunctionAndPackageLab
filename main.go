// main.go
package main

import (
    "fmt"
    "FunctionAndPackageLab/mathops"
)

func main() {
    // Test the functions from the mathops package

	// using add func from the mathops package
    result := mathops.Add(10, 5)
    fmt.Println("Addition:", result)

	// using subtract func from the mathops package
    result = mathops.Subtract(10, 5)
    fmt.Println("Subtraction:", result)

	// using multiply func from the mathops package
    result = mathops.Multiply(10, 5)
    fmt.Println("Multiplication:", result)

	// using divide func from the mathops package
    quotient, remainder := mathops.Divide(10, 3)
    fmt.Printf("Division: Quotient - %d, Remainder - %d\n", quotient, remainder)

	// using power func from the mathops package
	result = mathops.Power(2, 3)
    fmt.Println("Power:", result)
}