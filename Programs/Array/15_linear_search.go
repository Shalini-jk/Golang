// write a program to search an element in arra using linear search.
/*
Linear search : In linear search we search the array element sequentially. till the target 
element found. if the target element found then we return the element along with the index
otherwise we return msg with element not found.

Algorithm : 
    1. declare an array 
	2. take target element as input.
	3. now iterate the array from starting index to last index.
	4. while iterating element check that the iterating element is equal to target element or 
	not 
	5. then print the result if the condition matched

Handle required error
*/

package main 

import(
	"fmt"
)

func main() {
	//declare array
	arr1 := [] int{45,62,2,3,4,5,6,78,8,9,42,88}

	//declare variable
	var input int

	//taking input from user
	fmt.Println("Enter the number which you want to search :")
	_,err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Incorrect Input, input must be integer",err)
	}
	var get_value = false

	//now iterate the array element
	for i := 0 ; i < len(arr1) ; i++ {
		if input == arr1[i] {
			get_value = true
			//fmt.Println("element found",input,"at the index of",i)//we can also use println in this way
			break
		} 
	}
	if get_value == true {
		fmt.Println("element found",input)	
	} else {
		fmt.Println("element  not found")
	} 
}


