// write a program to calculate area of rectangles.
// area = 2(l+b)

package main

import(
	"fmt"
)

func main()  {
	//declaring varible
	var rect_length float64
	var rect_breath float64
	var rect_area float64

	//taking input from user
	fmt.Println("Enter the length of the rectangle")
	_,err := fmt.Scanln(&rect_length)
	if err != nil {
		fmt.Println("Incorrect input",err)
	}
	
	fmt.Println("Enter the breadth of the rectangle")
	_,err = fmt.Scanln(&rect_breath)
	if err != nil {
		fmt.Println("Incorrect input",err)
	}
	
	// calculating the area of rectangles
	rect_area = 2 * (rect_length + rect_breath)
	fmt.Println("the area of rectangle is :", rect_area)
}