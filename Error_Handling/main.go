package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Denominator mustn't be negative")
	}
	return a / b, nil
}

func main() {
	fmt.Println("Error handling in Go Language")
	ans, _ := divide(6, 0)
	fmt.Println("Division of two numbers is ", ans)
}
