// write a program to  check whether a number is even or odd

package main

import(
	"fmt"
)

func main() {
	// declare the variable
	var first_number int

	// taking input from the user
	fmt.Println("Enter the number")
	_, err := fmt.Scanln(&first_number)
	if err != nil {
		fmt.Println("the input is not in correct format",err)
		return
	}
	if first_number % 2 == 0 {
		fmt.Println("The number is even number")
	} else {
		fmt.Println("The number is odd number")
	}

}