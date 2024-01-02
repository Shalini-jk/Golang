In Go, an array is a fixed-size sequence of elements of the same type. The size of an array is part of its type, which means you cannot change the size of an array once it is declared. The array concept in Go is straightforward, and here are the key aspects:

Declaration:
To declare an array, you specify the type of its elements and the number of elements enclosed in square brackets:

``` go
var myArray [5]int // An array of 5 integers
```
Initialization:
You can initialize an array during declaration or later:

```go
var myArray [5]int = [5]int{1, 2, 3, 4, 5}
// OR
myArray := [5]int{1, 2, 3, 4, 5}
```

Accessing Elements:
You access elements using zero-based indices:

```go
value := myArray[2] // Access the third element (index 2)
```
Length:
You can get the length of an array using the len function:

```go
length := len(myArray)
```

Iterating:
You can iterate over the elements of an array using a for loop:

```go
for i := 0; i < len(myArray); i++ {
    fmt.Println(myArray[i])
}
// OR
for index, value := range myArray {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}
```
Array Literals:
Go provides a shorthand syntax called array literals for declaring and initializing arrays:

```go
myArray := [5]int{1, 2, 3, 4, 5}
```
Arrays are Value Types:
Assigning one array to another creates a copy of the array:

```go
arr1 := [3]int{1, 2, 3}
arr2 := arr1 // arr2 is a copy of arr1
```
Multi-dimensional Arrays:
Go supports multi-dimensional arrays:

```go
var matrix [3][3]int // A 3x3 matrix
```
Arrays vs. Slices:
Arrays have a fixed size, whereas slices are dynamic and more flexible. Slices are often preferred for most use cases due to their versatility.

```go
mySlice := []int{1, 2, 3, 4, 5} // Slice
```
In summary, arrays in Go are useful for scenarios where a fixed-size, contiguous block of memory is appropriate. Slices are more commonly used in Go because of their flexibility and dynamic nature. Understanding both arrays and slices is important for effective Go programming.