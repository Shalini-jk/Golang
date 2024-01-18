// golang program to use the expression in case statement
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

	var day int = 2
	// here it first execute the switch case ie 7-2 and then the result comes which is 
	//5 is then  compared with the initailzed value of the variable which we used as a expression
	// and then if the compaarision is true then it gives the result.

	switch day {
	case 7- 1 :
		fmt.Println("Sunday")	
	case 7 - 2:
		fmt.Println("monday")
	case 7 -3:
		fmt.Println("Tuesday")
	case 7 - 4:
		fmt.Println("wednesday")
	case 7 - 5:
		fmt.Println("Thursday")
	case 7 - 6:
		fmt.Println("Friday")
	case 7 - 7:
		fmt.Println("Saturday")
	default :
	     fmt.Println("Invalid input")



	}
}
