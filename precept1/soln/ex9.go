package main

/** COS 418 Fall 2019
Professor Wyatt Lloyd, Professor Michael Freedman

Author: Jennifer Lam
**/

import "fmt"

func mergesort(s []int) []int {

	c := make([]int, len(s))
	copy(c, s)

	if len(s) <= 1 {
		return c
	}

	var firstHalf, secondHalf []int
	ch := make(chan bool, 2)
	go func() {
		firstHalf = mergesort(c[:len(s)/2])
		ch <- true
	}()

	go func() {
		secondHalf = mergesort(c[len(s)/2:])
		ch <- true
	}()

	// block, waiting for mergesort to finish
	for i := 0; i < 2; i++ {
		<-ch
	}

	i, j, k := 0, 0, 0
	for i < len(firstHalf) && j < len(secondHalf) {
		if firstHalf[i] < secondHalf[j] {
			c[k] = firstHalf[i]
			i++
		} else {
			c[k] = secondHalf[j]
			j++
		}
		k++
	}

	for _, val := range firstHalf[i:] {
		c[k] = val
		k++
	}
	for _, val := range secondHalf[j:] {
		c[k] = val
		k++
	}

	return c
}

func main() {

	s := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}

	fmt.Printf("Original %+v\n", s)

	s = mergesort(s)
	fmt.Printf("Sorted %+v\n", s)
}
