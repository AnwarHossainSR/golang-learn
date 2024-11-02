package main

import "fmt"

// Function with two parameters and a return value
func add(a int, b int) int {
    return a + b
}

// Function with multiple return values
func divide(a, b int) (int, int) {
    return a / b, a % b
}

func main() {
    sum := add(3, 4)
    fmt.Println("Sum:", sum)

    quotient, remainder := divide(10, 3)
    fmt.Println("Quotient:", quotient)
    fmt.Println("Remainder:", remainder)
}
