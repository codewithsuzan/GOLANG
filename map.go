package main

import (
    "fmt"
)

func main() {
	// fmt.Println("we are learning about Maps in Go")
	// //maps in Go, objects in JavaScript, and dictionaries in Python serve similar purposes

	studentGrade:=make(map[string]int)
	studentGrade["Alice"]=95
	studentGrade["Bob"]=85
	// fmt.Println("Marks of Alice:",studentGrade["Alice"])
	// studentGrade["Alice"]=100
	// fmt.Println("New marks of Alice:",studentGrade["Alice"])
	// delete(studentGrade,"Alice")
	// fmt.Println("Marks of Alice after deletion:",studentGrade["Alice"]) //0

	// grades,exists:=studentGrade["David"]
	// fmt.Println("Marks of David:",grades)
	// fmt.Println("Does David exist in the map?",exists) 

	Grades,Exists:=studentGrade["Bob"]
	fmt.Println("Does Bob exist in the map?",Exists) 
	fmt.Println("Marks of Bob:",Grades)

	for index,value:=range studentGrade{
		fmt.Printf("Index :%s, Value :%d\n", index, value)
	}

	persons:=map[string]string{
		"Alice":"Female",
        "Bob":"Male",
        "Charlie":"Female",
	}
	fmt.Println("Persons:",persons)	
}