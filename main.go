package main

import "fmt"

func main() {
    // If-else statement
    age := 18
    if age >= 18 {
        fmt.Println("Adult")
    } else {
        fmt.Println("Minor")
    }

    // For loop (Go's only loop structure)
    sum := 0
    for i := 1; i <= 5; i++ {
        sum += i
    }
    fmt.Println("Sum:", sum)

    // Switch statement
    day := "Monday"
    switch day {
    case "Saturday", "Sunday":
        fmt.Println("Weekend")
    default:
        fmt.Println("Weekday")
    }
}
