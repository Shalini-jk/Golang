// golang program to find smallest number among the number  using math function

package main 

import(
	"fmt"
	"math"
)

func main() {
	// declaring variable
	var first_number float64
	var second_number float64
	

	//taking input from the user
	fmt.Println("Enter the first number")
	fmt.Scanln(&first_number)

	fmt.Println("Enter the second number")
	fmt.Scanln(&second_number)

	smallest_number := math.Min(first_number,second_number)
	fmt.Println("The smallest number:", smallest_number)


}