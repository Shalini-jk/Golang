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
	fmt.Scan(&first_number)

	fmt.Println("Enter the second number")
	fmt.Scan(&second_number)

	result := first_number + second_number
	fmt.Println("the addition of two integer value",result)
}