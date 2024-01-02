package main 

import (
	"fmt"
)

func main() {
	// array

	// declare  and initialize array
	arr := [4]int{1 , 2, 4, 6}

	// accessing the array element
	for i := 0 ; i < len(arr) ; i++ {
		fmt.Println("Print array  by using loop",arr[i])
	} 
	fmt.Println("Print array by using index",arr[2])

	// length of array
	length := len(arr)
	fmt.Println("length of array",length)

	// assiging value of one array to another
	arr1 := arr
	fmt.Println(arr1)

	//updating array elements
	arr[3] = 78
	fmt.Println(arr)

	// allow user to input the each element of the array
	var numbers [5]int
	fmt.Println(".........Enter the value of array..............")
	for i := 0 ; i < len(numbers) ; i++ {
		fmt. Printf("enter the element of array %d :\n",i)
		fmt.Scan(&numbers[i])
	}
	fmt.Println("array element with user input",numbers)

	//slice

    //declaration of slice
	slice1 := []int{1,7,8,9,0,6,4}
	fmt.Println(slice1)

	slice2 := make([]int, 6)
	fmt.Println(slice2)




}