//write a program to find the largest element in arrray

package main

import(
	"fmt"
)

func main()  {
	// declare array
	arr := [] int{2,3,7,8,9,77,88,99,102,100}
	get_size := len(arr)
	var largest_number int


	//finding largest number from array
	for i := 0 ; i < get_size ; i++ {
		//fmt.Println(arr[i])
		if largest_number < arr[i] {
			largest_number = arr[i]
		}
    }
	fmt.Println("largest number in array",largest_number)
}