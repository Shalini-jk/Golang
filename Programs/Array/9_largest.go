//write a program to find the largest element in arrray

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

	//finding largest number from array
	for i := 0 ; i < get_size ; i++ {
		//fmt.Println(arr[i])
		store_value = arr[i]
		//fmt.Println("store value",store_value)
	}
	if largest_number < store_value {
		largest_number = store_value
	}
	fmt.Println("largest number in array",largest_number)

}