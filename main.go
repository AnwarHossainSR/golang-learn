package main

import "fmt"

func main() {
    var name, city string = "Bob", "New York"
    age, isEmployed := 28, true

    fmt.Println("Name:", name)
    fmt.Println("City:", city)
    fmt.Println("Age:", age)
    fmt.Println("Is Employed:", isEmployed)
}
