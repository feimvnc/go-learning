package main

import (
    "fmt"
    "time"
    "unsafe"
)

func main() {
    t := time.Now()
    fmt.Printf("a: %T, %d\n", t, unsafe.Sizeof(t))
	// static var type int64
	var i int64 = 100
    fmt.Printf("a: %T, %d\n", i, unsafe.Sizeof(i))
	// static variable declaration 
	var j = 100
	fmt.Printf("a: %T, %d\n", j, unsafe.Sizeof(j))
	// dynamic declaration 
	m := 100
	fmt.Printf("a: %T, %d\n", m, unsafe.Sizeof(m))
	a := [32]byte{}
	fmt.Printf("a: %T, %d\n", a, unsafe.Sizeof(a))
	aa := string(a[:])
	fmt.Printf("a: %T, %d\n", a, unsafe.Sizeof(aa))
	b := []byte("abcd")
	fmt.Println(b)
	c := string(b[:])
	fmt.Println(c)
	var v int64 
	fmt.Println(v)

	var Start1851 = -time.Date(1851, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano()
	fmt.Println(Start1851)
	var Start1851b = time.Date(1851, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano()
	fmt.Println(Start1851b)

	
}
