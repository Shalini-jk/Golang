// write a program to swap two integer value

package main

import (
	"fmt"
)

func main()  {
	//decalare the variable
	var first_number int
	var second_number int
	var temp int
	
	// taking input from user
	fmt.Println("Enter the first number")
	_, err := fmt.Scanln(&first_number)
	if err != nil {
		fmt.Println("The Entered number is not INteger",err)
		return
	}

	fmt.Println("Enter the second number")
	_,err = fmt.Scan(&second_number)
	if err != nil {
		fmt.Println("The Entered number is not Integer",err)
		return
	}

	fmt.Println("the number before swapping",first_number,second_number)

	//swap two number by using third number
	temp = first_number
	first_number = second_number
	second_number = temp

	//print the result
	fmt.Println("the number after swapping",first_number,second_number)

}