// write a program  to find second largest element in array

/*
Algorithm : 
    1. define an array.
	2. Then iterate the each elements of array 
	3. while iterating array store the largest element in a variable.
	4. then in new  variale also store the value of largest element.
	5. then compared both the element and if the second variale is smaller than the
	 first variale  we get the reult. 
*/

package main

import(
	"fmt"
)

func main() {
	//declare variable
	var largest_number int
	var second_largest_number int
	//var value_no int
	//declare an array 
	arr1 := [] int{1,3,4,5,6,7,8,9,99,88,102,101}

	// then iterate each element of array 
	get_size := len(arr1)

	for i := 0 ; i < get_size ; i++ {
		//get the largest number
		if largest_number < arr1[i] {
			// store the largest number in another variable also
			second_largest_number = largest_number
			//fmt.Println(second_largest_number)
			// put the argest numer in one variale
			largest_number = arr1[i]
			// then compare oth variable	
		} else if second_largest_number < arr1[i] {
			second_largest_number = arr1[i]
		}			
	}
	fmt.Println("Largest",largest_number)
	fmt.Println("second Largest number",second_largest_number)				
}
