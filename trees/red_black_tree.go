package structure

import "errors"

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

func (x *RedBlackTreeNode) RightRotate() error {

	var err error

	y := x.Parent

	b := x.Right

	if y.Parent != nil {
		switch {
		case y == y.Parent.Left:
			y.Parent.Left = x
		case y == y.Parent.Right:
			y.Parent.Right = x
		default:
			err = errors.New("Invalid decendent field")
			return err
		}
	}

	x.Right = y
	y.Left = b

	return nil

}

func (y *RedBlackTreeNode) LeftRotate() error {

	var err error

	x := y.Parent
	b := y.Left

	if x.Parent != nil {
		switch {
		case x == x.Parent.Left:
			x.Parent.Left = y
		case x == x.Parent.Right:
			x.Parent.Right = y
		default:
			err = errors.New("Invalid decendent field")
			return err
		}
	}

	x.Right = b
	y.Left = x

	return nil

}
