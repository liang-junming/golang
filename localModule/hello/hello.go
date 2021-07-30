package main

import (
    "fmt"

    "example.com/greetings/morning"
    "example.com/greetings/night"

    "example.com/hello/localPackage"
)

func main() {
    message := morning.GetMorningGreeting("Allen")
    fmt.Println(message)

    message = night.GetNightGreeting("Allen")
    fmt.Println(message)

    message = morning.GetMorningGreeting2("Allen")
    fmt.Println(message)

    message = hi.GetHi("Allen")
    fmt.Println(message)
}
