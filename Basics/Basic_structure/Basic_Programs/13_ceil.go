// write a program using math.ceil function
// it return the smallest integer after roundup

package main

import(
	"fmt"
	"math"
)

func main() {
	//declare the variable
	var input float64

	//taking input from user
	fmt.Println("Enter the number")
	fmt.Scanln(&input)

	result := math.Ceil(input)
	fmt.Println("the ceil of this number",result)

}