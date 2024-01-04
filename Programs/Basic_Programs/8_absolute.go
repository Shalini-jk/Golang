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
	_,err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("The number is not float or Integer",err)
		return
	}
	//calculate the absolute value using math function
	absolute_value := math.Abs(number)
	fmt.Println("The absolute value is",number,absolute_value)
}