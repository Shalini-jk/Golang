In Go, a struct is a composite data type that groups together variables (fields) under a single name. Structs are used to represent and organize related data, similar to how a class or record works in other programming languages. Each field in a struct can have a different data type, and the fields are accessed using dot notation.

Declaring a Struct:
Here's the basic syntax for declaring a struct in Go:

```go
package main

import "fmt"

// Define a struct named 'Person'
type Person struct {
    FirstName string
    LastName  string
    Age       int
    Address   Address // Embedded struct
}

// Define another struct named 'Address'
type Address struct {
    Street  string
    City    string
    Country string
}

func main() {
    // Create an instance of the 'Person' struct
    person1 := Person{
        FirstName: "John",
        LastName:  "Doe",
        Age:       30,
        Address: Address{
            Street:  "123 Main St",
            City:    "Anytown",
            Country: "USA",
        },
    }

    // Access fields of the struct
    fmt.Println("First Name:", person1.FirstName)
    fmt.Println("Last Name:", person1.LastName)
    fmt.Println("Age:", person1.Age)
    fmt.Println("Address:", person1.Address)

    // Access embedded struct fields
    fmt.Println("Street:", person1.Address.Street)
    fmt.Println("City:", person1.Address.City)
    fmt.Println("Country:", person1.Address.Country)
}
```
In this example, the Person struct has fields FirstName, LastName, Age, and an embedded Address struct. The Address struct has its own fields (Street, City, Country). Instances of these structs can be created and accessed using dot notation.

# Initializing Structs:
Structs can be initialized in various ways:

```go
// Initialize with values
person2 := Person{
    FirstName: "Alice",
    LastName:  "Smith",
    Age:       25,
}

// Initialize without field names (order matters)
person3 := Person{"Bob", "Johnson", 35, Address{"456 Oak St", "Sometown", "USA"}}
```
# Anonymous Structs:
```go

// Anonymous struct without a defined type
anonStruct := struct {
    X int
    Y int
}{10, 20}
```
# Methods on Structs:
Go supports attaching methods to structs, allowing you to define behaviors associated with instances of a struct.

```go
func (p Person) FullName() string {
    return p.FirstName + " " + p.LastName
}

// Usage
full := person1.FullName()
fmt.Println("Full Name:", full)
```
# Tags in Structs:
You can add tags to struct fields, which are often used for encoding and decoding data, validation, or providing additional metadata.

```go
type TaggedStruct struct {
    Field1 string `json:"field_one"`
    Field2 int    `json:"field_two"`
}
```
Structs are a powerful and flexible feature in Go, enabling you to model and organize data effectively. They play a central role in Go's composition model, and their simplicity aligns with Go's design philosophy.