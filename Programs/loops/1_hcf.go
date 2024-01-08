// write a program to calculate  HCF of two number

package main

import(
	"fmt"
)

func main() {
	//declare variable 
	var num1 int
	var num2 int
	var i, j int

	//taking input from user
	fmt.Println("Enter the first number")
	fmt.Scanln(&num1)

	for i = 1 ; i <= num1 ; i++ {
		if num1 % i == 0 {
			fmt.Println("Factor of first number",i)
		}	
	}
	fmt.Println("Enter the Second number")
	fmt.Scanln(&num2)

	for j = 1 ; j <= num2 ; j++ {
		if num2 % j == 0 {
			fmt.Println("Factor of second number",j)
		}
		max_factor := 0
	    if i == j  {
			if max_factor < i {
				max_factor = i
			}
			fmt.Println("HCF",max_factor)
	}
		
	}
	

	
}