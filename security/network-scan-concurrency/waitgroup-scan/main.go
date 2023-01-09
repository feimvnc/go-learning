package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

var host string
var fromPort string
var toPort string

func init() {
	flag.StringVar(&host, "host", "127.0.0.1", "host to scan")
	flag.StringVar(&fromPort, "from", "1", "start port")
	flag.StringVar(&toPort, "to", "65535", "end port")
}

func main() {
	start := time.Now()

	flag.Parse()

	fp, err := strconv.Atoi(fromPort)
	if err != nil {
		log.Fatalln("invalid 'from' port")
	}

	tp, err := strconv.Atoi(toPort)
	if err != nil {
		log.Fatalln("invalid 'to' port")
	}

	if fp > tp {
		log.Fatal("invalid values for 'from' and 'to' port")
	}
	// declare wait group
	// specify a counter to wait on
	var wg sync.WaitGroup
	numGoRoutinesToWaitOn := tp - fp + 1
	// must run before wg.Done() to make it work
	wg.Add(numGoRoutinesToWaitOn)
	for i := fp; i <= tp; i++ {
		go func(p int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, p))
			if err != nil {
				//log.Printf("%d closed (%s)\n", p, err)
				return
			}
			conn.Close()
			log.Printf("%d open\n", p)
		}(i)
	}
	// tell main goroutine to wait, until wait counter is reduced
	// otherwise blocking happens
	wg.Wait()
	//log.Println("done")

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
