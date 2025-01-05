package main

import (
    "fmt"
    "example.com/greetings"
)

func main() {
    //Define the name of the user being greeted
    message := greetings.Hello("George")
    fmt.Println(message)
}