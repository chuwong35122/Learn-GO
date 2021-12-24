package main

import "fmt"

/*
	struct is collection of fields
*/

// Define type of Stack using struct
type stack struct {
	index int
	data  [5]int
}

func (s *stack) push(k int) {
	s.data[s.index] = k
	s.index++
}

func (s *stack) pop() int {
	s.index--
	return s.data[s.index]
}

func main() {
	s := new(stack)
	s.push(23)
	s.push(14)
	fmt.Println("stack ", *s)
}
