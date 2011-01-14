package main

import (
        "fmt"
        )

type MyInt int
type MyString string

type 整数 MyInt
type 字符串 MyString

func main() {
    var myint MyInt = 1
    var mystring MyString = "Hello, World"
    var 我的整数 整数 = 11
    var 我的字符串 字符串 = "UTF8 中文"

    fmt.Printf("MyInt: %v\n", myint)
    fmt.Printf("MyString: %v\n", mystring)
    fmt.Printf("整数: %v\n", 我的整数)
    fmt.Printf("字符串: %v\n", 我的字符串)
}
