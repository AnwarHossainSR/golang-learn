package main

import "fmt"

// Function with pointer parameter
func increment(x *int) {
    *x = *x + 1 // Modify the original value
}

func main() {
    value := 10
    increment(&value)
    fmt.Println("Incremented Value:", value) // Outputs 11
}
