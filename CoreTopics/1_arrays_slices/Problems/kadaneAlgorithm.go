package main

import (
	"fmt"
)

func kadaneAlgorithm(arr []int) (int, []int) {
	maxEndingHere := arr[0]
	maxSoFar := arr[0]
	startIndex, endIndex := 0, 0
	tempStartIndex := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] > maxEndingHere+arr[i] {
			maxEndingHere = arr[i]
			tempStartIndex = i
		} else {
			maxEndingHere += arr[i]
		}

		if maxEndingHere > maxSoFar {
			maxSoFar = maxEndingHere
			startIndex = tempStartIndex
			endIndex = i
		}
	}

	return maxSoFar, arr[startIndex : endIndex+1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Creating array
	arr := []int{1, -8, 5, 6}

	// Find the maximum sum subarray using Kadane's Algorithm
	maxSum, maxSumSubarray := kadaneAlgorithm(arr)

	// Print the result
	fmt.Printf("Maximum Sum Subarray: %v, Sum: %d\n", maxSumSubarray, maxSum)
}
