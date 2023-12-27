/*
Even or Odd:
problem :
Write a program that prompts the user to enter an integer and determines whether the entered 
number is even or odd. Display an appropriate message based on the result.

Algorithm :
  1. Take a numer as input from user
  2. Now check whether it is  even or odd
  3. Then  print the message accrding to the result 
*/

package main 
import (
	"fmt"
)
func main() {
  // declare variable
  var input int
  
  // Take input from the user
  fmt.Println("Enter the numer which you can want to check")
  fmt.Scan(&input)

  // implement the Logic
  if (input > 0) {
    if ( input % 2 == 0) {
      fmt.Println("The Entered Number is even number")
    } else {
      fmt.Println("The Entered Number is odd number")
    } 
  } else {
    fmt.Println("The entered number is negative,Can not find whether the number is even or odd")
    
  }
}