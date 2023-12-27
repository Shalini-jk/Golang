In Go, control flow refers to the order in which statements are executed in a program. Go provides several control flow statements that allow you to control the flow of execution based on conditions, loops, and other criteria. Here are the main control flow statements in Go:

1. Conditional Statements: if, else, else if
The if statement is used for conditional execution of code based on a Boolean expression.

if condition {
    // Code to be executed if the condition is true
} else if anotherCondition {
    // Code to be executed if anotherCondition is true
} else {
    // Code to be executed if none of the conditions are true
}
Example:

```go
Copy code
package main

import "fmt"

func main() {
    age := 25

    if age < 18 {
        fmt.Println("You are a minor.")
    } else if age >= 18 && age < 60 {
        fmt.Println("You are an adult.")
    } else {
        fmt.Println("You are a senior citizen.")
    }
}
```

2. Switch Statement: switch, case, default
The switch statement is used for selective control flow based on the value of an expression. It provides a concise way to express a multi-way branch.

```go
switch expression {
case value1:
    // Code to be executed if expression equals value1
case value2:
    // Code to be executed if expression equals value2
default:
    // Code to be executed if none of the cases match
}
```
Example:

```go
package main

import "fmt"

func main() {
    day := "Monday"

    switch day {
    case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
        fmt.Println("It's a weekday.")
    case "Saturday", "Sunday":
        fmt.Println("It's a weekend.")
    default:
        fmt.Println("Invalid day.")
    }
}
```

In this example, switch can test against multiple values, these are separated by ,. Cool feature. Other languages need multiple case statements for this.


3. Loops: for
The for statement is used for looping in Go. It supports different forms of loops, including a basic loop, a for loop with a condition, and a for loop that ranges over elements.

Basic Loop:

```go
for i := 0; i < 5; i++ {
    // Code to be executed in each iteration
}
//Loop with Condition:

for condition {
    // Code to be executed as long as the condition is true
}

// Loop with Range:
for index, value := range collection {
    // Code to be executed for each element in the collection
}
```
Example:

```go
package main

import "fmt"

func main() {
    // Basic loop
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    // Loop with condition
    j := 0
    for j < 3 {
        fmt.Println(j)
        j++
    }

    // Loop with range
    numbers := []int{1, 2, 3, 4, 5}
    for index, value := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
}
```

4. Control Transfer: break, continue, goto, fallthrough
break: Terminates the innermost for, switch, or select statement.
continue: Skips the rest of the loop body and starts the next iteration.
goto: Transfers control to a labeled statement (not commonly used and considered bad practice in most cases).
fallthrough: Used in a switch statement to fall through to the next case (not recommended in most cases).

Example:

```go

package main

import "fmt"

func main() {
    // break example
    for i := 0; i < 5; i++ {
        if i == 3 {
            break
        }
        fmt.Println(i)
    }

    // continue example
    for i := 0; i < 5; i++ {
        if i == 2 {
            continue
        }
        fmt.Println(i)
    }

    // goto example (avoid using goto)
    gotoLabel := 0
    if gotoLabel == 0 {
        fmt.Println("Using goto.")
        goto End
    }
    // ...


End:
    fmt.Println("End of the program.")
}
```

These control flow statements provide the necessary tools for writing structured and flexible Go programs. It's important to use them judiciously and maintain code readability.