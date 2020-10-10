package algorithms

var (
	keys  []int
	count int = 0
)

// Node represents a leaf component in the binary tree
type Node struct {
	Key        int
	LeftChild  *Node
	RightChild *Node
}

// Insert adds a new node to the tree
func (node *Node) Insert(key int) {
	if key > node.Key {
		if node.RightChild == nil {
			node.RightChild = &Node{Key: key}
			count++
		} else {
			node.RightChild.Insert(key)
		}
	} else if key < node.Key {
		if node.LeftChild == nil {
			node.LeftChild = &Node{Key: key}
			count++
		} else {
			node.LeftChild.Insert(key)
		}
	}
}

// Search finds a key in the tree
func (node Node) Search(key int) bool {
	var status bool

	if node.Key == key {
		status = true
	} else if node.Key > key && node.LeftChild != nil {
		return node.LeftChild.Search(key)
	} else if node.Key < key && node.RightChild != nil {
		return node.RightChild.Search(key)
	} else {
		status = false
	}

	return status
}

// InOrderTraversal returns keys in a inorder sequence
func (node Node) InOrderTraversal() []int {
	if node.LeftChild != nil {
		node.LeftChild.InOrderTraversal()
	}

	keys = append(keys, node.Key)

	if node.RightChild != nil {
		node.RightChild.InOrderTraversal()
	}

	return keys
}
