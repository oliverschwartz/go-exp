package main

/** COS 418 Fall 2019
Professor Wyatt Lloyd, Professor Michael Freedman

Author: Jennifer Lam
**/

import "fmt"

func printClockwise(matrix [][]int) {

	top := 0              // start row index k
	bottom := len(matrix) // ending row index m
	left := 0             // starting column index l
	right := len(matrix)  // ending column index n

	for top < bottom && left < right {

		// Print first of the remaining rows
		for i := left; i < right; i++ {
			fmt.Printf("%d ", matrix[top][i])
		}
		top++

		// Print the last column from the remaining columns
		for i := top; i < bottom; i++ {
			fmt.Printf("%d ", matrix[i][right-1])
		}
		right--

		// Print the last row from the remaining rows
		if top < bottom {
			for i := right - 1; i >= left; i-- {
				fmt.Printf("%d ", matrix[bottom-1][i])
			}
			bottom--
		}

		// Print the first of the remaining columns
		if left < right {
			for i := bottom - 1; i >= top; i-- {
				fmt.Printf("%d ", matrix[i][left])
			}
			left++
		}
	}

	fmt.Println()
}

func printCounterClockwise(matrix [][]int) {

	top := 0
	bottom := len(matrix)
	left := 0
	right := len(matrix)

	for top < bottom && left < right {

		// Print first of the remaining rows
		for i := right - 1; i >= left; i-- {
			fmt.Printf("%d ", matrix[top][i])
		}
		top++

		// Print first of the remaining columns
		for i := top; i < bottom; i++ {
			fmt.Printf("%d ", matrix[i][left])
		}
		left++

		// Print the last of the remaining rows
		if top < bottom {
			for i := left; i < right; i++ {
				fmt.Printf("%d ", matrix[bottom-1][i])
			}
			bottom--
		}

		// Print the last of the remaining columns
		if left < right {
			for i := bottom - 1; i >= top; i-- {
				fmt.Printf("%d ", matrix[i][right-1])
			}
			right--
		}
	}
	fmt.Println()
}

func main() {

	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	fmt.Printf("Original matrix\n")
	for _, row := range matrix {
		fmt.Printf("%+v\n", row)
	}

	fmt.Printf("Clockwise\n")
	printClockwise(matrix)

	fmt.Printf("Counter clockwise\n")
	printCounterClockwise(matrix)
}
