package main

import "fmt"



func add(a int, b int) int {
	return a + b
}

func subtract(a,b int)(result int){
	result=a-b
	return
}

func multiply(a, b int) int {
    return a * b
}

func divide(a,d int) int{
	if d == 0 {
        fmt.Println("Error: Division by zero is not allowed")
        return 0
    }
    return a / d
}


func main() {
	fmt.Println("We are learning about the Function in Go")
	var num1 int
	var num2 int
	fmt.Println("Enter first number:")
	fmt.Scan(&num1)
	fmt.Println("Enter second number:")
	fmt.Scan(&num2)
	fmt.Printf("Addition of %d and %d is:%d\n",num1,num2,add(num1,num2))
	fmt.Printf("Subtraction of %d and %d is:%d\n",num1,num2,subtract(num1,num2))
	fmt.Printf("Subtraction of %d and %d is:%d\n",num1,num2,multiply(num1,num2))
	fmt.Printf("Division of %d and %d is :%d\n",num1,num2,divide(num1,num2))

}
