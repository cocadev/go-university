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
	// var BMI = 21.0
	// if BMI < 18.5 {
	// 	fmt.Println("You are underweight");
	// } else if BMI >= 18.5 && BMI < 25.0 {
	// 	fmt.Println("Your weight is normal");
	// } else if BMI >= 25.0 && BMI < 30.0 {
	// 	fmt.Println("You're overweight")
	// } else {
	// 	fmt.Println("You're obese")
	// }

	//====================== 4
	// var dayOfWeek = 6
	// switch dayOfWeek {
	// 	case 1: fmt.Println("Monday")
	// 	case 2: fmt.Println("Tuesday")
	// 	case 3: fmt.Println("Wednesday")
	// 	case 4: fmt.Println("Thursday")
	// 	case 5: fmt.Println("Friday")
	// 	case 6: {
	// 		fmt.Println("Saturday")
	// 		fmt.Println("Weekend. Yaay!")
	// 	}
	// 	case 7: {
	// 		fmt.Println("Sunday")
	// 		fmt.Println("Weekend. Yaay!")
	// 	}
	// 	default: fmt.Println("Invalid day")
	// }

	//======================= 5
	switch dayOfWeek := 5; dayOfWeek {
	case 1, 2, 3, 4, 5:
		fmt.Println("Weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid Day")		
	}
}
	
