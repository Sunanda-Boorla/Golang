package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	items []int
}

//push
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)

}

//peek
func (s *Stack) Peek() int {
	return s.items[len(s.items)-1]
}

//pop

func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return 0, errors.New("no more elements to pop")
	}
	l := len(s.items) - 1
	toPop := s.items[l]
	s.items = s.items[:l]
	return toPop, nil

}

//size
func (s *Stack) Size() int {
	return len(s.items)
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)
	fmt.Println(myStack)
	fmt.Println(myStack.Pop())
	fmt.Println(myStack)
	fmt.Println("Peek element: ", myStack.Peek())
	fmt.Println("size of Stack:", myStack.Size())
	myStack.Pop()
	myStack.Pop()
	myStack.Pop()
	myStack.Pop()
	fmt.Println(myStack.Pop())
	fmt.Println(myStack)

}
