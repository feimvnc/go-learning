// linked list
// value is stored in a node
// each node has address of next node
// when change node value, we go from first node, go through until we find it
// slower O(n) when going to the Kth element
// but adding and removing values at the beginning of the list is simple
// Faster O(1) constant time
// comparing with array, add an element all element will be shifted
// doubly linked list, node contains both next and previous node address

package main

import "fmt"

type node struct {
	data int
	next *node // pointer, address of next node
}

type linkedList struct {
	head   *node // for linked list, only need to store first node
	length int
}

// func (l linkedList) prepend(n *node) {   // value receiver, change value on the copy

func (l *linkedList) prepend(n *node) { // method pointer receiver, pointer to change value
	second := l.head     // make a temp node to store current head
	l.head = n           //	assign n to first node
	l.head.next = second // update first node with second node address
	l.length++
}

// print out data of every node of the list
func (l linkedList) printListData() { // use value receiver, because of changing data, just list
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data) // print head node
		toPrint = toPrint.next          // update node to next
		l.length--
	}
	fmt.Printf("\n") // print a line break
}

// delete a node of given value
func (l *linkedList) deleteWithValue(value int) { // use pointer receiver
	if l.length == 0 { // when list is empty
		return // come out of the method when list is empty
	}

	if l.head.data == value { // if head is the node have the given value
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil { // if given value not found till end
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	myList := linkedList{}
	node1 := &node{data: 10}
	node2 := &node{data: 20}
	node3 := &node{data: 30}
	node4 := &node{data: 1}
	node5 := &node{data: 3}
	node6 := &node{data: 5}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.prepend(node6)
	fmt.Println(myList) // print myList node3 address and count of nodes

	fmt.Println(myList.head) // head node data and address of next node
	fmt.Println(myList.head.next)
	fmt.Println(myList.head.next.next) // first added node has no linked node, only data

	myList.printListData()

	myList.deleteWithValue(100)
	myList.deleteWithValue(2)

	emptyList := linkedList{}
	emptyList.deleteWithValue(10)

	myList.printListData()

}
