package main

import (
    "fmt"
    "time"
    "unsafe"
)

func main() {
    t := time.Now()
    fmt.Printf("a: %T, %d\n", t, unsafe.Sizeof(t))
}
