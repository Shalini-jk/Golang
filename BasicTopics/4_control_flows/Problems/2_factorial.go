/* Factorial Calculation:
prolem :
Write a program that calculates the factorial of a given non-negative integer. The factorial 
of a number is the product of all positive integers up to that number. For example, the 
factorial of 5 is 5 * 4 * 3 * 2 * 1.
*/

package main

import "fmt"

func main() {
	var input int
	var factorial int
	
	// Taking input from user
	fmt.Print("Enter a non-negative integer: ")
	fmt.Scan(&input)

	// Check if the input is negative or positive
	if input > 0 {
		factorial = 1
	    for i := 1; i <= input; i++ {
			factorial *= i	
	    }
    } else {
		fmt.Println("Factorial is not defined for negative numbers.")
	}

	// Print the result
	fmt.Printf("Factorial of %d is: %d\n", input, factorial)
}
