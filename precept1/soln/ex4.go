package main

/** COS 418 Fall 2019
Professor Wyatt Lloyd, Professor Michael Freedman

Author: Jennifer Lam
**/

import "fmt"

func reverse(s []int) {

	for i := 0; i < len(s)/2; i++ {
		temp := s[i]
		s[i] = s[len(s)-1-i]
		s[len(s)-1-i] = temp
	}
}

func main() {

	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("unreversed %+v\n", s)

	reverse(s)
	fmt.Printf("reversed: %+v\n", s)
}
