// golang program to demonstrate switch case with optional statement and expression
/*
optional statement : 
expression : 
*/
/*
syntax :
 switch (optional statement); (expression) {
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
	//  day= 4 ; is optional statement
	// day is expression
	switch day:= 5 ; day {
	case 1:
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

