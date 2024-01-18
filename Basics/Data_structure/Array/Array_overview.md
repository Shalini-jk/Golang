Array is a type of data structure in go language. Array is a collection of fixed size  element of same type and the size of array is defined at the time of declaration and cannot be changed during the tyme of excution of the program. if we do not define the size of array at the time of declaration then by default it is declared as a slice not an array. 

the operations which can be performed on array are as follow:

    1. Declaration and initialization
    Declare an array: var arrayName [size]dataType.
    numbers := [5]int{1, 2, 3, 4, 5} (initialization of an array)

    2. Accessing Elements:
    Access individual elements using their index
    num := arr1[index]

    3. Updating Elements:
    Modify the value of an element by assigning a new value:
    numbers[index] = newValue

    4. Length:
    Get the length (number of elements) of an array:
    len(arr1)

    5. Iterating Through Array:
    Use a loop to iterate through the elements:
    for i := 0; i < len(numbers); i++ {
    // process array elements
    }
    Alternatively, use the range keyword:
    for index, value := range numbers {
    // process array elements
    }

    6. Copying Arrays:
    Copy the contents of one array into another:
    copy(arr3[:],arr1[:])
    now during copy first search at the last index of arr3 after coping arr1 and then copy arr2
	copy(arr3[len(arr1):], arr2[:])

    7. Comparing Arrays:
    Compare two arrays for equality:

    8. Multidimensional Arrays:
    Declare and initialize a multidimensional array:
    var matrix [3][3]int

    9. Passing Arrays to Functions:
    Pass an array as a parameter to a function:
    func processArray(arr [5]int) {
    // function logic
    }

    10.Arrays as Values:
    Arrays in Go are value types. When passed to a function or assigned to another variable, a copy of the array is made.
    
    11. Sorting:
    Sort the elements of an array using the sort package:
    import "sort"
    sort.Ints(numbers)

    12. Searching:
    Perform searching operations, such as linear search or binary search, on the array.





