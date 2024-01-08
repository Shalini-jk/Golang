// write a program to calculate second largest element in array

package main

import(
	"fmt"
)

func main()  {
	// declare array
	arr := [] int{2,3,7,8,9,77,88,99,102}
	get_size := len(arr)
	var store_value int 
	store_value = 0 
	var largest_number int
	largest_number = 0
	var second_largest_number int
	second_largest_number = 0

	//finding largest number from array
	for i := 0 ; i < get_size ; i++ {
		store_value = arr[i]
	}
	if largest_number < store_value {
		largest_number = store_value
		second_largest_number = store_value
	}
	fmt.Println("largest number in array",largest_number)

	if second_largest_number <= largest_number {
		fmt.Println("second largest number in array",second_largest_number)
	}

}