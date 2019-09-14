package main

import "fmt"

// Print the matrix in spiral order, clockwise
func printSpiralClkwise(matrix [][]int) {
	if len(matrix) != len(matrix[0]) {
		fmt.Println("matrix must be square. Aborting...")
		return
	}

	n := len(matrix)
	top := 0
	bottom := n
	left := 0
	right := n

	for top < bottom && left < right {
		// print first of remaining rows
		for i := left; i < right; i++ {
			fmt.Println(matrix[top][i])
		}
		top++

		// print last of remaining cols
		for i := top; i < bottom; i++ {
			fmt.Println(matrix[i][right-1])
		}
		right--

		// print last of remaining rows
		for i := right; i > left; i-- {
			fmt.Println(matrix[bottom-1][i])
		}
		bottom--

		// print first of remaining rows
		for i := bottom; i > top; i-- {
			fmt.Println(matrix[i][left])
		}
		left++
	}
}

func main() {
	size := 5
	matrix := make([][]int, size)

	for i := 0; i < size; i++ {
		matrix[i] = make([]int, size)
		for j := 0; j < size; j++ {
			matrix[i][j] = i*size + j
		}
	}

	fmt.Println(matrix)
	printSpiralClkwise(matrix)
}
