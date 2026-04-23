package main

import "fmt"

func main() {
	root := NewNode(50)
	vals := []int{40, 60, 45, 65, 55, 35, 70}
	for _, val := range vals {
		root.Add(val)
	}
	root.printTree(0)
	fmt.Println("--------------")
	root.delete(50)
	root.printTree(0)
}
