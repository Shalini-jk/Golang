// write a program to merge two array in the third array
/*
Example : a = {1,2,3} , b = {4,5,6} then c = {1,23,4,5,6}

Algorithm : 
   1. Declare two array.
   2. then use two arrray to determine the size of third/new array.
   3. Then assign the value of each array to the new array one by one
*/

package main

import(
	"fmt"
)

func main() {
	//declare array
	arr1 := [] int{2,3,4}
	arr2 := [] int{5,6,7,8,9}
	var arr3[8]int //

	// geting the size of both array
	get_size1 := len(arr1)
	get_size2 := len(arr2)
	new_size :=  get_size1 + get_size2
	fmt.Println(new_size) 
	//merge the two array
	//copy the element of one array to new array
	copy(arr3[:],arr1[:])
	//now during copy first search at the last index of arr3 after coping arr1 and then 
	//copy arr2
	copy(arr3[len(arr1):], arr2[:])
	fmt.Println(arr3)
}