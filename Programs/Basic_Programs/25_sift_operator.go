package main

import(
	"fmt"
)

func main() {
	//declare variable
	var first_number int
	var result_value int

	//taking input from user
	fmt.Println("Enter the number")
	fmt.Scanln(&first_number)

	result_value = first_number >> 3

	fmt.Println("the reslut  using right sift operator",result_value)

}