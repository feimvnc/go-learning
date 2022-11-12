// go run main
// (base) user:read-csv user$ go run main.go
// [1 2 3]
// [10 20 30]
// [1x3] DataFrame

//     1     2     3
//  0: 10    20    30
//     <int> <int> <int>

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func main() {
	// open file
	csvfile, err := os.Open("data/test.csv")
	if err != nil {
		log.Fatal(err)
	}
	// Do something
	// parse / process
	// method 1
	r := csv.NewReader(csvfile)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)
	}

	// method 2: using package gota, qframes, dataframe-go
	csvfile1, err := os.Open("data/test.csv")
	if err != nil {
		log.Fatal(err)
	}

	df := dataframe.ReadCSV(csvfile1)
	fmt.Println(df)
}
