// write a program to calculate the sum odf all elements in array

package main

import(
	"fmt"
)

func main() {
	//declare array
	arr := [] int{2,3,4,5,6,7,8,9}
	var sum int = 0 
	
	get_size := len(arr)
	fmt.Println("size of array",get_size)

	//calculating sum of array elements
	for i := 0 ; i < get_size ; i++  {
		sum = sum + arr[i] 
	}
	fmt.Println("sum of elements of array",sum)
}