// write a program to find the average of three number
// average = a+b+c/3

package main 

import (
	"fmt"
)

func main()  {
	// declare the variable
	var first_number int
	var second_number int
	var third_number int

	//taking input from user
	fmt.Println("Enter the first numebr")
	fmt.Scan(&first_number)

	fmt.Println("Enter the second number")
	fmt.Scan(&second_number)

	fmt.Println("Enter the third number")
	fmt.Scan(&third_number)

        // printing the result
	average := (first_number + second_number + third_number) /3
	fmt.Println("The average of three number:", average)
}
