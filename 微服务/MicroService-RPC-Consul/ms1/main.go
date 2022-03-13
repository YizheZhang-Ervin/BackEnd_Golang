package main

import (
    "configGet"
    "fmt"
)

func main() {
    config := configGet.Get()
    data := config.Get("micro", "config", "local").Bytes()
    fmt.Println(string(data))
}
