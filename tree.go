package main

import "fmt"

type node struct {
	value  int
	height int
	left   *node
	right  *node
}

func NewNode(val int) *node {
	return &node{
		value:  val,
		height: 0,
	}
}

func (n *node) Add(val int) {
	if val <= n.value {
		if n.left != nil {
			n.left.Add(val)
		} else {
			n.left = NewNode(val)
		}
	} else {
		if n.right != nil {
			n.right.Add(val)
		} else {
			n.right = NewNode(val)
		}
	}

	n.updateHeight()
}

func (n *node) updateHeight() {
	leftHeight := -1
	rightHeight := -1

	if n.left != nil {
		leftHeight = n.left.height
	}
	if n.right != nil {
		rightHeight = n.right.height
	}

	n.height = 1 + max(leftHeight, rightHeight)
}

func (n *node) traverse() {
	if n.left != nil {
		n.left.traverse()
	}
	fmt.Println(n.value)
	if n.right != nil {
		n.right.traverse()
	}
}

func getSuccesor(root *node) *node {
	var curr *node = root.right
	for curr != nil && curr.left != nil {
		curr = curr.left
	}
	return curr
}

func getPredecessor(root *node) *node {
	var curr *node = root.left
	for curr != nil && curr.right != nil {
		curr = curr.right
	}
	return curr
}

func (n *node) delete(val int) *node {
	if val < n.value {
		if n.left != nil {
			n.left = n.left.delete(val)
		}

		n.updateHeight()
		return n
	} else if val > n.value {
		if n.right != nil {
			n.right = n.right.delete(val)
		}

		n.updateHeight()
		return n
	}

	// both children not exist
	if n.left == nil && n.right == nil {
		return nil
	}

	// one child exists
	if n.left != nil && n.right == nil {
		return n.left
	}
	if n.right != nil && n.left == nil {
		return n.right
	}

	// both children exists
	replacementNode := getSuccesor(n)
	n.value = replacementNode.value
	n.right = n.right.delete(replacementNode.value)
	n.updateHeight()

	return n
}

func (n *node) toArray(v *[]int) {
	if n.left != nil {
		n.left.toArray(v)
	}
	*v = append(*v, n.value)
	if n.right != nil {
		n.right.toArray(v)
	}
}

func (n *node) printTree(level int) {
	if n.right != nil {
		n.right.printTree(level + 1)
	}

	for range level {
		fmt.Printf("\t")
	}

	fmt.Println(n.value)

	if n.left != nil {
		n.left.printTree(level + 1)
	}
}
