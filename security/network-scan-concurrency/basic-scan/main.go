package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	p := 27017
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", p))

	if err != nil {
		log.Fatalf("%d closed(%s)\n", p, err)
	}

	conn.Close()
	log.Printf("%d open\n", p)
}
