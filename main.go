package main

import "fmt"

// Struct definition
type Person struct {
    name string
    age  int
}

// Method for the Person struct
func (p *Person) greet() string {
    return "Hello, " + p.name
}

func main() {
    person := Person{name: "Alice", age: 25}
    fmt.Println(person.greet()) // Outputs "Hello, Alice"
}
