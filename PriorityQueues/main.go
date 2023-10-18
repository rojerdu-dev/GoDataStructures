package main

import "errors"

var ErrEmptyQueue = errors.New("Empty Queue")

type PriorityQueue struct {
	High []int
	Low  []int
}

// 5 Methods
// Length
// IsEmpty
// Peek
// Enqueue
// Dequeue

func (pq *PriorityQueue) Length() int {
	return len(pq.High) + len(pq.Low)
}

func (pq *PriorityQueue) IsEmpty() bool {
	return pq.Length() == 0
}

func (pq *PriorityQueue) Peek() (int, error) {
	if len(pq.High) != 0 {
		return pq.High[0], nil
	}
	if len(pq.Low) != 0 {
		return pq.Low[0], nil
	}
	return 0, ErrEmptyQueue
}

func (pq *PriorityQueue) Enqueue(n int, highPriority bool) {
	if highPriority {
		pq.High = append(pq.High, n)
	} else {
		pq.Low = append(pq.Low, n)
	}
}

func (pq *PriorityQueue) Dequeue() (int, error) {
	if len(pq.High) != 0 {
		var removedHigh int
		removedHigh, pq.High = pq.High[0], pq.High[1:]
		return removedHigh, nil
	}

	if len(pq.Low) != 0 {
		var removedLow int
		removedLow, pq.Low = pq.Low[0], pq.Low[1:]
		return removedLow, nil
	}

	return 0, ErrEmptyQueue
}
