// write a golang program to find the remainder for floating point number using mod functions

package main

import(
	"fmt"
	"math"
)

func main() {
	//declare the variable
	var first_number float64
	var second_number float64

	// taking the input
	fmt.Println("enter the first number")
	fmt.Scanln(&first_number)

	fmt.Println("enter the second number")
	fmt.Scanln(&second_number)

	remainder_value := math.Mod(first_number,second_number)
	fmt.Println("the remaiinder of a floating point number", remainder_value)

}