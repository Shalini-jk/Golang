// write a program to sort element by implementing buble sort

/*
Bubble sort Bubble Sort is a simple sorting algorithm that repeatedly steps through the list,
compares adjacent elements, and swaps them if they are in the wrong order

Algorithm :
   1.  declare array
   2.  compare adjacent element and then swap them accordingly
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
	fmt.Println("Sort array in ascending order",unsorted_arr)
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
	fmt.Println("Sort array in descending order",unsorted_arr)
}
func main() {
	Ascen_bubble_sort()
	Desc_bubble_sort()
}

/*
shalini@Ubuntu:~/Golang/Programs/Array$ go run 18_bubble_sort.go
Sort array in ascending order [1 2 4 5 6 8 9 24 33 37 66 99]
Sort array in descending order [99 66 37 33 24 9 8 6 5 4 2 1]
*/