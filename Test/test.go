package main

import "fmt"

func main() {
	age := 10
	name := "Alice"
	height := 5.6

	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)
	fmt.Println("Height: ", height)
	fmt.Print("Hi ")
	fmt.Print(name)

	
	fmt.Printf("Name: %s\n" ,name)
	fmt.Printf("Age: %d\n",age)
	fmt.Printf("Height: %.3f\n",height)
}
