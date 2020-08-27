package main

import (
    "fmt"
    // "path/filepath"
)

type Test interface {
    Tfunc(test string)
}

type A []Test

type B struct {

}

func (self B) Tfunc(test string) {

}

func main() {
    var a uint32 = 1010102333
    b := int(a)
    fmt.Printf("%T\t%T\n",a,b)
}