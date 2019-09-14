package main

import "fmt"

func sort(slice []int) {
	if len(slice) > 1 {
		mid := len(slice) / 2

		left := make([]int, len(slice[:mid]))
		right := make([]int, len(slice[mid:]))
		copy(left, slice[:mid])
		copy(right, slice[mid:])

		sort(left)
		sort(right)

		// merge the left and right halves
		i, j, k := 0, 0, 0
		for i < len(left) && j < len(right) {
			if left[i] < right[j] {
				slice[k] = left[i]
				i++
			} else {
				slice[k] = right[j]
				j++
			}
			k++
		}

		for i < len(left) {
			slice[k] = left[i]
			k++
			i++
		}

		for j < len(right) {
			slice[k] = right[j]
			k++
			j++
		}
	}
}

func main() {
	slice := []int{9, 8, 7, 4, 1, 2, 3, 5, 6, 0}
	sort(slice)
	fmt.Println(slice)
}
