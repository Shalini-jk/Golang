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
	_, err := fmt.Scanln(&first_number)
	if err != nil {
		fmt.Println("The entered number is not Integer",err)
		return
	}
	
	fmt.Println("Enter the second number")
	_, err = fmt.Scanln(&second_number)
	if err != nil {
		fmt.Println("The entered number is not Integer",err)
		return
	}

	
	fmt.Println("Enter the third number")
	_, err = fmt.Scanln(&third_number)
	if err != nil  {
		fmt.Println("The entered numer is not Integer",err)
		return
	}

    // printing the result
	average := (first_number + second_number + third_number) / 3
	fmt.Println("The average of three number:", average)
}