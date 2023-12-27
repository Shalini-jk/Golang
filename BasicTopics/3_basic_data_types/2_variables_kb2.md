 In GoLang, the var keyword is used to declare variables. It's a fundamental concept in Go and is used to allocate memory for a variable and associate a name with that memory location. Here's an explanation of the var concept in GoLang:

Declaration Syntax:

var variableName dataType
var: Keyword used to declare a variable.
variableName: The name of the variable.
dataType: The type of the variable.
Example:

package main

import "fmt"

func main() {
    var age int
    age = 25

    var height float64
    height = 1.75

    var isAdult bool
    isAdult = age >= 18

    fmt.Println("Age:", age)
    fmt.Println("Height:", height)
    fmt.Println("Is Adult:", isAdult)
}
In this example:

var age int: Declares a variable named age of type int.
var height float64: Declares a variable named height of type float64.
var isAdult bool: Declares a variable named isAdult of type bool.
Short Variable Declaration:
Go also provides a shorter syntax for variable declaration and initialization within a function:

func main() {
    age := 25
    height := 1.75
    isAdult := age >= 18

    fmt.Println("Age:", age)
    fmt.Println("Height:", height)
    fmt.Println("Is Adult:", isAdult)
}
Syntax:
variableName := value
Type Inference:
The type of the variable is inferred from the provided value.
#Key Points:
1. Zero Values:

If a variable is declared without an explicit initialization, it is assigned the zero value for its type. For numeric types, the zero value is 0, and for bool, it is false.
2. Redeclaration:

Variables cannot be redeclared in the same scope.

3. Short Variable Declaration:

The short variable declaration (:=) is often preferred for concise code when initializing variables.
4. Explicit Type Declaration:

While Go supports type inference, you can also explicitly declare the type if needed.
The var concept allows you to declare variables in a clear and explicit manner, and the short variable declaration provides a concise alternative for local variables. It's a fundamental part of Go's syntax for working with data in programs.

# reserved words
In Go, there are a set of reserved keywords that have special meanings and cannot be used as identifiers (names for variables, functions, etc.) in your program. Here is the list of reserved keywords in Go:

break       default     func        interface   select
case        defer       go          map         struct
chan        else        goto        package     switch
const       fallthrough if          range       type
continue    for         import      return      var
These keywords have specific roles in the language, and you cannot use them as names for variables, constants, functions, or other identifiers in your Go code.

It's essential to be aware of these reserved keywords to avoid naming conflicts and to write clean and readable code. If you attempt to use a reserved keyword as an identifier, you will receive a compilation error.

# Pre-declared identifiers

In Go, names like int, float64, and others are called predeclared identifiers. Predeclared identifiers are special names that are built into the language and serve as the names for fundamental types, constants, and some built-in functions.

Here are some examples of predeclared identifiers for basic types:

## Basic Types:

bool: Represents Boolean values (true or false).
int, int8, int16, int32, int64: Signed integer types of various sizes.
uint, uint8, uint16, uint32, uint64: Unsigned integer types of various sizes.
float32, float64: Floating-point types.
complex64, complex128: Complex number types.

## Constants:

true, false: Boolean constants.
nil: The zero value for pointers, channels, functions, interfaces, maps, and slices.

## Functions:

make: Used to create slices, maps, and channels.
new: Used to allocate memory for a new value.
panic, recover: Used for error handling in exceptional cases.
These predeclared identifiers are always available for use in any Go program without the need for explicit declaration. They are part of the language specification and play a crucial role in the core functionality of the language. When you see names like int or bool in Go code, you are working with these predeclared identifiers representing basic types.








