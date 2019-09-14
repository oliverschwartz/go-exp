package main

/** COS 418 Fall 2019
Professor Wyatt Lloyd, Professor Michael Freedman

Author: Jennifer Lam
**/

import "fmt"

func distinct(s []int) int {
	count := 0
	distMap := make(map[int]int)
	for _, v := range s {
		if _, ok := distMap[v]; !ok {
			count = count + 1
			distMap[v] = 0 // placeholder value 0
		}
	}
	return count
}

func main() {

	s := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
	fmt.Printf("distinct: %d\n", distinct(s))
}
