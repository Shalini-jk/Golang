// write a program to add two integer number

package main

import (
	"fmt"
)

func main() {
	//decalre the variable
	var first_number int
	var second_number int

	//taking input from user
	fmt.Println("Enter the first number")
	_, err := fmt.Scanln(&first_number)
	if err != nil {
		fmt.Println("The Entered value is not an integer",err)
		return
	} 

	fmt.Println("Enter the second number")
	_, err = fmt.Scanln(&second_number)
	if err != nil {
		fmt.Println("The Entered value is not an integer",err)
		return
	}

	//printing the result
	result := first_number + second_number
	fmt.Println("the addition of two integer value",result)
}