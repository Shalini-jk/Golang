package operation

import (
	"fmt"
)

//unc Addition() //{
	//declare variable
	//var First_number float64
	//var Second_number float64
	//var Result float64

	// Take input
	/*fmt.Println("Enter the First Number")
	fmt.Scan(&First_number)

	fmt.Println("Enter the second Number")
	fmt.Scan(&Second_number)

	Result = First_number + Second_number
	
	fmt.Println("The Addition of two number :",Result)
}*/
func Add(a, b int) int {
    return a + b
}

func Subtraction() {
	//declare variable
	var First_number float64
	var Second_number float64
	var Result float64

	// Take input
	fmt.Println("Enter the First Number")
	fmt.Scan(&First_number)

	fmt.Println("Enter the second Number")
	fmt.Scan(&Second_number)

	Result = First_number - Second_number
	
	fmt.Println("The Sutraction of two number :",Result)
}

func Multiplication() {
	//declare variable
	var First_number float64
	var Second_number float64
	var Result float64

	// Take input
	fmt.Println("Enter the First Number")
	fmt.Scan(&First_number)

	fmt.Println("Enter the second Number")
	fmt.Scan(&Second_number)

	Result = First_number * Second_number
	
	fmt.Println("The Multiplication of two number :",Result)
}

func Division() {
	//declare variable
	var First_number float64
	var Second_number float64
	var Result float64

	// Take input
	fmt.Println("Enter the First Number")
	fmt.Scan(&First_number)

	fmt.Println("Enter the second Number")
	fmt.Scan(&Second_number)

	Result = First_number / Second_number
	
	fmt.Println("The Division of two number :",Result)
}