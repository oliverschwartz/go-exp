package main

/** COS 418 Fall 2019
Professor Wyatt Lloyd, Professor Michael Freedman

Author: Jennifer Lam
**/

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("fizzbuzz ")
		} else if i%3 == 0 {
			fmt.Printf("fizz ")
		} else if i%5 == 0 {
			fmt.Printf("buzz ")
		} else {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
