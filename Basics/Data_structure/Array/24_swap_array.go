//write a program to swap adjacent element of array

/*
Algorithm :
   1. declare array
   2. iterate that array and store the previous index and net index every time.
   3. while storing the index swap that index.
   4. then print the swapped index.
*/

package main

import(
	"fmt"
)

func main() {
	//declare array
	arr1 := [] int{2,3,4,5,6,7}
	//declare variable
	var temp_value int
	//now iterate the array
	fmt.Println(arr1)
	for i:= 0 ; i < len(arr1) ; i++ {
		temp_value = arr1[i]
		arr1[i] = arr1[i+1]
		arr1[i+1] = temp_value
		i = i + 1
	}
	fmt.Println(arr1)
}

/*
output:

go run 24_swap_array.go
[2 3 4 5 6]
[3 2 4 6 5]
shalini@Ubuntu:~/Golang/Programs/Array$ go run 24_swap_array.go
[2 3 4 5 6 7]
[3 2 4 6 5 7]
shalini@Ubuntu:~/Golang/Programs/Array$ go run 24_swap_array.go
[2 3 4 5]
[3 2 4 5]
shalini@Ubuntu:~/Golang/Programs/Array$ go run 24_swap_array.go
[2 3 4 5]
[3 4 5 2]
shalini@Ubuntu:~/Golang/Programs/Array$ go run 24_swap_array.go
[2 3 4 5 6 7]
[3 2 4 6 5 7]
shalini@Ubuntu:~/Golang/Programs/Array$ go run 24_swap_array.go
[2 3 4 5 6 7]
[3 2 4 6 5 7]
shalini@Ubuntu:~/Golang/Programs/Array$ go run 24_swap_array.go
[2 3 4 5 6 7]
[3 2 5 4 7 6]
*/