package tree

type node struct {
	value int
	left  *node
	right *node
}

type BinaryTree struct {
	root *node
}

func (t *BinaryTree) Insert(val int) *BinaryTree {
	if t.root == nil {
		t.root = &node{value: val, left: nil, right: nil}
	} else {
		t.root.Insert(val)
	}

	return t
}

func (n *node) Insert(val int) {
	if n == nil {
		return
	}

	if val <= n.value {
		if n.left == nil {
			n.left = &node{value: val, left: nil, right: nil}
		} else {
			n.left.Insert(val)
		}
	} else {
		if n.right == nil {
			n.right = &node{value: val, left: nil, right: nil}
		} else {
			n.right.Insert(val)
		}
	}
}
