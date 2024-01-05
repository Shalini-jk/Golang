// write a program to find the largest among two number

package main 

import(
	"fmt"
)

func main() {
	//decalre the variable
	var first_number int
	var second_number int
	
	//taking input from user
	fmt.Println("Enter the first number")
	_, err := fmt.Scanln(&first_number)
	if err != nil {
		fmt.Println("input is not in not in correct format",err)
		return
	}

	fmt.Println("Enter the second number")
	_, err = fmt.Scanln(&second_number)
	if err != nil {
		fmt.Println("input is not in not in correct format",err)
		return
	}

	if first_number > second_number {
		fmt.Println("first number is greater than second number")
	} else {
		fmt.Println("second number is greater than first number")
	}
}