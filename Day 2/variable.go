package main

import "fmt"

func main() {
	//variable with different data types

	// A shorthand way to declare and initialize variables
	// person:="Sujan"
	person := true
	age := 18
	fmt.Println(person, age)

	// var (
	// 	num =10
	// 	sentence ="Hello Go"
	// 	boolean =true
	// )

	// fmt.Println("num:",num)
	// fmt.Println("sentence:", sentence)
	// fmt.Println("boolean:", boolean)

	// var a int = 10
	// var b float32 = 0.55
	// var c string = "Hello, World!"
	// var d bool = true
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)

	var a int     // 0
	var b float64 // 0.0
	var c string  // ""
	var d bool    // false
	fmt.Println(a, b, c, d)

}
