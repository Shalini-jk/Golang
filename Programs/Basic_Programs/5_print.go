// write a golang program to read and print integer value

package main

import(
	"fmt"
)

func main()  {
	//declare the variable
	var number int 

	//taking input from user
	fmt.Println("Enter the number")
	fmt.Scanf("%d",&number)
	
	fmt.Println("The Integer input",number)
	

}