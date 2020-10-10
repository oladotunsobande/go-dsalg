package algorithms

import "fmt"

var arraySize int

// HashTable defines the type for the hash table array
type HashTable struct {
	Array []*bucket
}

// bucket defines the linkedlist as an item in the hash table
type bucket struct {
	Head *bucketNode
}

// bucketNode defines each node in the linkedlist
type bucketNode struct {
	Key  string
	Next *bucketNode
}

// Insert adds a new value to the hash table
func (h *HashTable) Insert(val string) {
	index := hash(val)
	arrayBucket := h.Array[index]
	arrayBucket.bucketInsert(val)
	fmt.Printf("value: %s, index: %d\n", val, index)
}

// Search finds a value in the hash table
func (h *HashTable) Search(val string) (int, bool) {
	index := hash(val)
	arrayBucket := h.Array[index]
	status := arrayBucket.bucketSearch(val)

	return index, status
}

// bucketInsert inserts a bucket node
func (b *bucket) bucketInsert(value string) {
	if b.Head == nil {
		b.Head = &bucketNode{Key: value}
	} else {
		temp := b.Head

		b.Head = &bucketNode{Key: value, Next: temp}
	}
}

// bucketSearch finds a key in the linked list
func (b *bucket) bucketSearch(value string) bool {
	var status bool = false

	currentNode := b.Head

	if currentNode.Key == value {
		status = true
	} else {
		if currentNode.Next != nil {
			currentNode = currentNode.Next

			for currentNode.Next != nil {
				if currentNode.Key == value {
					status = true
					break
				} else {
					currentNode = currentNode.Next
				}
			}
		}
	}

	return status
}

// Init sets the array size for the hash table
func Init(size int) *HashTable {
	arraySize = size

	var array []*bucket

	for i := 1; i <= size; i++ {
		array = append(array, &bucket{})
	}

	var result = &HashTable{Array: array}

	return result
}

// hash is the hashing function for generating indexes
func hash(key string) int {
	sum := 0

	for _, v := range key {
		sum += int(v)
	}

	return sum % arraySize
}
