package main

import "fmt"

func square(slice []int) {
	ch := make(chan int, len(slice))

	for i := 0; i < len(slice); i++ {
		go func(index int) {
			ch <- index
			slice[index] *= slice[index]
		}(i)
	}

	for i := 0; i < len(slice); i++ {
		<-ch
	}
}

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	square(slice)
	fmt.Println(slice)
}
