// write a program using math.floor function
//it gives largest integer after roundup

package main

import(
	"fmt"
	"math"
)

func main() {
	//declaring variable
	var input float64

	// taking input from user
	fmt.Println("Enter the number")
	fmt.Scan(&input)

	result := math.Floor(input)
	fmt.Println("the floor of the number", result)
}