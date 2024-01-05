// write a program to find the largest among three number

package main 

import(
	"fmt"
)

func main() {
	//decalre the variable
	var first_number int
	var second_number int
	var third_number int
	
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

	fmt.Println("Enter the third number")
	_, err = fmt.Scanln(&third_number)
	if err != nil {
		fmt.Println("input is not in not in correct format",err)
		return
	}

	if first_number > second_number &&  first_number > third_number {
		fmt.Println("first number is greatest among all")
	} else if  second_number > first_number &&  second_number > third_number {
		fmt.Println("second number is greatest among all")
	} else {
		fmt.Println("third number is greatest among all")
	}
}