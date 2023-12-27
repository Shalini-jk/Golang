/*
BMI Calculator:
problem : 
Write a program that calculates the Body Mass Index (BMI) based on the user's weight in 
kilograms and height in meters. Prompt the user to enter their weight and height, perform the 
calculation, and display the BMI.

formula : BMI = weight(kg)/height(m) ** 2

Algorithm : 
  1. Take height and weight as input from the user
  2. Calculate the bmi b using formula.
  3. print the calculated BMI
*/
package main
import (
	"fmt"
)
func main() {
	//declare variable
	var height float64
	var weight float64
	var bmi float64

	// Take input from user
	fmt.Println("Enter the Height in ft:")
	fmt.Scan(&height)
	fmt.Println("Enter the weight in kg:")
	fmt.Scan(&weight)

	// calculations of BMI
	bmi = weight / (height * height)

	//print the BMI
	fmt.Printf("BMI: %.2f\n", bmi)

}