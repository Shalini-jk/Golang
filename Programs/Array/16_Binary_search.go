// write a program to search an element in arra using binary search

/*
Binary search : binary search is a searching algorithm. in binary search we can search the 
element  dividing it into two half and then compare it with the target element 

condition : to perform binary search algorithm the array must e sorted

Algorithm :
   1. First decalare the array
   2. Then check whether the array is sorted or not.if array is not sorted then sort the array.
   3. Then take target numer as input from the user
   4. Now assign the lowest index and highest index
   5. Then calculate the middle index (low+high)/2
   6. Then check whether the middle index value is equal to the target value or not
    if found then return the value.
   7. if does not match then again we set the low index and high index and calculate the middle index
   and then compare. follow this step until u find the target number. 
*/

package main

import(
	"fmt"
	"sort"
)

func main() {
	//declare variable
	var low_index int
	var highest_index int
	var middle_index int
	var input_value int

	//declare array
	arr1 :=[] int{2,8,9,4,5,3,7,6,10}// {2,3,4,5|,6,|7,8,9,10}
	// now sort the array
	sort.Ints(arr1)
	fmt.Println(arr1)
	// take input from the user
	fmt.Println("Enter the number which you want to search")
	fmt.Scanln(&input_value)
	//now assign the lowest index and highest index
	low_index = 0
	highest_index = len(arr1)-1
	fmt.Println("Lowest index",low_index)
	fmt.Println("Highest index",highest_index)
    get_found := false
	for low_index <= highest_index {
		// calculating the mid value
	    middle_index = (low_index + highest_index) / 2
	    fmt.Println("Middle index",middle_index)
		// now compare the value of middle index with the target value
		if arr1[middle_index] == input_value {
			get_found = true
			break
			//fmt.Println(middle_index)
		} else if arr1[middle_index] < input_value {
			// set new lowest index as the element which we searched is greater than middle element
			low_index = middle_index + 1 
		} else {
			// set new highest index as the element which we searched is smaller than the middle element
			highest_index = middle_index - 1
		}	
	}
	if get_found == true {
		fmt.Println("element found",input_value, "at index",middle_index)
	} else {
		fmt.Println("Element not found")
	}
}

/*
shalini@Ubuntu:~/Golang/Programs/Array$ go run 16_Binary_search.go
[2 3 4 5 6 7 8 9 10]
Enter the number which you want to search
8
Lowest index 0
Highest index 8
Middle index 4
Middle index 6
element found 8 at index 6
*/