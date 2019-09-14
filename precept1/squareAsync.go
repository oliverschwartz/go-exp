package main

import (
	"fmt"
	"sync"
)

func square(iter int, slice []int, wg *sync.WaitGroup) {
	fmt.Printf("Starting worker on iteration %v\n", iter)
	slice[iter] *= slice[iter]
	wg.Done()
}

/* For every element in the slice, spin up a
goRoutine to compute the square of the element
and replace it in the slice. */
func main() {
	var wg sync.WaitGroup
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 0; i < len(slice); i++ {
		wg.Add(1)
		go square(i, slice, &wg)
	}

	wg.Wait()
	fmt.Println(slice)
}
