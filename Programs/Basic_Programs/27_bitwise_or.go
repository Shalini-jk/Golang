package main

import(
	"fmt"
)

func main() {
	//declare variable
	var first_number int
	var second_number int
	var result_value int

	//taking input from user
	fmt.Println("Enter the number")
	fmt.Scanln(&first_number)

	fmt.Println("Enter the number")
	fmt.Scanln(&second_number)

	result_value = first_number|second_number

	fmt.Println("the reslut using or sift operator",result_value)

}