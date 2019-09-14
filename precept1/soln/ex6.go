package main

/** COS 418 Fall 2019
Professor Wyatt Lloyd, Professor Michael Freedman

Author: Jennifer Lam
**/

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (root *Node) sum() int {

	result := root.Val

	if root.Left != nil {
		result += root.Left.sum()
	}

	if root.Right != nil {
		result += root.Right.sum()
	}

	return result
}

func main() {

	one := Node{1, nil, nil}
	three := Node{3, nil, nil}
	two := Node{2, &one, &three}

	five := Node{5, nil, nil}
	seven := Node{7, nil, nil}
	six := Node{6, &five, &seven}

	four := Node{4, &two, &six}

	fmt.Printf("%d\n", four.sum())

}
