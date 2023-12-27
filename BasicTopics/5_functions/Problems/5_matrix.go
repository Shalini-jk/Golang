/*
Matrix Multiplication:

Implement a function that performs matrix multiplication. 
Given two matrices as input, return the resulting matrix. Make sure to check the compatibility 
of the matrices for multiplication.
*/
package main
import (
	"fmt"
)

func main(){
      // declaration of array
	  var matrix1 [2][2] int
	  
	  // giving value to the array
	  matrix1[0][0]= 2
	  matrix1[0][1] = 5

	  matrix1[1][0] = 6
	  matrix1[1][1] = 8

	// declaration of array
	var matrix2 [2][2] int
	  
	// giving value to the array
	matrix2[0][0]= 4
	matrix2[0][1] = 9

	matrix2[1][0] = 3
	matrix2[1][1] = 7

	 fmt.Println("First matrix")
	 fmt.Println(matrix1[0][0], matrix1[0][1])   
	 fmt.Println(matrix1[1][0], matrix1[1][1])
	 fmt.Println("second matrix")
	 fmt.Println(matrix2[0][0], matrix2[0][1])   
	 fmt.Println(matrix2[1][0], matrix2[1][1])
	 // check conditions for matri multiplications
	 // no of column of first matrix must be equal to number of row of second matrix(cols1 = rows2)
	rows1, cols1 :=  len(matrix1),len(matrix1[0])
    rows2, cols2 := len(matrix2), len(matrix2[0])

    if cols1 == rows2 {
        fmt.Println("conditions satisfied !!, the resultant matrix is of", rows1 * cols2)
		
		// matrix multiplications

		result := make([][]int, rows1)
		for i := range result {
			result[i] = make([]int, cols2)
		}

		for i := 0; i < rows1; i++ {
			for j := 0; j < cols2; j++ {
				for k := 0; k < cols1; k++ {
					result[i][j] += matrix1[i][k] * matrix2[k][j]
					
				}
			}
		}
		fmt.Println("Resultant matrix")
		fmt.Println(result[0][0], result[0][1])
		fmt.Println(result[1][0], result[1][1])

		
    }else {
        fmt.Println("Conditions does not match.")
    }


}
