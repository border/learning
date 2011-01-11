package main
/*
http://go.hokapoka.com/go/gos-interfaces-5/
*/

import (
    "fmt"
)

type foo struct {
    s string
}

type bar struct {
    a []byte
}

type Tester interface {
    Test() (r string)
}

func (f *foo) Test() (r string) {
    return f.s + " -> In Test of foo"
}

func (b *bar) Test() (r string) {
    return string(b.a) + " -> In Test of bar"
}

func receiveTester(o Tester) (r string) {
    return o.Test() + " -> via receiveTest"
}

func main() {
    f := &foo{s: "Type foo"}
    b := &bar{a: []byte("Type bar")}

    s1 := receiveTester(f)
    s2 := receiveTester(b)

    fmt.Println("s1: " + s1)
    fmt.Println("s2: " + s2)
}
