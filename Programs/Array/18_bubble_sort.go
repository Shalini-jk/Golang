// write a program to sort element by implementing buble sort

/*
Bubble sort Bubble Sort is a simple sorting algorithm that repeatedly steps through the list,
compares adjacent elements, and swaps them if they are in the wrong order

Algorithm :
   1.  declare array
   2.  compare adacent element and then swap them accordingly
*/

package main 

import(
	"fmt"
)
// we implement this in assecending order
func Ascen_bubble_sort() {
	//declare variable
	var temp_value int
	// declare array
	unsorted_arr := []int{8,2,6,33,5,9,99,37,66,24,1,4}
	//in first  pass we compare 8 and 2 since it is not in correct way so it swapped
    // first loop iterate it from the zero index
	for start_index := 0; start_index < len(unsorted_arr)-1; start_index++ {
		for end_index := 0 ; end_index < len(unsorted_arr) - start_index -1; end_index++ {
			// Compare adjacent elements and swap if needed
			if unsorted_arr[end_index] > unsorted_arr[end_index+1] {
				temp_value = unsorted_arr[end_index]
				unsorted_arr[end_index] = unsorted_arr[end_index+1]
				unsorted_arr[end_index+1] = temp_value
			}
		}
	}
	fmt.Println(unsorted_arr)
}
// sort elements in descending order
func Desc_bubble_sort() {
	//declare variable
	var temp_value int
	// declare array
	unsorted_arr := []int{8,2,6,33,5,9,99,37,66,24,1,4}
	//in first pass we compare 8 and 2 since it is in correct way so it remain as it is
    // first loop iterate it from the zero index
	for start_index := 0; start_index < len(unsorted_arr)-1; start_index++ {
		for end_index := 0 ; end_index < len(unsorted_arr) - start_index -1; end_index++ {
			// Compare adjacent elements and swap if needed
			if unsorted_arr[end_index] <  unsorted_arr[end_index+1] {
				temp_value = unsorted_arr[end_index]
				unsorted_arr[end_index] = unsorted_arr[end_index+1]
				unsorted_arr[end_index+1] = temp_value
			}
		}
	}
	fmt.Println(unsorted_arr)
}
func main() {
	Ascen_bubble_sort()
	Desc_bubble_sort()
}
