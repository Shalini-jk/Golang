//write a program to delete an item from array.

/*
Algorithm :
   1. first declare an array
   2. then iterate the array element
   3. then remove that element which you want from the array while iterating
*/

package main

import(
	"fmt"
)

func main()  {
	//declare array
   arr1 := [6] int{2,4,5,6,7,8}
   fmt.Println(arr1)

   //iterate the array using for loop
   for i := 0 ; i < len(arr1) ; i++ {
      fmt.Println("print array after iteration",arr1)
   }
	
}