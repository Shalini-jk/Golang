/*
problem :
Write a program that takes numeric input as a string, converts it to a numeric type using 
the strconv package, performs a mathematical operation, and then converts the result back 
to a string for display.

Algorithm :
  1. Take a numberic input as a string
  2. then convert that string into integer.
  3. then perorm the mathematical operations.
  4. then convert the result agan into string
  5. then print the input and output in the form of string. 
*/

package main

import (
  "fmt"
  "strconv"
)

func Takeinput() {
  //Declaration of variable
  var First_string string
  var Second_string string

  //Taking input from user
	fmt.Print("Enter the first numeric string: ")
	fmt.Scanln(&First_string)

  fmt.Print("Enter the second numeric string: ")
	fmt.Scanln(&Second_string)

  First_int, err := strconv.Atoi(First_string)
	if err == nil {
		// Display the integer
		fmt.Printf("Integer value: %d\n", First_int)
  }

  Second_int, err := strconv.Atoi(Second_string)
	if err == nil {
		// Display the integer
		fmt.Printf("Integer value: %d\n", Second_int)
  }

}

func Arthmeticoperations() {
  var result int 

  result = First_int - Second_int 
  fmt.Println("The sutraction of the two string :", result)

  result = First_int + Second_int
  fmt.Println("The Addition of two string :", result)
}
func main()  {
  Takeinput()
  Arthmeticoperations()
}