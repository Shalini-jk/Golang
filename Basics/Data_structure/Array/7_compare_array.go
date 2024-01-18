// write a program to compare two array

package main

import(
	"fmt"
)

func main() {
	//declare array 
	arr1 := [6] int{2,4,6,7,8,9}
	arr2 := [6] int{6,7,99,77,88,55}
	arr3 := [6] int{2,4,6,7,8,9}
	if arr1 == arr2 {
		fmt.Println("two array 1 and array 2 are equals")
	} else {
		fmt.Println("array 1 and array 2 are  not equals")
	}
	if arr1 == arr3 {
		fmt.Println("array 2 and array 3 is same")
	} else {
		fmt.Println("array 2 and array 3 is same")
		
	}
	
	
}
