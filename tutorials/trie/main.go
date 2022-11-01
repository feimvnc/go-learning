// trie is a tree like data structure which stores words
// each node represents a word or part of word
// search word pretty fast, like autocomplete
// google search uses trie, sometimes it suggests next word
// each node has 26 chars, but all nil except with letters showing
// each node holds value of address of next node
// node has a boolean value indicating if node is an end of a word
// time complexity
// trie has trade of space and time
package main

import "fmt"

// define data structure first
const AlphabeSize = 26

// Node
type Node struct {
	children [26]*Node // size  26, with pointer to another node
	isEnd    bool
}

// Trie, a trie has a pointer point to root
type Trie struct {
	root *Node
}

// create new Trie
func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

// Insert, add word to Trie
func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a' // convert char index to number, 'a' = 97
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true // indicate end of word
}

// Search word, return true if found
func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}

	if currentNode.isEnd == true {
		return true
	}
	return false
}

func main() {
	testTrie := InitTrie()
	fmt.Println(testTrie)
	fmt.Println(testTrie.root)

	myTrie := InitTrie()
	myTrie.Insert("book")
	fmt.Println(myTrie.Search("booka"))

	toAdd := []string{
		"aragorn",
		"aragon",
		"argon",
		"eragon",
		"oregon",
		"oregano",
		"oreo",
	}

	for _, v := range toAdd {
		myTrie.Insert(v)
	}

	fmt.Println(myTrie.Search("oreo"))
	fmt.Println(myTrie.Search("not"))
}
