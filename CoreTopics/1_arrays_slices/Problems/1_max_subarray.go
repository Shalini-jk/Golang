/*
Problem Statement:
Given an array of integers, find the contiguous subarray (containing at least one number) with the largest sum and return the sum.

Example:
Input: [-2, 1, -3, 4, -1, 2, 1, -5, 4]

Output: subarray [4, -1, 2, 1] maxSum 6

Explanation: The contiguous subarray [4, -1, 2, 1] has the largest sum of 6.

Algorithm :
 1. first we create a array
 2. then we create all the possible sub arrray from the existing array
 3. and print the sum of all the created sub array 
 4. find the largest sum and after finding that print its equivalent su array.
*/

package main 

import (
	"fmt"
)

func main() {
	// Creating array
	arr := [4]int{1, -8, 5, 6}
	// {1}, {1, -8}, {1, -8, 5}, {1, -8, 5, 6}, {-8}, {-8, 5}, {-8, 5, 6}, {5}, {5, 6}, {6}
	var all_possible_sub_array [][]int
	var sub_array []int
	var maxSum = arr[0]
	var startIndex, endIndex int

	// Initialize maximum sum and its indexes
	//maxSum := arr[0]
	startIndex = 0
	endIndex = 0

	// try to create all possible subarrays
	for start_number := 0; start_number < len(arr); start_number++ {
		for end_number := start_number + 1; end_number <= len(arr); end_number++ {
			sub_array = arr[start_number:end_number]
			// Now use append function to add the element in slice
			all_possible_sub_array = append(all_possible_sub_array, sub_array)

			// try to do the sum of the subarray
			sum := 0
			for _, num := range sub_array {
				sum += num
			}
			//fmt.Println(sum)

			// change the value of max sum if the value of sum is greater than assign that value to ma sum             
				if sum > maxSum {
				maxSum = sum
				startIndex, endIndex = start_number, end_number-1
				//fmt.Println(startIndex)
			}
		}
	}

		for _, sub_array := range all_possible_sub_array {
			fmt.Println(sub_array)
			// try to do the sum of the  each subarray
			sum := 0
			for _, num := range sub_array {
				sum += num
			}
		fmt.Printf("Subarray: %v, Sum: %d\n", sub_array, sum)
	
    }
	fmt.Println("maximum sum among all the sum of subarray",maxSum)
	fmt.Printf("Subarray along with maximum sum: %v\n", arr[startIndex:endIndex+1])
  }

	
