//write a program to add two integer array

package main

import(
	"fmt"
	"reflect"
)

func main() {
	//declare two array
	first_arr := [] int{2,3,4,5}
	second_arr := [] int{6,7,8,9}
	var sum_arr[]int  //{8,10,12,14}
	var arr_sum[]int
	xType := reflect.TypeOf(arr_sum)
	fmt.Println(xType)
	arrKind := xType.Kind()
	fmt.Println(arrKind)
	//if e declare array without size then y default it  is a slice not an array


	//iterate both array
	for i := 0 ; i < len(first_arr) ; i++ {
		for j := 0 ; j < len(second_arr) ; j++ {
			sum_arr = append(sum_arr,first_arr[i] + second_arr[j])
		}
	}
	fmt.Println(sum_arr)
    
	// now Add corresponding elements and append to arr_sum
	for k := 0; k < len(first_arr); k++ {
		arr_sum = append(arr_sum, first_arr[k]+second_arr[k])
	}

	fmt.Println(arr_sum)
}

/*
go run 23_add_array.go
[8 9 10 11 9 10 11 12 10 11 12 13 11 12 13 14]
[8 10 12 14]
*/