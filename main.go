package main

import "fmt"

func main() {
	root := NewNode(50)
	vals := []int{40, 30, 25, 35, 45, 60, 65, 64, 70, 55, 54, 56}
	for _, val := range vals {
		root.Add(val)
	}
	root.traverse()
	fmt.Println("----------------")
	root.printTree(0)

	root = root.delete(50)
	fmt.Println("----------------")
	root.traverse()
	fmt.Println("----------------")

	root.printTree(0)
}
