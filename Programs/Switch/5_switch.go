// golang program to  create a simple calculator using switch case.

package main

import(
	"fmt"
)

func main() {
	var firs_number int
	var second_number int
	var choice int

	fmt.Println("Enter the first number")
	fmt.Scan(&firs_number) 

	fmt.Println("Enter the second number")
	fmt.Scan(&second_number) 

	//using switch statement without using optional statement 
	fmt.Println("Enter your choice between 1 to 4")
	fmt.Scan(&choice) 

	switch choice {
	case 1:
		add :=  firs_number + second_number
		fmt.Println("addition of two number :",add)	

	case 2:
		sub := firs_number - second_number
		fmt.Println("subtraction of two number:",sub)

	case 3:
		mult := firs_number * second_number
		fmt.Println("Multiplication of two number :",mult)

	case 4:
		div := firs_number / second_number
		fmt.Println("division of two number :",div)
	default :
	    fmt.Println("Invalid Choice")
	}
}