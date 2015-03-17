package structure

import "testing"

type Integer int

func (i Integer) Less(j interface{}) bool {

	return int(i) < j.(int)
}

func TestRightRotation(t *testing.T) {

	a := &RedBlackTreeNode{}
	b := &RedBlackTreeNode{}
	c := &RedBlackTreeNode{}

	y := &RedBlackTreeNode{
		Right: c,
	}

	x := &RedBlackTreeNode{
		Left:   a,
		Right:  b,
		Parent: y,
	}

	y.Left = x

	err := x.RightRotate()

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if x.Left != a {
		t.Error("x subtree incorrect, left != a")
	}

	if x.Right != y {
		t.Error("x subtree incorrect, right != y")
	}

	if y.Left != b {
		t.Error("y subtree incorrect, left != b")
	}

	if y.Right != c {
		t.Error("y subtree incorrect, right != c")
	}

}

func TestLeftRotation(t *testing.T) {

	a := &RedBlackTreeNode{}
	b := &RedBlackTreeNode{}
	c := &RedBlackTreeNode{}

	y := &RedBlackTreeNode{
		Left:  b,
		Right: c,
	}

	x := &RedBlackTreeNode{
		Left:  a,
		Right: y,
	}

	y.Parent = x

	err := y.LeftRotate()

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if x.Left != a {
		t.Error("x subtree incorrect, left != a")
	}

	if x.Right != b {
		t.Error("x subtree incorrect, right != b")
	}

	if y.Left != x {
		t.Error("y subtree incorrect, left != x")
	}

	if y.Right != c {
		t.Error("y subtree incorrect, right != c")
	}

}
