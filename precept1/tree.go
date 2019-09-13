package main

import "fmt"

type Tree struct {
	Root Node
}

type Node struct {
	Left *Node
	Right *Node
	Val int 
}

func add(n *Node) int {
	if n == nil {
		return 0
	}
	return n.Val + add(n.Left) + add(n.Right)
}

func main() {
	root := Node{nil, nil, 1}
	root.Left = &Node{nil, nil, 2}
	root.Right = &Node{nil, nil, 3}
	root.Left.Left = &Node{nil, nil, 4}
	root.Left.Right = &Node{nil, nil, 5}
	root.Right.Left = &Node{nil, nil, 6}
	root.Right.Right = &Node{nil, nil, 7}
	fmt.Println(add(&root))
}



