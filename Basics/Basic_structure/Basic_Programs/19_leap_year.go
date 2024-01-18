// write a program to find the leap year(leap ear is when the total no of da is 366)

package main

import(
	"fmt"
)

func main() {
	//declare the variable
	var year_value int

	//taking input from user
	fmt.Println("Enter the year")
	_, err :=fmt.Scanln(&year_value)
	if err != nil {
		fmt.Println("the input is not in a correct format",err)
		return
	}
	if year_value % 4 == 0 {
		fmt.Println("the year is a leap year")
	} else {
		fmt.Println("the year is not a leap year")
	}
}