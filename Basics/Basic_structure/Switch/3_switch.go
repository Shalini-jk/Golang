// golang program to demonstrate switch case with expression
/*
optional statement : 
expression : 
*/
/*
syntax :
 switch (expression) {
	case1 :
	statement
	......
	default :
	statement
 }
*/
package main

import(
	"fmt"
)
func main () {

	var day int
	fmt.Println("Enter the switch case etween 1 to 7")
	fmt.Scan(&day)
	switch day {
	case 1 :
		fmt.Println("Sunday")	
	case 2:
		fmt.Println("monday")
	case 3:
		fmt.Println("Tuesday")
	case 4:
		fmt.Println("wednesday")
	case 5:
		fmt.Println("Thursday")
	case 6:
		fmt.Println("Friday")
	case 7:
		fmt.Println("Saturday")
	default :
	     fmt.Println("Invalid input")



	}
}

