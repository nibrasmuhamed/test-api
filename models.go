package main

import (
	"fmt"
	"time"
)

type Values struct {
	sum     float32
	average float32
	max     float32
	min     float32
	count   float32
}

type Node struct {
	data    float32
	updated time.Time
	next    *Node
}

type Queue struct {
	front *Node
	rear  *Node
}

func (q *Queue) IsEmpty() bool {
	if q.rear == nil && q.front == nil {
		return true
	}
	return false
}

func (q *Queue) Enqueue(data float32, update time.Time) {
	newNode := &Node{data, update, nil}
	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
		return
	}
	q.rear.next = newNode
	q.rear = newNode
}

func (q *Queue) Dequeue() {
	if q.front == nil {
		fmt.Println("queue empty")
		return
	}
	q.front = q.front.next
}
