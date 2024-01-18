// write a go program to print the ASCII value of a character

package main

import(
	"fmt"
)

func main()  {
	var result string
	// taking input from user 
	fmt.Println("Enter the character")
	fmt.Scan(&result)
	for _, v := range result {
		fmt.Println(v)
	}
	
}