package main

import ("./filestuts")

func main() {
    file := &filestuts.File{1, "String"}
    fmt.Printf("%v\n", file)
}
