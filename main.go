package main

import (
    "fmt"
    "os"
    "os/user"
    "interpreter/repl"
)


func main() {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }

    fmt.Printf("Hello %s! This is a cursed programming language!\n", user.Username)
    fmt.Printf("Feel free to type some commands\n")
    repl.Start(os.Stdin, os.Stdout)
}