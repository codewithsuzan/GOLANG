package main

import "fmt"

func main() {
	fmt.Println("We are learning about for loop in go")
	// for i:=0;i<=10;i++ {
	// 	fmt.Printf("Number :%d\n",i)
	// }

	//infinite loop
	// counter :=1
	// for{
	// 	fmt.Printf("Number :%d\n",counter)
	//     counter++
	//     if counter == 10 {
	//         break
	//     }
	// }

	// range
	// numbers := []int{1, 2, 3, 4, 5}
	// for index, value := range numbers {
	// 	fmt.Printf("Index :%d, Value :%d\n", index, value)
	// }

	// name:=[]string{"John","John2", "John3", "John4", "John5"}
	// for index,value:=range name{
	// 	fmt.Printf("Index :%d, Value :%s\n", index, value)
	// }

	data:="Hello,World!"
	for index,char:=range data{
		fmt.Printf("Index :%d, Value :%c\n", index, char)
	}
}
