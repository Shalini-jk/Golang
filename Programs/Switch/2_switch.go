// golang program to demonstrate switch case without optional statement and expression
/*
optional statement : 
expression : 
*/
/*
syntax :
 switch  {
	case variale == condition :
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
	switch {
	case day == 1 :
		fmt.Println("Sunday")	
	case day == 2:
		fmt.Println("monday")
	case day == 3:
		fmt.Println("Tuesday")
	case day == 4:
		fmt.Println("wednesday")
	case day == 5:
		fmt.Println("Thursday")
	case day == 6:
		fmt.Println("Friday")
	case day == 7:
		fmt.Println("Saturday")
	default :
	     fmt.Println("Invalid input")



	}
}

