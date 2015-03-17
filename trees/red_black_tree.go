package structure

// We can only elements of a poset
type Orderer interface {
	Less(interface{}) bool
}

// Elements of the tree
type RedBlackTreeNode struct {
	Red   bool
	Value Orderer

	// Pointers to branches
	Left   *RedBlackTreeNode
	Right  *RedBlackTreeNode
	Parent *RedBlackTreeNode
}

// A general tree type
type RedBlackTree struct {
	Root *RedBlackTreeNode
}
