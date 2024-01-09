// write a program to create a new array from the existing array
package main

import(
	"fmt"
)

func main() {
	// declare a array
	arr1 := [] int{2,4,7,9,8,5,3,2,66,89,96}

	get_size := len(arr1)
	fmt.Println("size of previous array :",get_size)
	fmt.Println("Previous array :",arr1)

	// create another array by using the existing array
	var arr2 []int
	for i := 2 ; i < len(arr1)-3 ; i++ {
		arr2 = append(arr2, arr1[i])
	}
	fmt.Println("size o newl created array :",len(arr2))
	fmt.Println("newly created array from previous one :",arr2)
}