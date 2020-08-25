package main

import (
    "fmt"
    "path/filepath"
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
    a := []Test{}
    b := B{}
    a = append(a,  b)

    fmt.Println("On Unix:")
    fmt.Println(filepath.Join("a", "b", "c"))
    fmt.Println(filepath.Join("a", "b/c"))
    fmt.Println(filepath.Join("a/b", "c"))
    fmt.Println(filepath.Join("a/b", "/c"))
}