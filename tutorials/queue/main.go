// queue and stack inserting data is same,
// removing data is different

package main

import "fmt"

// Queue represents a queue that holds a slice
type Queue struct {
	items []int
}

// Enqueue
func (q *Queue) Enqueue(i int) { // pointer receiver, no return value
	q.items = append(q.items, i)
}

// Dequeue
func (q *Queue) Dequeue() int { // no input parameter, return removed value
	toRemove := q.items[0] // first item in queue
	q.items = q.items[1:]  // from index 1 to end
	return toRemove
}

func main() {
	myQueue := Queue{}
	fmt.Println(myQueue)
	myQueue.Enqueue(2)
	myQueue.Enqueue(4)
	myQueue.Enqueue(6)
	fmt.Println(myQueue)

	myQueue.Dequeue()
	fmt.Println(myQueue)
}
