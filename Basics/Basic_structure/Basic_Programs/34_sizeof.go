// write a program to demostrate size of operators.

package main 

import(
	"fmt"
	"unsafe"
)

func main() {
	//declaring variable
	var number int
	var value string

	// input the number
	fmt.Println("enter the number")
	fmt.Scanln(&number)

	fmt.Println("enter the string value")
	fmt.Scanln(&value)

	result := unsafe.Sizeof(number)
	fmt.Println("the size of entered number",result)

	result1 := unsafe.Sizeof(value)
	fmt.Println("the size of entered number",result1)
}