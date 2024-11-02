package main

import (
	"fmt"
	"time"
)

func printNumbers() {
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
        time.Sleep(100 * time.Millisecond)
    }
}

func printLetters() {
    for ch := 'a'; ch <= 'e'; ch++ {
        fmt.Println(string(ch))
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    go printNumbers() // Run in a separate goroutine
    go printLetters() // Run in a separate goroutine

    time.Sleep(1 * time.Second) // Wait for goroutines to finish
}
