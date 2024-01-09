// write a program to get the size of array 

package main

import(
	"fmt"
	"unsafe"
)
func main() {
	//declare array
	arr1 := []int{1,2,3,4,5,7,8,96,84,98,57}
	result :=  unsafe.Sizeof(arr1)// it gives size of array in bytes
	result1 := len(arr1)
	fmt.Println("size getting unsafe package ",result)
	fmt.Println(result1) // it gives the size of array
}