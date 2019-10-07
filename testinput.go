package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Test input")
    fmt.Println("----------")

    for {
        fmt.Print("-> ")
        text, _ := reader.ReadString('\n')
        fmt.Println(text)

        text = strings.Replace(text, "\n", "", -1)
        fmt.Println(text)

        if strings.Compare("hi", text) == 0 {
            fmt.Println("hello, yourself!")
        }
    }
}