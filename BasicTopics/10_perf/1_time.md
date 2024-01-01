You can use the time package in Go to measure the elapsed time spent in a function. Here's a simple example demonstrating how to use the time package to measure the execution time of a function:

```go
package main

import (
	"fmt"
	"time"
)

// Function to be measured
func exampleFunction() {
	// Simulate some work
	for i := 0; i < 1000000; i++ {
		_ = i
	}
}

func main() {
	// Record the start time
	startTime := time.Now()

	// Call the function to be measured
	exampleFunction()

	// Record the end time
	endTime := time.Now()

	// Calculate the elapsed time
	elapsedTime := endTime.Sub(startTime)

	// Print the elapsed time
	fmt.Printf("Function took %s\n", elapsedTime)
}
```
In this example:

The time.Now() function is used to record the current time before and after the function execution.
The Sub method is used to calculate the duration between the start and end times.
The resulting elapsedTime represents the duration the function took to execute.
Remember, if your function has parameters or returns values, you can modify the example accordingly. This approach is suitable for measuring the elapsed time for functions with relatively short execution times. For more accurate measurements, consider running the function multiple times and calculating the average elapsed time.

Additionally, if you want to profile the execution time of a specific function multiple times, you might find the testing.B type used in benchmark tests helpful, as it automatically manages the timing and iterations for you.