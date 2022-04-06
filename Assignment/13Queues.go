package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	items []int
}

//push
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)

}

//pop

func (q *Queue) Dequeue() (int, error) {
	if len(q.items) == 0 {
		return 0, errors.New("empty queue")
	}
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove, nil

}

//size
func (q *Queue) Size() int {
	return len(q.items)
}

//Isempty
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

//Peek
func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("empty queue")
	}
	return q.items[0], nil
}

func main() {
	myQueue := Queue{}
	fmt.Println(myQueue)
	myQueue.Enqueue(1)
	myQueue.Enqueue(2)
	myQueue.Enqueue(3)
	myQueue.Enqueue(4)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	myQueue.Dequeue()
	peek, _ := myQueue.Peek()
	fmt.Println(myQueue)
	fmt.Println("peek element:", peek)
	fmt.Println("size of queue: ", myQueue.Size())
	myQueue.Dequeue()
	myQueue.Dequeue()
	fmt.Println(myQueue.Peek())
	fmt.Println(myQueue)
	ele, _ := myQueue.Dequeue()
	fmt.Println(ele)
	fmt.Println(myQueue.Dequeue())
	fmt.Println(myQueue.IsEmpty())

}
