// write a program to find the occurance of an item in array

package main

import(
	"fmt"
)

func main()  {
	//declare array
	arr1 := [] int{2,3,4,2,3,4,5,6,4,8,4,8,4} 
	// 2 = 2 , 3 = 2 , 4 = 5 , 8 = 2
	var count int
	var store_value int
	var temp int
	fmt.Println("Enter the value whose occurance you want")
	fmt.Scanln(&store_value)

	//iterate the array
	for i := 0 ; i < len(arr1) ; i++{
		if store_value == arr1[i] {
			count++
		}	
	}
	fmt.Println("number of occurance :",count)

	//now i want to calculate occurance of each element

}