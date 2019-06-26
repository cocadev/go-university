package main

import (
	 "fmt"
	//  "math"
	)

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

	// var (
	// 	firstName, lastName string
	// 	age                 int
	// 	salary              float64
	// 	isConfirmed         bool
	// )

	// fmt.Printf("firstName: %s, lastName: %s, age: %d, salary: %f, isConfirmed: %t\n",
	// 	firstName, lastName, age, salary, isConfirmed)

	//================================== 5
	// var myInt8 int8 = 97

	// /*
	//   When you don't declare any type explicitly, the type inferred is `int`
	//   (The default type for integers)
	// */
	// var myInt = 1200

	// var myUint uint = 500

	// var myHexNumber = 0xFF  // Use prefix '0x' or '0X' for declaring hexadecimal numbers
	// var myOctalNumber = 034 // Use prefix '0' for declaring octal numbers

	// var myFloat32 float32 = 4.5
	// var myFloat = 9.12 // // Type inferred as `float64` (the default type for floating-point numbers)

	// fmt.Printf("%d, %d, %d, %#x, %#o %f %f\n", myInt8, myInt, myUint, myHexNumber, myOctalNumber, myFloat32, myFloat)

	//========= 6
	// var myByte byte = 'a'
	// var myRune rune = '♥'

	// fmt.Printf("%c = %d and %c = %U\n", myByte, myByte, myRune, myRune)

	//========= 7
	// var a, b = 4, 5
	// var res1 = (a + b) * (a + b)/2  // Arithmetic operations

	// a++ // Increment a by 1

	// b += 10 // Increment b by 10

	// var res2 = a ^ b // Bitwise XOR

	// var r = 3.5
	// var res3 = math.Pi * r * r  // Operations on floating-point type

	// fmt.Printf("res1 : %v, res2 : %v, res3 : %v\n", res1, res2, res3)

	//========== 8
	// var myBoolean bool = true
	// var anotherBoolean = false // Inferred type

	// var truth = 3 <= 5
	// var falsehood = 10 != 10

	// // Short Circuiting
	// var res1 = 10 > 20 && 5 == 5     // Second operand is not evaluated since first evaluates to false
	// var res2 = 2*2 == 4 || 10%3 == 0 // Second operand is not evaluated since first evaluates to true

	// fmt.Println(myBoolean, anotherBoolean, truth, falsehood, res1, res2)

	//=========== 9
	// var x complex64 = 3.4 + 2.9i
	// var y = 5 + 7i // Type inferred as `complex128` (default type for complex numbers)

	// fmt.Println(x, y)

	// // Creating complex no from variables
	// var a1 = 4.5
	// var a2 = 7.1

	// var c = complex(a1, a2) // a1 + a2*i won't work
	// fmt.Println(c)

	// // ===== Complex No Operations =====
	// var a = 3 + 5i
	// var b = 2 + 4i

	// var res1 = a + b
	// var res2 = a - b
	// var res3 = a * b
	// var res4 = a / b

	// fmt.Println(res1, res2, res3, res4)

	//=============== 10
	var website = "\thttps://www.callicoder.com\t\n"

	// Raw String (Can span multiple lines. Escape characters are not interpreted)
	var siteDescription = `\t\tCalliCoder is a programming blog where you can find
                           practical guides and tutorials on programming languages, 
                           web development, and desktop app development.\t\n`

	fmt.Println(website, siteDescription)

}