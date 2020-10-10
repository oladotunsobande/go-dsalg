package algorithms

import "fmt"

// RunBst executes the BST algorithm
func RunBst() {
	binaryTree := &Node{Key: 100}
	binaryTree.Insert(250)
	binaryTree.Insert(45)
	binaryTree.Insert(432)
	binaryTree.Insert(21)
	binaryTree.Insert(96)
	binaryTree.Insert(309)
	fmt.Println("tree: ", binaryTree)

	status := binaryTree.Search(309)
	fmt.Println("search for 309 in tree: ", status)

	ioTraversal := binaryTree.InOrderTraversal()
	fmt.Println("InOrder Traversal: ", ioTraversal)
}

// RunHashTable executes the hash table algorithm
func RunHashTable() {
	vals := [6]string{"bolu", "glory", "damola", "blessing", "mayokun", "taiwo"}

	newHashTable := Init(len(vals))

	for i := 0; i < len(vals); i++ {
		newHashTable.Insert(vals[i])
	}

	index, status := newHashTable.Search("tayo")
	fmt.Println("value: 'tayo', index: ", index, ", status: ", status)
}
