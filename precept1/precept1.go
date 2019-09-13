package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	
	// Print first 10 Fibonacci Numbers
	prev := 1
	prevprev := 0
	cur := 1
	for i := 0; i < 10; i++ {
		fmt.Println(cur)
		cur = prev + prevprev
		prevprev = prev
		prev = cur
	}
	
	s := []int{1,2,3,4,5,6,7,8}
	reverse(s)
	fmt.Println(s)
	
	s2 := []int{1,1,1,2,3}
	uniq := unique(s2)
	fmt.Println(uniq)
}


func reverse(s []int) {
	l := len(s)
	for i := 0; i < l / 2; i++ {
		tmp := s[i]
		s[i] = s[l - i - 1]
		s[l - i - 1] = tmp
	}
}

func unique(s []int) int {
	count := 0
	m := make(map[int]bool)
	
	for i := 0; i < len(s); i++ {
		_, ok := m[s[i]]
		if ( !ok ) {
			m[s[i]] = true
			count++
		}
	}
	
	return count
}
