package main

import "fmt"

// pass a string, return true if string is palindrome
func IsPalindrome(s string) bool {

	// start from 0 to len(s) / 2 mid-way
	// compare against from end backward, len(s)-i-1, "-1" for 0-indexing
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}

func main() {
	s := "hello"
	fmt.Printf("input: %s - answer is: %v\n", s, IsPalindrome(s))
}
