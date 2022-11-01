// desigh HashMap
// what is hash tables, hash function, collision handling
// why hashtable is fast at insert/delete/search data
// for hash, when storing data, calculate data with hash function,
// and store location with hash value as index
// the data structure storing this data is called hash table
// hash function is used to calculate the index of data storage
// collision handling - open addressing, separate chaining
// open addressing, add to next after first insert, but this may lose benefit of hash table search
// separate chaining, like storing multiple names in one slot, by using linked list
//
// hash table average time complexity
// insert, search, delete O(1)
//
// Hash Table Part (array)
// struct hashtable
// method Insert()
// method Search()
// method Delete()

// Bucket Part (linked list), for separate chaining
// structure bucket
// structure bucketNode
// method insert()
// method search()
// method delete()

package main

import "fmt"

const ArraySize = 10 // hashtable size
// HashTable structure
type HashTable struct {
	array [ArraySize]*bucket
}

// bucket structure
type bucket struct {
	head *bucketNode
}

// bucketNode structure
type bucketNode struct {
	key  string
	next *bucketNode // address of next node
}

// Hashtable
// Insert
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Search
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// Delete
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// bucket structure
// insert
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println("key already exists")
	}
}

// search bucket for matching key
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// delete
func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			// delete
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

// hash function
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// Init
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}
func main() {
	testHashTable := &HashTable{}
	fmt.Println(testHashTable)

	myHashTable := Init()
	fmt.Println(myHashTable)

	testBucket := &bucket{}
	testBucket.insert("bucket")
	testBucket.insert("bucket")
	testBucket.delete("bucket")

	fmt.Println(testBucket.search("bucket"))
	fmt.Println(testBucket.search("not"))

	list := []string{
		"apple",
		"banana",
		"cherry",
		"donut",
	}

	for _, v := range list {
		testBucket.insert(v)
	}
	testBucket.delete("cherry")
	fmt.Println(testBucket.search("cherry"))
	fmt.Println(testBucket.search("apple"))
}
