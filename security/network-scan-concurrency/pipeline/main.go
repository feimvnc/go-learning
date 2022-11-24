package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var ports string
var outFile string

func init() {
	flag.StringVar(&ports, "ports", "27017-27020", "port(s), 80, 27000-27020")
	flag.StringVar(&outFile, "outfile", "scans.csv", "scan result file scans.csv")
}

func main() {
	flag.Parse()
	portsToScan, err := parsePortsToScan(ports)
	if err != nil {
		if err != nil {
			fmt.Printf("failed to parse ports to scan: %s\n", err)
			os.Exit(1)
		}
	}

	dest, err := os.Create(outFile)
	if err != nil {
		fmt.Printf("failed to create scan resutls file: %s\n", err)
		os.Exit(2)
	}

	//pipeline
	scanChan := store(dest, filter(scan(gen(portsToScan...))))

	//unfiltered
	//scannChan := store(dest, scan(gen(portsToScan...)))

	// break up for explainability
	// var scanChan <- chan scanOp
	// scanChan = gen(portsToScan...)
	// scanChan = scan(scanChan)
	// scanChan = filter(scanChan)
	// scanChan = store(dest, scanChan)

	for s := range scanChan {
		if !s.open && s.scanErr != fmt.Sprintf("dial 127.0.0.1:%d: connect: refused", s.port) {
			fmt.Println(s.scanErr)
		}
	}

	fmt.Print("completed, check scans.csv for results\n\n ")
}

func parsePortsToScan(portsFlag string) ([]int, error) {
	p, err := strconv.Atoi(portsFlag) // single port
	if err == nil {
		return []int{p}, nil
	}
	ports := strings.Split(portsFlag, "-")
	if len(ports) != 2 {
		return nil, errors.New("unable to determins port(s) to scann")
	}

	minPort, err := strconv.Atoi(ports[0])
	if err != nil {
		return nil, fmt.Errorf("failed to convert %s to a valid port number", ports[0])
	}
	maxPort, err := strconv.Atoi(ports[1])
	if err != nil {
		return nil, fmt.Errorf("failed to convert %s to a valid port number", ports[1])
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

type scanOp struct {
	port         int
	open         bool
	scanErr      string
	scanDuration time.Duration
}

func (so scanOp) csvHeaders() []string {
	return []string{"port", "open", "scanError", "scanDuration"}
}

func (so scanOp) asSlice() []string {
	return []string{
		strconv.FormatInt(int64(so.port), 10),
		strconv.FormatBool(so.open),
		so.scanErr,
		so.scanDuration.String(),
	}
}

func gen(ports ...int) <-chan scanOp {
	out := make(chan scanOp, len(ports))
	go func() {
		defer close(out)
		for _, p := range ports {
			out <- scanOp{port: p}
		}
	}()
	return out
}

func scan(in <-chan scanOp) <-chan scanOp {
	out := make(chan scanOp)
	go func() {
		defer close(out)
		for scan := range in {
			address := fmt.Sprintf("127.0.0.1:%d", scan.port)
			start := time.Now()
			conn, err := net.Dial("tcp", address)
			scan.scanDuration = time.Since(start)
			if err != nil {
				scan.scanErr = err.Error()
			} else {
				conn.Close()
				scan.open = true
			}
			out <- scan
		}
	}()
	return out
}

func filter(in <-chan scanOp) <-chan scanOp {
	out := make(chan scanOp)
	go func() {
		defer close(out)
		for scan := range in {
			if scan.open {
				out <- scan
			}
		}
	}()
	return out
}

func store(file io.Writer, in <-chan scanOp) <-chan scanOp {
	csvWriter := csv.NewWriter(file)
	out := make(chan scanOp)
	go func() {
		defer csvWriter.Flush()
		defer close(out)
		var headerWritten bool
		for scan := range in {
			if !headerWritten {
				headers := scan.csvHeaders()
				if err := csvWriter.Write(headers); err != nil {
					fmt.Println(err)
					break
				}
				headerWritten = true
			}
			values := scan.asSlice()
			if err := csvWriter.Write(values); err != nil {
				fmt.Println(err)
				break
			}
		}
	}()
	return out
}
