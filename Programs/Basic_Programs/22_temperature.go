/*
Problem:
Temperature Conversion:

Write a program that converts a temperature in Celsius to Fahrenheit. Prompt the user to enter 
a temperature in Celsius, perform the conversion, and print the result in Fahrenheit.
 
Formula :  celsius to Fahrenheit;  F=(9/5)​C+32

Algorithm : 
      1. First take input for temperature in celsius
	  2. Then we use formula for the conversion of temperature from celsius to fahrenheit
	  3. then print both the temperature on the terminal 
*/

package main
import (
	"fmt"
)
func main() {
	// declaration of variable
	var celsius float64
	var fahrenheit float64

	// taking input from the user
	fmt.Println("Enter the temperature in Celsius")
	fmt.Scan(&celsius)

	// Conversion of temperature from celsius to fahrenheit
	fahrenheit = (9.0/5.0) * celsius + 32

	// print the temperature in celsius and fahrenhiet
	fmt.Printf("The Temperature in Celsius: %.2f °C\n", celsius)
    fmt.Printf("The Temperature in Fahrenheit: %.2f °F\n", fahrenheit)
	

}