# Type Conversion

In Go, type conversion is the process of converting a value from one type to another. This is necessary when you want to assign a value of one type to a variable of another type, or when you need to perform operations that involve values of different types. Here's how type conversion works in Go:

## Basic Syntax:
newType(valueOfTypeToConvert)
newType: The target type to which you want to convert the value.
valueOfTypeToConvert: The value you want to convert.
Example:

package main

import (
	"fmt"
)

func main() {
	var integer int = 42
	var floatingPoint float64 = float64(integer)

	fmt.Println("Integer:", integer)
	fmt.Println("Floating Point:", floatingPoint)
}

In this example, we convert an int to a float64 using the float64() conversion function.


## Notes:
1. Implicit Type Conversion:

Go does not allow implicit type conversion between different types. You need to explicitly use a conversion function or type assertion. Type assertion is part of interface, so can be ignored for now.

2. Numeric Type Conversion:

Converting between numeric types (e.g., int to float64) is common, and Go requires an explicit conversion.

3. Type Assertion with Interfaces:

Type assertion is commonly used with interface types to retrieve the underlying value and type.

4. String Conversion:

Converting between string and other types is also explicit. You can use the strconv package for more complex string conversions.

5. Compatibility:

Not all types are compatible for conversion. For example, you cannot convert a custom type to a built-in type directly.
Type conversion in Go is explicit and requires you to use conversion functions or type assertion, ensuring a clear and intentional transformation of values between different types.