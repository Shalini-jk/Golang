// write a program to create a new array from the existing array
package main

import(
	"fmt"
)

func main() {
	// declare a array
	arr1 := [] int{2,4,7,9,8,5,3,2,66,89,96}

	get_size := len(arr1)
	fmt.Println(get_size)
	fmt.Println(arr1)

	// create another array by using the existing array
	var arr2 []int
	arr2 =arr1
	fmt.Println("new array :",arr2)
}