// golang program to calculate power of number using the power function 

package main

import(
	"fmt"
	"math"
)

func main()  {
	//declare variable 
	var number float64
	var power_value float64

	// taking input from user
	fmt.Println("Enter the number")
	_,err := fmt.Scanln(&number)
	if err != nil {
		fmt.Println("The number is not Integer or float",err)
		return
	}

	fmt.Println("Enter the power")
	_,err = fmt.Scanln(&power_value)
	if err != nil {
		fmt.Println("The number is not Integer or Float")
		return
	}
	

    //calculating the power
	get_power := math.Pow(number,power_value)

	//printing the result
	fmt.Println("The result",get_power)
}