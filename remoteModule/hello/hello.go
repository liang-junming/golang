package main

import (
    "fmt"

    "github.com/liang-junming/golang/remoteModule/say"
)

func main() {
    message := say.Say("hello")
    fmt.Println(message)
}
