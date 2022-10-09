package gochallenges

import (
	"errors"
	"fmt"
)

type Queue struct {
	Items []Flight
}

func (q *Queue) Pop() (Flight, error) {
	if q.IsEmpty() {
		return Flight{}, errors.New("Queue is Empty")
	} else {
		topElement := 0
		flight := q.Items[topElement]
		q.Items = q.Items[1:]
		fmt.Println("Popped Element From Queue", flight)
		return flight, nil
	}
}

func (q *Queue) Push(f Flight) {
	q.Items = append(q.Items, f)
	fmt.Println("Pushed Element Into Queue", f)
}

func (q *Queue) Peek() (Flight, error) {
	if q.IsEmpty() {
		return Flight{}, errors.New("Queue is Empty")
	} else {
		flight := q.Items[0]
		fmt.Println("Element at the top of the Queue", flight)
		return flight, nil
	}
}

func (q *Queue) IsEmpty() bool {
	return len(q.Items) == 0
}

func QueueImplementation() {
	fmt.Println("Go Queue Implementation")
	flights := []Flight{
		{Price: 30},
		{Price: 20},
		{Price: 50},
		{Price: 1000},
	}
	q := Queue{Items: flights}
	q.Push(Flight{Origin: "Mumbai"})
	q.Push(Flight{Destination: "Toronto"})
	fmt.Println("Queue Elements At The Moment", q.Items)
	q.Peek()
	q.Pop()
	q.Pop()
	q.Push(Flight{Origin: "Mumbai"})
	q.Push(Flight{Destination: "Toronto"})
	q.Push(Flight{Origin: "Mumbai"})
	q.Push(Flight{Destination: "Toronto"})
	fmt.Println("Queue Elements At The Moment", q.Items)
}
