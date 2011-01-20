package main

import (
        "./numbers"
        "./filestuts"
        "fmt"
        )

func main() {
    ret := numbers.Double(10)
    fmt.Println(ret)

    file := &filestuts.File{1, "String"}
    fmt.Printf("%v\n", file)
}
