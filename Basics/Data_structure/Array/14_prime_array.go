// write a program to get all prime number in the arrray
/*
Algorithm :
  1. first declare the array
  2. then iterate the array element. 
  3. while iterating the array element we check the element is prime or not  appling prime umer logic
  4. stored all the no which is prime no in new arra and then print that arary
*/

package main

import(
	"fmt"
	"math"
)

func main() {
	//declare array
	
	arr1 := []int{2,3,4,5,6,7,8,9} 
	var prime_arr[]int
	// here we want an array with element{3,5,7}

	//iterate the element of array
	for i := 0 ; i < len(arr1) ; i++ {
		//fmt.Println(arr1[i])
		a := math.Sqrt(float64(arr1[i]))
		//fmt.Println("square root of each element of array",a)
		// now use for loop to check that the iterated element is prime or not
		var check_prime = true

		for j := 2 ; j <= int(a) ; j++ {
			if arr1[i] % j == 0 {
				check_prime = false
				//break
			}
		}
        if check_prime && arr1[i] > 1 {
			prime_arr = append(prime_arr, arr1[i])
			//fmt.Println("Prime number:", prime_arr)		
		}
	}
	fmt.Println("original array",arr1)
	fmt.Println("Prime number:", prime_arr)

}