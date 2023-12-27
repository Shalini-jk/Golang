/*
Number Palindrome:
problem :
Create a program that checks whether a given integer is a palindrome. A palindrome is a number
that reads the same backward as forward. For example, 121 is a palindrome.

Alogrithm : 
  1. Take a number as input
  2. Then reverse te number.
  3. Then stored the original number and reversed number in seperate variable
  4. then check whether the original numer is equal to reversed number 
  5. print the msg according to the requirement
*/
package main 
import (
	"fmt"
)
func main() {
	// decalare variable
	var original_number int
	var reverse_number = 0
	var temp int

	// Taking input from user
	fmt.Println("Enter the number to check palindrom or not")
	fmt.Scan(&original_number)
	temp = original_number

	// reverse the number
	for temp > 0 {
		digit := temp % 10
		reverse_number = reverse_number * 10 + digit
		temp = temp / 10
	}
	fmt. Println("The Reversed Number is:", reverse_number)
	
	if (original_number == reverse_number) {
		fmt.Println("The Numer is Palindrom Number")
	}  else {
		fmt.Println("The numer is not a palindrom number")
	}
}