package main

import "fmt"

func main() {
	// fmt.Println("We are learning about switch-case in Golang")

	// var day string
	// fmt.Println("Enter the day (e.g., Sunday or 1):")
	// fmt.Scan(&day)

	// switch day {
	// case "Sunday", "1":
	// 	fmt.Println("Today is a Sunday")
	// case "Monday", "2":
	// 	fmt.Println("Today is a Monday")
	// case "Tuesday", "3":
	// 	fmt.Println("Today is a Tuesday")
	// case "Wednesday", "4":
	// 	fmt.Println("Today is a Wednesday")
	// case "Thursday", "5":
	// 	fmt.Println("Today is a Thursday")
	// case "Friday", "6":
	// 	fmt.Println("Today is a Friday")
	// case "Saturday", "7":
	// 	fmt.Println("Today is a Saturday")
	// default:
	// 	fmt.Println("Invalid day")
	// }


	month:="january"
	switch month{
	case "january","february", "march":
		fmt.Println("Winter")
		break 
	case "april","may","june":
		fmt.Println("Spring")
        break
	case "july","august","september":
		fmt.Println("Summer")
        break
	case "october","november","december":
		fmt.Println("Autumn")
        break
	default:
		fmt.Println("Invalid month")
        break
	} 

}
