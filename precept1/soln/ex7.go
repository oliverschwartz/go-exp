package main

/** COS 418 Fall 2019
Professor Wyatt Lloyd, Professor Michael Freedman

Author: Jennifer Lam
**/

import "fmt"

func squareInParallel(s []int) {

	ch := make(chan int, len(s))
	for i, val := range s {

		// fmt.Printf("%d, %d\n", i, val)
		go func(index int, value int) {
			s[index] = value * value
			// fmt.Printf("go i=%d, val=%d\n", index, value)
			ch <- index
		}(i, val)
	}

	for i := 0; i < len(s); i++ {
		<-ch
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("original: %+v\n", s)

	squareInParallel(s)
	fmt.Printf("squared: %+v\n", s)
}
