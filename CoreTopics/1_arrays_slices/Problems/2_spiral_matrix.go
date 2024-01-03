/*
The Spiral Matrix problem involves traversing a 2D matrix in a spiral order, starting from the top-left corner and moving in a clockwise direction. The task is to return all elements of the matrix in the order of their traversal. Here's a more detailed explanation:

# Problem Statement:
Given an m x n matrix, return all elements of the matrix in spiral order.

# Example:
Consider the following 3x3 matrix:

1  2  3  4   
5  6  7  8      
9 10 11 12        

The spiral order traversal of this matrix would be [1, 2, 3, 6, 9, 8, 7, 4, 5].

Algorithm :
   1. first take a matrix as input.
   2. then take a empty slice.
   3. then use a loop to traverse upper outer row
   4. then traverse the outer column
   5. then traverse the lower outer row 
   6. then traverse the second(which is middle row) row which is now became the outer
      row after completing previous traverse.
   7.  then after each traversal append the element into the empty slice. 
*/

package main

import (
	"fmt"
)

func main() {
   // created a matrix along with the value
	matrix1 := [][]int{
      {1,2,3,4}, 
      {5,6,7,8}, 
      {9,10,11,12},
   }
   // creating an empty slice
   var traversal_slice1 []int
   //var traversal_slice2 []int
   //var traversal_slice3 []int
   //var traversal_slice4 []int
   
   fmt.Println(matrix1)

   // then use a loop to traverse upper outer row
   //fmt.Println("......first row of matrix.......")
   for i := 0 ; i < len(matrix1[0])-1 ; i++ {
      //fmt.Println(matrix1[0][i])
      //append the traverse element in empty slice
      traversal_slice1 = append(traversal_slice1,matrix1[0][i]) 
   }
   //fmt.Println(traversal_slice1)

   //then traverse the outer column
   //fmt.Println(".....traverse the outer column......")
   column_index := 3 // it specify that we are using last column 
   for m := 0 ; m < len(matrix1) ; m++ {
      //fmt.Println(matrix1[m][column_index])
      //append the traverse element again in the half filled slice
      traversal_slice1 = append(traversal_slice1,matrix1[m][column_index]) 
       
   }
   //fmt.Println(traversal_slice4)

   // then traverse the last row
   //fmt.Println("..........last row of matrix......")
   for j := len(matrix1[2])- 2 ; j>=0 ; j-- {
      //fmt.Println(matrix1[2][j])
      //append the traverse element again in the half filled slice
      traversal_slice1 = append(traversal_slice1,matrix1[2][j]) 
   }
   //fmt.Println(traversal_slice1)

   //then traverse the row(middle alement of array)
   //fmt.Println(".........second row of matrix..........")
   for k := 0 ; k < len(matrix1[1])-1 ; k++ {
      //fmt.Println(matrix1[1][k])
      //append the traverse element again in the half filled slice
      traversal_slice1 = append(traversal_slice1,matrix1[1][k]) 
   }
   fmt.Println(".........sprial traversal of matrix..........")
   fmt.Println(traversal_slice1)

   
}