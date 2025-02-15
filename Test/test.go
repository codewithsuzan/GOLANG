package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// age := 10
	// name := "Alice"
	// height := 5.6

	// fmt.Println("Name: ", name)
	// fmt.Println("Age: ", age)
	// fmt.Println("Height: ", height)
	// fmt.Print("Hi ")
	// fmt.Print(name)

	// fmt.Printf("Name: %s\n" ,name)
	// fmt.Printf("Age: %d\n",age)
	// fmt.Printf("Height: %.3f\n",height)



	
	
	//take input from user

	fmt.Println("Enter your name:")
	// var name string
	// fmt.Scan(&name)
	// fmt.Printf("Hello %s!\n", name)
	// fmt.Println("Length of your name:",len(name))

	reader:=bufio.NewReader(os.Stdin)
	name,_:=reader.ReadString('\n')
	fmt.Printf("Hello %s\n", name)
}
