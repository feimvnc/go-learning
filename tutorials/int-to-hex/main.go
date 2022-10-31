package main

import (
	"fmt"
	"strconv"
)

func main() {
	hex_string := "98"
	o, _ := strconv.ParseInt(hex_string, 16, 64)
	fmt.Printf("%d \n", o)    // hex string to int
	fmt.Printf("%d \n", o/16) // division
	fmt.Printf("%d \n", o%16) // remnant
}
