package main

import (
	"fmt"
)

func main() {
	//slice
	// numbers:=[]int {0,1,2,3,4,5,6,7}
	// numbers = append(numbers, 56,24,24321,13453421,2345321)
	// fmt.Println("Original Slice:", numbers)
	// fmt.Println("Length of Slice:", len(numbers))
	// fmt.Println("Capacity of Slice:", cap(numbers))
	// numbers = append(numbers, 8)
	// fmt.Println("Appended Slice:", numbers)

	// prices :=[]string{"1","23","4","5","6"}
	// fmt.Println(prices)
	// prices = append(prices, "Hi")
	// fmt.Println(prices)

	// Slice with 0 size
	// names := []string{}
	// fmt.Println(names)
	// fmt.Println(cap(names))

	// number:=[]int{1,2,3}
	// fmt.Println("Slices:",number)
	// fmt.Println("Length:",len(number))
	// fmt.Println("Capacity:",cap(number))

	// Use of make function
	number := make([]int, 3, 5) //make([]datatype,length,capacity)
	number=append(number, 4)
	number=append(number, 5)
	number=append(number, 6)
	number=append(number, 6)
	fmt.Println("Slices:", number)
	fmt.Println("Length:", len(number))
	fmt.Println("Capacity:", cap(number))

}
