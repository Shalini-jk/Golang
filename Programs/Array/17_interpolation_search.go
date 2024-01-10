// write a program to search an element using Interpolation search algorithm.

/*
Interpolation Search Algorithm : this algorithm is used to search the element in uniformed 
array y finding proabable index using formula
probablePos = low + ((target - arr[low]) * (high - low) / (arr[high] - arr[low]))

uniformed array :  uniformed array  means the diference between two consecutive elements  are constant 
{2,4,6,8,10,12} the dirence is uniformly constant ie 2. in interpolation search we find the
probable position of the target element

probale position : in probable position we calculate the estimated index by thinking that at this
position the target element may be located 

conditions to implement proabable element :
  1. The array must be uniformly distriuted

Algorithm :
    1. declare array
	2. take input element
	3. now  assign lower inde and higher index
	4. then calculate the probable position using formula
	5. if the probable position matches the target element then return that
	6. if the target is greater than the probable position then update higher index
	7. if the target is smaller than the probable position then update lower index
*/

package main 

import(
	"fmt"
)

func main() {
	// declare the array
	arr1 := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	//declare variable
	var give_value int
	var lower_index int
	var higher_index int
	var probable_position int

	//take input from user
	fmt.Println("Enter the number which ou want to search")
	fmt.Scanln(&give_value)

	//assign lower and higher index
	lower_index = 0
	higher_index = len(arr1) - 1
	
	found_value := true
	// now start the loop amd calculate the probable position
	for lower_index <= higher_index {
		//probablePos = low + ((target - arr[low]) * (high - low) / (arr[high] - arr[low]))
		probable_position = lower_index + ((give_value - arr1[lower_index]) * (higher_index - lower_index) / (arr1[higher_index] - arr1[lower_index]))
		fmt.Println("The estimated index :",probable_position)
	    if arr1[probable_position] == give_value {
			found_value = false
		    break	
	    } else if arr1[probable_position] < give_value {
		// if the target element is lower than the proable position
		//update lower index 
		lower_index = probable_position + 1	
	    } else {
		// if the target element is higher than the proale position
		//update higher index
		higher_index = probable_position - 1	
	    }
    }
	if found_value == false {
		fmt.Println("Searched Element",give_value,"founded at estimated index",probable_position)
	} else {
		fmt.Println("Elemnt not found")
	}
}

/*
shalini@Ubuntu:~/Golang/Programs/Array$ go run 17_interpolation_search.go
Enter the number which ou want to search
16
The estimated index : 7
Searched Element 16 founded at estimated index 7
*/