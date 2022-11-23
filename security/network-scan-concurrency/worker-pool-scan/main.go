package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

var host string
var ports string
var numWorkers int

func init() {
	flag.StringVar(&host, "host", "127.0.0.1", "host to scan")
	flag.StringVar(&ports, "ports", "27000-27020", "port(s), 80, 22-1024")
	flag.IntVar(&numWorkers, "workers", runtime.NumCPU(), "number of workers, defaults to system's number of cpus")
}

func main() {
	flag.Parse()
	var openPorts []int

	// handle system interruption , capture user signal
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		printResults(openPorts)
		os.Exit(0)
	}()

	portsToScan, err := parsePortsToScan(ports)
	//fmt.Println(portsToScan)
	if err != nil {
		fmt.Printf("failed to parse ports to scan: %s\n", err)
		os.Exit(1)
	}

	portsChan := make(chan int, numWorkers) // create a buffered chan
	resultsChan := make(chan int)

	for i := 0; i < cap(portsChan); i++ { // feeds numWorkers chan to workers
		go worker(host, portsChan, resultsChan) // create new work to do work
	}

	go func() {
		for _, p := range portsToScan {
			portsChan <- p
		}
	}()

	for i := 0; i < len(portsToScan); i++ {
		if p := <-resultsChan; p != 0 { // non-zero port means it is open
			openPorts = append(openPorts, p)
		}
	}
	close(portsChan)
	close(resultsChan)
	printResults(openPorts)

}

// portsChan <- chan int, read only chan
// resultsChan chan <- int , write to chan
func worker(host string, portsChan <-chan int, resultsChan chan<- int) {
	for p := range portsChan { // as long as portsChan has values
		address := fmt.Sprintf("%s:%d", host, p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("%d closed (%s)\n", p, err)
			resultsChan <- 0 // send 0 to indicate work is done
			continue
		}
		conn.Close()
		resultsChan <- p // till here , port is open and send to resultsChan

	}
}

func printResults(ports []int) {
	//sort.Ints(ports)
	fmt.Println("\nResults\n---x")
	for _, p := range ports {
		fmt.Printf("%d - open\n", p)
	}
}

func parsePortsToScan(portsFlag string) ([]int, error) {
	p, err := strconv.Atoi(portsFlag)
	if err == nil {
		return []int{p}, err
	}

	ports := strings.Split(portsFlag, "-")
	if len(ports) != 2 {
		return nil, errors.New("unable to determine port(s) to scan")
	}

	minPort, err := strconv.Atoi(ports[0])
	if err != nil {
		return nil, fmt.Errorf("failed to convert %s to a valid port number")
	}
	maxPort, err := strconv.Atoi(ports[1])
	if err != nil {
		return nil, fmt.Errorf("failed to convert %s to a valid port number")
	}

	if minPort <= 0 || maxPort <= 0 {
		return nil, fmt.Errorf("port numbers must be greater than 0")
	}

	var results []int
	for p := minPort; p <= maxPort; p++ {
		results = append(results, p)
	}
	return results, nil
}
