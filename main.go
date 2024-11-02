package main

import "fmt"

func main() {
    // Basic data types
    var age int = 30             // Integer
    var temperature float64 = 36.6 // Float
    var isStudent bool = true      // Boolean
    var name string = "Alice"      // String

    // Composite types: Array, Slice, Map, Struct
    var scores = [3]int{85, 90, 95} // Array of fixed size
    var names = []string{"John", "Doe"} // Slice (dynamic size)
    var grades = map[string]int{"math": 90, "science": 85} // Map (key-value pairs)

    // Struct: A custom data type
    type Person struct {
        name string
        age  int
    }
    var person = Person{name: "Bob", age: 32} // Initializing a struct

    fmt.Println("Age:", age)
    fmt.Println("Temperature:", temperature)
    fmt.Println("Is Student:", isStudent)
    fmt.Println("Name:", name)
    fmt.Println("Scores:", scores)
    fmt.Println("Names:", names)
    fmt.Println("Grades:", grades)
    fmt.Println("Person:", person)
}
