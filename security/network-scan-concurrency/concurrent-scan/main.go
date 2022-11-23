package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	for i := 27000; i < 27020; i++ {
		go func(p int) {
			conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", p))
			if err != nil {
				log.Printf("%d closed (%s)\n", p, err)
				return
			}
			conn.Close()
			log.Printf("%d open", p)
		}(i)
		log.Println("done")
	}
}
