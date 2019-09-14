package main

/** COS 418 Fall 2019
Professor Wyatt Lloyd, Professor Michael Freedman

Author: Jennifer Lam
**/

import "fmt"

func main() {

	a := 1
	b := 1

	fmt.Printf("%d %d ", a, b)

	for i := 2; i < 10; i++ {
		fmt.Printf("%d ", a+b)
		temp := b
		b = a + b
		a = temp
	}
	fmt.Println()
}
