package main

import ("./filestuts"
        "fmt")

func main() {
    file := &filestuts.File{1, "String"}
    fmt.Printf("%v\n", file)
}
