package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"time"

	"golang.org/x/sync/semaphore"
)

var host string
var ports string
var timeout int

func init() {
	flag.StringVar(&host, "host", "127.0.0.1", "host to scan")
	flag.StringVar(&ports, "ports", "80", "port(s), 80, 27000-27020")
	flag.IntVar(&timeout, "timeout", 5, "timeout in seconds (default is 5)")
	rand.Seed(time.Now().UnixNano())
}
func main() {
	flag.Parse()

	var openPorts []int
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		printResults(openPorts)
		os.Exit(0)
	}()

	portsToScan, err := parsePortsToScan(ports)
	//fmt.Println(portsToScan, ports)
	if err != nil {
		fmt.Printf("failed to parse ports to scan: %s\n", err)
		os.Exit(1)
	}

	var semMaxWeight int64 = 100_100
	var semAcquisitionWeight int64 = 100

	sem := semaphore.NewWeighted(semMaxWeight)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	for _, port := range portsToScan {
		if err := sem.Acquire(ctx, semAcquisitionWeight); err != nil {
			fmt.Printf("failed to acquire semaphore (port %d): %v\n", port, err)
			break
		}

		go func(port int) {
			defer sem.Release(semAcquisitionWeight)
			sleepy(10)
			p := scan(host, port)
			if p != 0 {
				openPorts = append(openPorts, p)
			}
		}(port) // pass parameter of same port value
		// intend to use loop's value, and change port value
		// can scan concurrently
	}

	// block here until done
	sem.Acquire(ctx, semMaxWeight)
	printResults(openPorts)

}

func parsePortsToScan(portsFlag string) ([]int, error) {

	ports := strings.Split(portsFlag, "-")
	//fmt.Println("xx", ports)
	if len(ports) != 2 {
		return nil, errors.New("unable to determine port(s) to scan")
	}

	minPort, err := strconv.Atoi(ports[0])
	if err != nil {
		return nil, fmt.Errorf("failed to convert %s to a valid port number", ports[0])
	}

	maxPort, err := strconv.Atoi(ports[1])
	if err != nil {
		return nil, fmt.Errorf("failed to convert %s to a valid port number", ports[0])
	}

	if minPort <= 0 || maxPort <= 0 {
		return nil, fmt.Errorf("port numbers must be greater than 0")
	}

	var results []int
	for p := minPort; p <= maxPort; p++ {
		results = append(results, p)
		// fmt.Println("xx", minPort, maxPort, p)
	}
	return results, nil
}

func scan(host string, port int) int {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("%d closed (%s)\n", port, err)
		return 0
	}
	conn.Close()
	return port
}

func sleepy(max int) {
	n := rand.Intn(max)
	time.Sleep(time.Duration(n) * time.Second)
}

func printResults(ports []int) {
	sort.Ints(ports)
	fmt.Println("\nResults\n---")
	for _, p := range ports {
		fmt.Printf("%d - open\n", p)
	}
}
