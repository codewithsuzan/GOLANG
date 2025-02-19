package main

import "fmt"

func main() {
	// x := 10
	// if x == 10 {
	// 	fmt.Println("X must be 10")
	// }else{
	// 	fmt.Println("X is not 10")
	// }

	var age int
	fmt.Println("Enter your age:")
	fmt.Scan(&age) // Correct syntax, no need for "%d"

	if age >= 18 {
		fmt.Println("You can vote")
	} else {
		fmt.Println("You can't vote")
	}
}
