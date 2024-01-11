// write a program to sort element by implementing buble sort
// in uble sort we compare the adjacent element and then perform swapping
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
func Ascen_bubble_sort_integer() {
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
			//new_value := len(unsorted_arr) - start_index -1
			//fmt.Println(new_value)
			fmt.Println(unsorted_arr)
		}
	}
	fmt.Println("Sort array in ascending order",unsorted_arr)
}
// sort elements in descending order
func Desc_bubble_sort_integer() {
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
			fmt.Println(unsorted_arr)
		}
	}
	fmt.Println("Sort array in descending order",unsorted_arr)
}

func Bubble_sort_float () {
	//declare variable
	var temp_value float64
	// declare array
	unsorted_arr := []float64{6.9, 7.2, 8.6, 9.1, 10.09, 2.4, 2.7, 4.5, 5.8, 3.6}
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
			//new_value := len(unsorted_arr) - start_index -1
			//fmt.Println(new_value)
			fmt.Println(unsorted_arr)
		}
	}
	fmt.Println("Sort float array in ascending order",unsorted_arr)
}

func Bubble_sort_string() {
	//declare variable
	var temp_value string
	// declare array
	unsorted_arr := []string{"shalini","apple","chitra","rajashri","mango"}
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
			//new_value := len(unsorted_arr) - start_index -1
			//fmt.Println(new_value)
			fmt.Println(unsorted_arr)
		}
	}
	fmt.Println("Sort  string array ",unsorted_arr)
}
func main() {
	Ascen_bubble_sort_integer()
	Desc_bubble_sort_integer()
	Bubble_sort_string()
	Bubble_sort_float()
}

