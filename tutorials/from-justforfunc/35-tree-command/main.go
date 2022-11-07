// tree utiltiy in go, print files / directories recursively
// tree $GOPATH/src/github.com/xxx/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	args := []string{"."}
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}
	// test output, go build -o tree
	// ./tree .
	// ./tree foo bar
	// fmt.Println(args) // test output, go build -o tree

	for _, arg := range args {
		err := tree(arg, "")
		if err != nil {
			log.Printf("tree %s: %v\n", arg, err)
		}
	}
}

// 2nd way
func tree(root, indent string) error {
	fi, err := os.Stat(root) // stat path
	if err != nil {
		return fmt.Errorf("could not stat %s: %v", root, err)
	}

	// fmt.Printf("%s%s\n", indent, fi.Name())
	fmt.Println(fi.Name())
	if !fi.IsDir() {
		return nil // if not dir we are done
	}
	fis, err := ioutil.ReadDir(root)
	if err != nil {
		return fmt.Errorf("could not read dir %s:%v", root, err)
	}

	var names []string
	for _, fi := range fis {
		if fi.Name()[0] != '.' {
			names = append(names, fi.Name()) // take care file name start with '-'
		}
	}

	//add := " "
	for i, name := range names {
		// if fi.Name()[0] == '.' { // skip dir start with '.'
		// 	continue
		// }

		add := "|  "
		if i == len(names)-1 {
			fmt.Printf(indent + "|__")
			add = "  "
		} else {
			fmt.Printf(indent + "|--")

		}

		if err := tree(filepath.Join(root, name), indent+add); err != nil {
			return err
		}

	}
	return nil
}

// // use filepath.Walk library - version 1
// func tree(root string) error {
// 	fmt.Println(root)
// 	err := filepath.Walk(root, func(path string, fi os.FileInfo, err error) error {
// 		//fmt.Println(path)
// 		if err != nil {
// 			return err
// 		}
// 		if fi.Name()[0] == '.' { // first char is '.', ignore
// 			return filepath.SkipDir // skip dir
// 		}

// 		rel, err := filepath.Rel(root, path)
// 		if err != nil {
// 			return fmt.Errorf("could not rel%s, %s): %v", root, path, err)
// 		}
// 		//fmt.Println(rel)
// 		//fmt.Println(fi.Name())
// 		//fmt.Println(path)
// 		depth := len(strings.Split(rel, string(filepath.Separator)))
// 		fmt.Printf("%s%s\n", strings.Repeat(" ", depth), fi.Name())
// 		return nil
// 	})
// 	return err
// }
