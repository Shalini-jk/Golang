// write a program to swap two numbers
/*
Algorithm :
   1. Take two number as input
   2. use the logic to assign the value of first number to second number and value of s
   econd number to the first number.
   3. then print the orginal value of two number as the swapped value of two number
*/
// program

package main

import (
	"fmt"
)

func main(){
	var first_number int
	var second_number int
	var temp_value int

	fmt.Println("Enter the value of First number :")
	fmt.Scan(&first_number)

	fmt.Println("Enter the value of Second number :")
	fmt.Scan(&second_number)
// swap the values with the help of temporary variable
fmt.Println("The original value of both number First number : %d, second number : %d ", first_number,second_number)
temp_value = first_number
first_number = second_number
second_number = temp_value

fmt.Println("The swapped value of both number First number : %d, second number : %d", first_number,second_number)

// swap the value without the help of third variable
first_number = first_number + second_number
second_number = first_number - second_number
first_number = first_number - second_number

fmt.Println("The swapped value of both number First number : %d, second number : %d", first_number,second_number)



}

