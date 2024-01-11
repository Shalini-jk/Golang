// write a program to sort the elements of array using selection sort

/*
Selection sort : in selection sort we divide the array into two part (sorted or unsorted part)
and then we pick the smallest element from the unsorted part and put it in the sorted part and
we do this again and again till the sorted arra get all the elements from unsorted part in
sorted way

Algorithm :-
  1. declare array
  2. then iterate the array elements
  3. while iterating the array elements find the minimum elements in the array
  4. then take that minimum elements and swap it with the 0th index
  5. and again find the minimum element in the array and swap it to next index where we swaped
   earlier
   6. then after completing the process print it on the terminal
*/

package main

import( 
	"fmt"
)

func Selection_sort_integer_ascen() {
	//declare array
	int_arr := []int{12, 11, 13, 5, 6, 15}
	//declare variable
	var smallest_number int

	//iterate the array elements
	for i := 0; i < len(int_arr); i++ {
		//now find the minimum element in the remaining unsorted part of the array
		smallest_number = i

		for j := i + 1; j < len(int_arr); j++ {
			if int_arr[j] < int_arr[smallest_number] {
				smallest_number = j
			}
		}

		// Swap the minimum element with the current element
		int_arr[i], int_arr[smallest_number] = int_arr[smallest_number], int_arr[i]

		// Print the current state of the array
		fmt.Println(int_arr)
	}

	fmt.Println("Sorted array in ascending order:", int_arr)
}
func Selection_sort_integer_desc() {
	//declare array
	int_arr := []int{12, 11, 13, 5, 6, 15}
	//declare variable
	var smallest_number int

	//iterate the array elements
	for i := 0; i < len(int_arr); i++ {
		//now find the minimum element in the remaining unsorted part of the array
		smallest_number = i

		for j := i + 1; j < len(int_arr); j++ {
			if int_arr[j] > int_arr[smallest_number] {
				smallest_number = j
			}
		}

		// Swap the minimum element with the current element
		int_arr[i], int_arr[smallest_number] = int_arr[smallest_number], int_arr[i]

		// Print the current state of the array
		fmt.Println(int_arr)
	}

	fmt.Println("Sorted array in ascending order:", int_arr)
}


func main() {
	Selection_sort_integer_ascen()
	Selection_sort_integer_desc()
	
}

/*
go run 19_selection_sort.go
[5 11 13 12 6 15]
[5 6 13 12 11 15]
[5 6 11 12 13 15]
[5 6 11 12 13 15]
[5 6 11 12 13 15]
[5 6 11 12 13 15]
Sorted array in ascending order: [5 6 11 12 13 15]
[15 11 13 5 6 12]
[15 13 11 5 6 12]
[15 13 12 5 6 11]
[15 13 12 11 6 5]
[15 13 12 11 6 5]
[15 13 12 11 6 5]
Sorted array in ascending order: [15 13 12 11 6 5]
*/

