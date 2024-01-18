// write a program  to get the absolute value of an integer number

package main 

import (
	"fmt"
	"math"
)

func main()  {
	//declare the variable
	var number float64

	//taking input from user
	fmt.Println("Enter the number whose asolute value is needed")
	_, err := fmt.Scanln(&number)
	if err != nil {
		fmt.Println("The input is not in correct format",err)
		return
	}
	if number > 0 {
		fmt.Println("The input is greater than 0 ")
		return
	}
	//calculate the absolute value using math function
	absolute_value := math.Abs(number)
	fmt.Println("The absolute value is",number,absolute_value)
}