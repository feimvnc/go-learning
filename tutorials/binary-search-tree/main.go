// trees has root, parent, children, leaves
// binary tree has no more than 2 children, left and right
// left children is less than parent, and right great than parent
// some binary tree has only 1 children
//
// insert - always start at root, compare, then add as leaf
// search - also start at root, similar to insert
//
// advantage is speed
// balanced binary tree - O(h), better than O(n)
// unbalanced binary tree - O(n), which traverse each node
//
// 3 main things
// Node
// Insert
// Search

package main

import "fmt"

var count int // count of node in binary tree

// Node represents the components of a binary search tree
type Node struct {
	Key   int   // compare against other node
	Left  *Node // each node may be a parent, which has left and right children
	Right *Node
}

// insert will add a node to the tree
// the key to add should not be already in the tree
func (n *Node) Insert(k int) {
	if n.Key < k { // k larger move right
		if n.Right == nil { // if right is empty, add
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k) // otherwise do same insert
		}
	} else if n.Key > k { // smaller move left
		if n.Left == nil { // if right is empty, add
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k) // otherwise do same insert
		}
	}
}

// Search will take in a key value
// and return true if there is a  node that value
func (n *Node) Search(k int) bool {
	count++ // increase count for each search
	if n == nil {
		return false
	}
	if n.Key < k {
		return n.Right.Search(k)
	} else if n.Key > k {
		return n.Left.Search(k)
	}
	return true
}

func main() {
	tree := &Node{Key: 5} // tree has to be an address
	//tree.Insert(1)        // insert at left
	tree.Insert(1) // insert at right
	tree.Insert(3)
	tree.Insert(5)
	tree.Insert(7)
	tree.Insert(9)
	fmt.Println(tree) // output &{1 <nil> <nil>}

	fmt.Println(tree.Search(2))
	fmt.Println(count) // count of node searched
}
