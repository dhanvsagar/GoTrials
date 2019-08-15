package main

import "fmt"

type Stack struct {
	slice []int
}

func (s *Stack) push(data int) {
	s.slice = append(s.slice, data)
}

func (s *Stack) pop() int {
	data := s.slice[len(s.slice)-1]
	s.slice = s.slice[0 : len(s.slice)-1]
	return data
}

func main() {
	Stack1 := Stack{}
	Stack1.push(10)
	Stack1.push(20)
	Stack1.push(12)
	fmt.Println(Stack1.slice)

	fmt.Println(Stack1.pop())
	fmt.Println(Stack1.slice)
}
