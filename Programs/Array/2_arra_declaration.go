// wriite a program to demonstrate the short hand declaration of array

package main

import(
	"fmt"
)

func main() {
	//declare array in short hand
	arr1 := [5] int{1,2,3,5,9}
	fmt.Println(arr1)

	// long declaration 
	var arr2[]int 
	fmt.Println(arr2)
}