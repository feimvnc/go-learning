# https://www.youtube.com/watch?v=9QZdQue_-1w

package main

import "demo/pkg"

var types = []string{
	pkg.PersonalComputerType,
	pkg.NotebookType,
	pkg.ServerType,
}

func main() {
	for _, typeName := range types {
		computer := pkg.New(typeName)
		if computer == nil {
			continue
		}
		computer.PrintDetails()
	}
}
