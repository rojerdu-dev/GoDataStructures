package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	Elements []string
}

var ErrEmptyQueue = errors.New("Empty Queue")

func (q *Queue) Length() int {
	return len(q.Elements)
}

func (q *Queue) IsEmpty() bool {
	return q.Length() == 0
}

func (q *Queue) Enqueue(item string) {
	q.Elements = append(q.Elements, item)
}

func (q *Queue) Dequeue(item string) (string, error) {
	if len(q.Elements) == 0 {
		return "", ErrEmptyQueue
	}
	var firstElement string
	firstElement, q.Elements = q.Elements[0], q.Elements[1:]

	return firstElement, nil
}

func (q *Queue) Peek() (string, error) {
	if q.IsEmpty() {
		return "", ErrEmptyQueue
	}
	return q.Elements[0], nil
}

func main() {
	fmt.Println("Implementing Queues in Golang")
}
