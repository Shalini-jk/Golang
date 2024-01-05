// write a program to check whether the numer is positive or negative

package main 

import(
	"fmt"
)

func main() {
	//declare the variable
	var number int

	//taking input from user
	fmt.Println("Enter the number")
	_,err := fmt.Scanln(&number)
	if err != nil {
		fmt.Println("The input is not in the correct format",err)
		return
	}
	if number >= 0 {
		fmt.Println("Positive number")
	} else {
		fmt.Println("negative number")
	}
}