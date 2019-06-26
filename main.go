package main

import "fmt"

func main() {


	//================================ 1

	//fmt.Println("Hello, World")

	//================================ 2

	// Declaring Variables
	// var myStr string = "Hello"
	// var myInt int = 100
	// var myFloat float64 = 45.12
	//fmt.Println(myStr, myInt, myFloat)


	// Multiple Declarations
	// var (
	// 	employeeId int = 5
	// 	firstName, lastName string = "Satoshi", "Nakamoto"
	// )
	//fmt.Println(employeeId, firstName, lastName)

	
	// Short variable declaration syntax
	// name := "Rajeev Singh"
	// age, salary, isProgrammer := 35, 50000.0, true

	//fmt.Println(name, age, salary, isProgrammer)


	//================================ 3
	// var name = "Rajeev Singh" // Type declaration is optional here.
	// fmt.Printf("Variable 'name' is of type %T\n", name)

	// //================================

	// // Multiple variable declarations with inferred types
	// var firstName, lastName, age, salary = "John", "Maxwell", 28, 50000.0

	// fmt.Printf("firstName: %T, lastName: %T, age: %T, salary: %T\n",
	// 	firstName, lastName, age, salary)

	//================================= 4

	var (
		firstName, lastName string
		age                 int
		salary              float64
		isConfirmed         bool
	)

	fmt.Printf("firstName: %s, lastName: %s, age: %d, salary: %f, isConfirmed: %t\n",
		firstName, lastName, age, salary, isConfirmed)

}