The strings package in Go provides a collection of functions for manipulating strings. These functions are designed to perform common string operations efficiently. The package is part of the Go Standard Library, and you can use it by importing the strings package in your Go program.

Here are some of the commonly used functions provided by the strings package:

1. Contains(s, substr string) bool:

Checks if the string s contains the substring substr. Returns true if substr is found in s, and false otherwise.

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    str := "Hello, World!"
    contains := strings.Contains(str, "World")
    fmt.Println(contains) // Output: true
}
```

2. Count(s, substr string) int:

Counts the number of non-overlapping instances of substr in the string s.

```go

count := strings.Count("ababab", "ab")
fmt.Println(count) // Output: 3
```

3. HasPrefix(s, prefix string) bool:

Checks if the string s starts with the specified prefix. Returns true if s has the specified prefix, and false otherwise.

```go

hasPrefix := strings.HasPrefix("Gopher", "Go")
fmt.Println(hasPrefix) // Output: true
HasSuffix(s, suffix string) bool:

```

4. Checks if the string s ends with the specified suffix. Returns true if s has the specified suffix, and false otherwise.

```go

hasSuffix := strings.HasSuffix("Gopher", "per")
fmt.Println(hasSuffix) // Output: true
Index(s, substr string) int:

```

5. Returns the index of the first occurrence of substr in s. If substr is not found, it returns -1.

```go

index := strings.Index("Hello, World!", "World")
fmt.Println(index) // Output: 7
Join(elems []string, sep string) string:

```

6. Concatenates the elements of a string slice elems into a single string using the specified sep as the separator.

```go

fruits := []string{"apple", "banana", "orange"}
result := strings.Join(fruits, ", ")
fmt.Println(result) // Output: apple, banana, orange
Repeat(s string, count int) string:

```

7. Returns a new string consisting of count copies of the string s.

```go

repeated := strings.Repeat("Go", 3)
fmt.Println(repeated) // Output: GoGoGo

```


8. Replace(s, old, new string, n int) string:

Replaces n occurrences of old with new in the string s. If n is negative, all occurrences are replaced.

```go

replaced := strings.Replace("ababab", "ab", "XY", 2)
fmt.Println(replaced) // Output: XYXYab

```

9. Split(s, sep string) []string:

Splits the string s into a slice of substrings using the specified sep as the delimiter.

```go

parts := strings.Split("apple,banana,orange", ",")
fmt.Println(parts) // Output: [apple banana orange]
ToLower(s string) string and ToUpper(s string) string:

```

10. Converts all characters in the string s to lowercase or uppercase, respectively.

``` go

lowercase := strings.ToLower("Hello, World!")
uppercase := strings.ToUpper("Hello, World!")
fmt.Println(lowercase) // Output: hello, world!
fmt.Println(uppercase) // Output: HELLO, WORLD!

```

These are just a few examples of the functions provided by the strings package. The package offers a range of functions for common string manipulations, making it a valuable tool for working with strings in Go programs.