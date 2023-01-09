package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {

	start := time.Now()

	for i := 1; i < 65535; i++ {
		go func(p int) {
			conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", p))
			if err != nil {
				//log.Printf("%d closed (%s)\n", p, err)
				return
			}
			conn.Close()
			log.Printf("%d open", p)
		}(i)
		//log.Println("done")
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
