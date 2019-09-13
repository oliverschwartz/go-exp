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

func main() {
	root := Node{nil, nil, 1}
	t := Tree{root}
	root.Left = &Node{nil, nil, 2}
	root.Right = &Node{nil, nil, 3}
}