/*
go run 18_bubble_sort.go
[2 8 6 33 5 9 99 37 66 24 1 4]
[2 6 8 33 5 9 99 37 66 24 1 4]
[2 6 8 33 5 9 99 37 66 24 1 4]
[2 6 8 5 33 9 99 37 66 24 1 4]
[2 6 8 5 9 33 99 37 66 24 1 4]
[2 6 8 5 9 33 99 37 66 24 1 4]
[2 6 8 5 9 33 37 99 66 24 1 4]
[2 6 8 5 9 33 37 66 99 24 1 4]
[2 6 8 5 9 33 37 66 24 99 1 4]
[2 6 8 5 9 33 37 66 24 1 99 4]
[2 6 8 5 9 33 37 66 24 1 4 99]
[2 6 8 5 9 33 37 66 24 1 4 99]
[2 6 8 5 9 33 37 66 24 1 4 99]
[2 6 5 8 9 33 37 66 24 1 4 99]
[2 6 5 8 9 33 37 66 24 1 4 99]
[2 6 5 8 9 33 37 66 24 1 4 99]
[2 6 5 8 9 33 37 66 24 1 4 99]
[2 6 5 8 9 33 37 66 24 1 4 99]
[2 6 5 8 9 33 37 24 66 1 4 99]
[2 6 5 8 9 33 37 24 1 66 4 99]
[2 6 5 8 9 33 37 24 1 4 66 99]
[2 6 5 8 9 33 37 24 1 4 66 99]
[2 5 6 8 9 33 37 24 1 4 66 99]
[2 5 6 8 9 33 37 24 1 4 66 99]
[2 5 6 8 9 33 37 24 1 4 66 99]
[2 5 6 8 9 33 37 24 1 4 66 99]
[2 5 6 8 9 33 37 24 1 4 66 99]
[2 5 6 8 9 33 24 37 1 4 66 99]
[2 5 6 8 9 33 24 1 37 4 66 99]
[2 5 6 8 9 33 24 1 4 37 66 99]
[2 5 6 8 9 33 24 1 4 37 66 99]
[2 5 6 8 9 33 24 1 4 37 66 99]
[2 5 6 8 9 33 24 1 4 37 66 99]
[2 5 6 8 9 33 24 1 4 37 66 99]
[2 5 6 8 9 33 24 1 4 37 66 99]
[2 5 6 8 9 24 33 1 4 37 66 99]
[2 5 6 8 9 24 1 33 4 37 66 99]
[2 5 6 8 9 24 1 4 33 37 66 99]
[2 5 6 8 9 24 1 4 33 37 66 99]
[2 5 6 8 9 24 1 4 33 37 66 99]
[2 5 6 8 9 24 1 4 33 37 66 99]
[2 5 6 8 9 24 1 4 33 37 66 99]
[2 5 6 8 9 24 1 4 33 37 66 99]
[2 5 6 8 9 1 24 4 33 37 66 99]
[2 5 6 8 9 1 4 24 33 37 66 99]
[2 5 6 8 9 1 4 24 33 37 66 99]
[2 5 6 8 9 1 4 24 33 37 66 99]
[2 5 6 8 9 1 4 24 33 37 66 99]
[2 5 6 8 9 1 4 24 33 37 66 99]
[2 5 6 8 1 9 4 24 33 37 66 99]
[2 5 6 8 1 4 9 24 33 37 66 99]
[2 5 6 8 1 4 9 24 33 37 66 99]
[2 5 6 8 1 4 9 24 33 37 66 99]
[2 5 6 8 1 4 9 24 33 37 66 99]
[2 5 6 1 8 4 9 24 33 37 66 99]
[2 5 6 1 4 8 9 24 33 37 66 99]
[2 5 6 1 4 8 9 24 33 37 66 99]
[2 5 6 1 4 8 9 24 33 37 66 99]
[2 5 1 6 4 8 9 24 33 37 66 99]
[2 5 1 4 6 8 9 24 33 37 66 99]
[2 5 1 4 6 8 9 24 33 37 66 99]
[2 1 5 4 6 8 9 24 33 37 66 99]
[2 1 4 5 6 8 9 24 33 37 66 99]
[1 2 4 5 6 8 9 24 33 37 66 99]
[1 2 4 5 6 8 9 24 33 37 66 99]
[1 2 4 5 6 8 9 24 33 37 66 99]
Sort array in ascending order [1 2 4 5 6 8 9 24 33 37 66 99]
[8 2 6 33 5 9 99 37 66 24 1 4]
[8 6 2 33 5 9 99 37 66 24 1 4]
[8 6 33 2 5 9 99 37 66 24 1 4]
[8 6 33 5 2 9 99 37 66 24 1 4]
[8 6 33 5 9 2 99 37 66 24 1 4]
[8 6 33 5 9 99 2 37 66 24 1 4]
[8 6 33 5 9 99 37 2 66 24 1 4]
[8 6 33 5 9 99 37 66 2 24 1 4]
[8 6 33 5 9 99 37 66 24 2 1 4]
[8 6 33 5 9 99 37 66 24 2 1 4]
[8 6 33 5 9 99 37 66 24 2 4 1]
[8 6 33 5 9 99 37 66 24 2 4 1]
[8 33 6 5 9 99 37 66 24 2 4 1]
[8 33 6 5 9 99 37 66 24 2 4 1]
[8 33 6 9 5 99 37 66 24 2 4 1]
[8 33 6 9 99 5 37 66 24 2 4 1]
[8 33 6 9 99 37 5 66 24 2 4 1]
[8 33 6 9 99 37 66 5 24 2 4 1]
[8 33 6 9 99 37 66 24 5 2 4 1]
[8 33 6 9 99 37 66 24 5 2 4 1]
[8 33 6 9 99 37 66 24 5 4 2 1]
[33 8 6 9 99 37 66 24 5 4 2 1]
[33 8 6 9 99 37 66 24 5 4 2 1]
[33 8 9 6 99 37 66 24 5 4 2 1]
[33 8 9 99 6 37 66 24 5 4 2 1]
[33 8 9 99 37 6 66 24 5 4 2 1]
[33 8 9 99 37 66 6 24 5 4 2 1]
[33 8 9 99 37 66 24 6 5 4 2 1]
[33 8 9 99 37 66 24 6 5 4 2 1]
[33 8 9 99 37 66 24 6 5 4 2 1]
[33 8 9 99 37 66 24 6 5 4 2 1]
[33 9 8 99 37 66 24 6 5 4 2 1]
[33 9 99 8 37 66 24 6 5 4 2 1]
[33 9 99 37 8 66 24 6 5 4 2 1]
[33 9 99 37 66 8 24 6 5 4 2 1]
[33 9 99 37 66 24 8 6 5 4 2 1]
[33 9 99 37 66 24 8 6 5 4 2 1]
[33 9 99 37 66 24 8 6 5 4 2 1]
[33 9 99 37 66 24 8 6 5 4 2 1]
[33 99 9 37 66 24 8 6 5 4 2 1]
[33 99 37 9 66 24 8 6 5 4 2 1]
[33 99 37 66 9 24 8 6 5 4 2 1]
[33 99 37 66 24 9 8 6 5 4 2 1]
[33 99 37 66 24 9 8 6 5 4 2 1]
[33 99 37 66 24 9 8 6 5 4 2 1]
[99 33 37 66 24 9 8 6 5 4 2 1]
[99 37 33 66 24 9 8 6 5 4 2 1]
[99 37 66 33 24 9 8 6 5 4 2 1]
[99 37 66 33 24 9 8 6 5 4 2 1]
[99 37 66 33 24 9 8 6 5 4 2 1]
[99 37 66 33 24 9 8 6 5 4 2 1]
[99 37 66 33 24 9 8 6 5 4 2 1]
[99 66 37 33 24 9 8 6 5 4 2 1]
[99 66 37 33 24 9 8 6 5 4 2 1]
[99 66 37 33 24 9 8 6 5 4 2 1]
[99 66 37 33 24 9 8 6 5 4 2 1]
[99 66 37 33 24 9 8 6 5 4 2 1]
[99 66 37 33 24 9 8 6 5 4 2 1]
[99 66 37 33 24 9 8 6 5 4 2 1]
[99 66 37 33 24 9 8 6 5 4 2 1]
[99 66 37 33 24 9 8 6 5 4 2 1]
[99 66 37 33 24 9 8 6 5 4 2 1]
[99 66 37 33 24 9 8 6 5 4 2 1]
[99 66 37 33 24 9 8 6 5 4 2 1]
[99 66 37 33 24 9 8 6 5 4 2 1]
[99 66 37 33 24 9 8 6 5 4 2 1]
Sort array in descending order [99 66 37 33 24 9 8 6 5 4 2 1]
[apple shalini chitra rajashri mango]
[apple chitra shalini rajashri mango]
[apple chitra rajashri shalini mango]
[apple chitra rajashri mango shalini]
[apple chitra rajashri mango shalini]
[apple chitra rajashri mango shalini]
[apple chitra mango rajashri shalini]
[apple chitra mango rajashri shalini]
[apple chitra mango rajashri shalini]
[apple chitra mango rajashri shalini]
Sort  string array  [apple chitra mango rajashri shalini]
[6.9 7.2 8.6 9.1 10.09 2.4 2.7 4.5 5.8 3.6]
[6.9 7.2 8.6 9.1 10.09 2.4 2.7 4.5 5.8 3.6]
[6.9 7.2 8.6 9.1 10.09 2.4 2.7 4.5 5.8 3.6]
[6.9 7.2 8.6 9.1 10.09 2.4 2.7 4.5 5.8 3.6]
[6.9 7.2 8.6 9.1 2.4 10.09 2.7 4.5 5.8 3.6]
[6.9 7.2 8.6 9.1 2.4 2.7 10.09 4.5 5.8 3.6]
[6.9 7.2 8.6 9.1 2.4 2.7 4.5 10.09 5.8 3.6]
[6.9 7.2 8.6 9.1 2.4 2.7 4.5 5.8 10.09 3.6]
[6.9 7.2 8.6 9.1 2.4 2.7 4.5 5.8 3.6 10.09]
[6.9 7.2 8.6 9.1 2.4 2.7 4.5 5.8 3.6 10.09]
[6.9 7.2 8.6 9.1 2.4 2.7 4.5 5.8 3.6 10.09]
[6.9 7.2 8.6 9.1 2.4 2.7 4.5 5.8 3.6 10.09]
[6.9 7.2 8.6 2.4 9.1 2.7 4.5 5.8 3.6 10.09]
[6.9 7.2 8.6 2.4 2.7 9.1 4.5 5.8 3.6 10.09]
[6.9 7.2 8.6 2.4 2.7 4.5 9.1 5.8 3.6 10.09]
[6.9 7.2 8.6 2.4 2.7 4.5 5.8 9.1 3.6 10.09]
[6.9 7.2 8.6 2.4 2.7 4.5 5.8 3.6 9.1 10.09]
[6.9 7.2 8.6 2.4 2.7 4.5 5.8 3.6 9.1 10.09]
[6.9 7.2 8.6 2.4 2.7 4.5 5.8 3.6 9.1 10.09]
[6.9 7.2 2.4 8.6 2.7 4.5 5.8 3.6 9.1 10.09]
[6.9 7.2 2.4 2.7 8.6 4.5 5.8 3.6 9.1 10.09]
[6.9 7.2 2.4 2.7 4.5 8.6 5.8 3.6 9.1 10.09]
[6.9 7.2 2.4 2.7 4.5 5.8 8.6 3.6 9.1 10.09]
[6.9 7.2 2.4 2.7 4.5 5.8 3.6 8.6 9.1 10.09]
[6.9 7.2 2.4 2.7 4.5 5.8 3.6 8.6 9.1 10.09]
[6.9 2.4 7.2 2.7 4.5 5.8 3.6 8.6 9.1 10.09]
[6.9 2.4 2.7 7.2 4.5 5.8 3.6 8.6 9.1 10.09]
[6.9 2.4 2.7 4.5 7.2 5.8 3.6 8.6 9.1 10.09]
[6.9 2.4 2.7 4.5 5.8 7.2 3.6 8.6 9.1 10.09]
[6.9 2.4 2.7 4.5 5.8 3.6 7.2 8.6 9.1 10.09]
[2.4 6.9 2.7 4.5 5.8 3.6 7.2 8.6 9.1 10.09]
[2.4 2.7 6.9 4.5 5.8 3.6 7.2 8.6 9.1 10.09]
[2.4 2.7 4.5 6.9 5.8 3.6 7.2 8.6 9.1 10.09]
[2.4 2.7 4.5 5.8 6.9 3.6 7.2 8.6 9.1 10.09]
[2.4 2.7 4.5 5.8 3.6 6.9 7.2 8.6 9.1 10.09]
[2.4 2.7 4.5 5.8 3.6 6.9 7.2 8.6 9.1 10.09]
[2.4 2.7 4.5 5.8 3.6 6.9 7.2 8.6 9.1 10.09]
[2.4 2.7 4.5 5.8 3.6 6.9 7.2 8.6 9.1 10.09]
[2.4 2.7 4.5 3.6 5.8 6.9 7.2 8.6 9.1 10.09]
[2.4 2.7 4.5 3.6 5.8 6.9 7.2 8.6 9.1 10.09]
[2.4 2.7 4.5 3.6 5.8 6.9 7.2 8.6 9.1 10.09]
[2.4 2.7 3.6 4.5 5.8 6.9 7.2 8.6 9.1 10.09]
[2.4 2.7 3.6 4.5 5.8 6.9 7.2 8.6 9.1 10.09]
[2.4 2.7 3.6 4.5 5.8 6.9 7.2 8.6 9.1 10.09]
[2.4 2.7 3.6 4.5 5.8 6.9 7.2 8.6 9.1 10.09]
Sort float array in ascending order [2.4 2.7 3.6 4.5 5.8 6.9 7.2 8.6 9.1 10.09]
*/