package main
import "fmt"

func main() {
	// // If Statement
	// var x = 25
	// if(x % 5 == 0) {
	// 	fmt.Printf("%d is a multiple of 5\n", x)
	// }

	// // Parentheses are Optional
	// var y = -1
	// if y < 0 {
	// 	fmt.Printf("%d is negative\n", y)
	// }
	
	// // If with a condition consisting of short circuit operators
	// var age = 21
	// if age >= 17 && age <= 30 {
	// 	fmt.Println("My Age is between 17 and 30")
	// }	

	// // If with a short statement	
	// if n := 10; n%2 == 0 {
	// 	fmt.Printf("%d is even\n", n)
	// } 

	//=====================2
	// var age = 18
	// if age >= 18 {
	// 	fmt.Println("You're eligible to vote!")
	// } else {
	// 	fmt.Println("You're not eligible to vote!")
	// }

	//======================3
	var BMI = 21.0
	if BMI < 18.5 {
		fmt.Println("You are underweight");
	} else if BMI >= 18.5 && BMI < 25.0 {
		fmt.Println("Your weight is normal");
	} else if BMI >= 25.0 && BMI < 30.0 {
		fmt.Println("You're overweight")
	} else {
		fmt.Println("You're obese")
	}
}
	
