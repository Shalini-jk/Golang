//write a program to get even number in arrray

package main

import(
	"fmt"
)

func main() {
	//declare array
	arr := [] int{2,3,7,8,9,77,88,99,102}
	get_size := len(arr)
	//value := 0

	//finding even number
	for i := 0 ; i < get_size ; i++ {
		//fmt.Println(arr[i])
		//value = arr[i]
		if arr[i] % 2 == 0 {
			fmt.Println(arr[i])
		}

	}

}