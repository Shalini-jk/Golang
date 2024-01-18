// write a program to reverse the integer/number inside the array
/*
Example : 
input :  arr1{1,2,3,4,5}  output:  arr= {5,4,3,2,1}
Alogorithm : 
   1. Declare array 
   2. then iterate that array in decending order means from last index to the first index
   3. then stored the value during iteration and then print that value
*/

package main

import(
	"fmt"
)

func main()  {
	//declare the variable
	var reverse_arr[]int
	arr1 := [] int{2,3,4,5,6,7,8,9}

	//now iterate the array in reversing order
	for i:= len(arr1)- 1 ; i >=0 ; i-- {
		reverse_arr = append(reverse_arr,arr1[i])	
	}
	fmt.Println("original array",arr1)
	fmt.Println("Reversed array",reverse_arr)
	
}