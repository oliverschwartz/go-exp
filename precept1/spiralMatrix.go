package main

import "fmt"

func main() {
	const size = 5
	var matrix [size][size]int

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			matrix[i][j] = i*size + j
		}
	}

	fmt.Println(matrix)
}
