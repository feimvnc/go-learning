package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	for i := 27000; i < 27020; i++ {
		conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", i))
		if err != nil {
			log.Printf("%d closed (%s)\n", i, err)
			continue
		}
		conn.Close()
		log.Printf("%d open\n", i)

	}
}
