/*
stack and queue are linear data structures
both can grow when adding data and shrink when removing data

stack - last in first out, lifo, push data, pop data

queue - first in first out, fifo, enqueue, dequeue
*/

package main

import "fmt"

// stack represents stack that hold a slice
type Stack struct {
	items []int
}

// Push will add a value at the end
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop will remove a value at the end
// and returns the removed value
func (s *Stack) Pop() int {
	l := len(s.items) - 1  // store last item index
	toRemove := s.items[l] // store last item
	s.items = s.items[:l]  // update items from beginning to leave one out
	return toRemove
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)

}
