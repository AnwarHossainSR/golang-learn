package main

import "fmt"

// Interface definition
type Greeter interface {
    greet() string
}

// Struct implementing Greeter
type Person struct {
    name string
}

func (p Person) greet() string {
    return "Hello, " + p.name
}

func main() {
    var greeter Greeter = Person{name: "Bob"}
    fmt.Println(greeter.greet()) // Outputs "Hello, Bob"
}
