//write a program to sort the element using Insertion sort

/*
Insertion sort : In insertion sort we take one element except the 0th index and then compare 
it with the left element and then swap accordingly.

Algorithm :
    1. declare the array
    2. Start with the second element (index 1) and compare it with the first element (index 0).
    3. If the second element is smaller, swap them. Otherwise, leave them as they are.
    4.Move to the third element (index 2) and compare it with the elements to its left
	(ie index 1) until you find the correct position for insertion.
    5. Repeat this process for the remaining elements in the list, always maintaining the 
	sorted portion of the list.
*/

package main

import(
	"fmt"
)

func Insertion_sort_integer_ascen() {
	//declare variale
	var store_value int
	var compare_value int
	//declare the array
	int_arr := [] int{12, 11, 13, 5, 6}
	//iterate the array element starting from (the index 1)

	for start_index := 1 ; start_index < len(int_arr) ; start_index++ {
		store_value = int_arr[start_index]// (here we can store the element from index 1)
		//fmt.Println(store_value)
		//now we have to compare  the stored value with its left inde value
		//say 1th inde is compared to 0th inde and so on
		compare_value =  start_index - 1 // here we get the left index value
		// now compare the left index with the iterated index
		for compare_value >= 0 && int_arr[compare_value] > store_value {
			int_arr[compare_value+1] = int_arr[compare_value]
			compare_value = compare_value - 1
		}
		int_arr[compare_value+1] = store_value
		fmt.Println(int_arr)
	}
	fmt.Println("Sort array in ascending order:",int_arr)
	
}

func Insertion_sort_integer_desc() {
	//declare variale
	var store_value int
	var compare_value int
	//declare the array
	int_arr := [] int{12, 11, 13, 5, 6}
	//iterate the array element starting from (the index 1)

	for start_index := 1 ; start_index < len(int_arr) ; start_index++ {
		store_value = int_arr[start_index]// (here we can store the element from index 1)
		//fmt.Println(store_value)
		//now we have to compare  the stored value with its left inde value
		//say 1th inde is compared to 0th index and so on
		compare_value =  start_index - 1 // here we get the left index value
		// now compare the left index with the iterated index
		for compare_value >= 0 && int_arr[compare_value] < store_value {
			int_arr[compare_value+1] = int_arr[compare_value]
			compare_value = compare_value - 1
		}
		int_arr[compare_value+1] = store_value
		fmt.Println(int_arr)
	}
	fmt.Println("Sort array in descending order:",int_arr)
	
}
func main() {
	Insertion_sort_integer_desc()
	Insertion_sort_integer_ascen()
	
}

/*
go run 20_insertion_sort.go
[12 11 13 5 6]
[13 12 11 5 6]
[13 12 11 5 6]
[13 12 11 6 5]
Sort array in descending order: [13 12 11 6 5]
[11 12 13 5 6]
[11 12 13 5 6]
[5 11 12 13 6]
[5 6 11 12 13]
Sort array in ascending order: [5 6 11 12 13]
*/