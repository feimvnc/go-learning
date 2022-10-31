/*
heap,

normal queues, first in first out
priority queues, highest priority served first
good to use for priority queue (abstract data type)

heap can be described as a complete tree
which means all tree levels are full, except at leaf level
lagest key stored at root
getting the largest key is very fast, because we know the root of max heap
for min heap, the lowest key is at root

heap can be seen as a tree, but behind the scene,
it is stored as an array, each node of the tree corresponds to an index of that array
heap is an array but operates like a tree
you can easily calculate the indices of children based on
the parent index, and vice versa

parent index 		left child index
		_ x 2 + 1 = _

parent index 		right child index
		_ x 2 + 2 = _

max heap insert, add at botttom right
then re-arrange nodes to maintain heap property
which keeps parent key larger than its children
we compare new node to its parent node, and swap if new node is higher
we keep this process until it completes
this process is called heapify

To extract a node from tree, it is also called heapify to re-arrange nodes
after remove largest node, fill root with last node
then look at root and its left and right node, which one is the largest
swap  nodes if child is larger than root
go through this process untile it completes

time complexity is O(h), the height of the tree

for insert or extract, the time complexity is O(log n)

coding needs
struct for maxheap
insert method
extract method


// accessible from outside
struct MaxHeap
method Insert
method Extract

//
method maxHeapifyUp
method maxHeapifyDown
method swap

//
function parent
function left
function right

*/

package main

import "fmt"

// MaxHeap struct has a slice that holds the array
type MaxHeap struct {
	array []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract returns the largest key or the root
// and remove it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	l := len(h.array) - 1

	// when array is empty, return
	if len(h.array) == 0 {
		fmt.Println("cannot extract because array length is 0")
	}
	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.maxHeapifyDown(0)
	return extracted
}

// maxHeapifyUp will heapify from bottom top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// maxHeapifyDown will heapify top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	// loop while index has at least one child
	for l <= lastIndex {
		if l == lastIndex { // when left child is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // when left child is larger
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}
	}
	// compare array value of current index to larger child and swap if smaller
	if h.array[index] < h.array[childToCompare] {
		h.swap(index, childToCompare)
		index = childToCompare
		l, r = left(index), right(index)
	} else {
		return
	}
}

// supporting function
// formula is from parent_index x 2 + 1 = left_child_index
// left child is odd number, right child is even number
func parent(i int) int {
	return (i - 1) / 2
}

// get the left child index
func left(i int) int {
	return 2*i + 1
}

// get the right child index
// from formula parent_index * 2 + 2 = right_child_index
func right(i int) int {
	return 2*i + 2
}

// swap keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

// go language specifications
// a, b = b, a // exchange a and b
/*
x := []int{1,2,3}
i := 0
i, x[i] = 1, 2 // set i = 1, x[0] = 2
*/

func main() {
	m := &MaxHeap{}
	fmt.Println(m) // output = &{[]}

	buildHeap := []int{10, 20, 30, 1, 3, 5, 7, 9}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
