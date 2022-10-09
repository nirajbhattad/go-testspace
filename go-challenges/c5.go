package gochallenges

import (
	"errors"
	"fmt"
)

type Stack struct {
	Items []Flight
}

func (s *Stack) Pop() (Flight, error) {
	if s.IsEmpty() {
		return Flight{}, errors.New("Stack is Empty")
	} else {
		topElement := len(s.Items) - 1
		flight := s.Items[topElement]
		s.Items = s.Items[:topElement]
		fmt.Println("Popped Element From Stack", flight)
		return flight, nil
	}
}

func (s *Stack) Push(f Flight) {
	s.Items = append(s.Items, f)
	fmt.Println("Pushed Element Into Stack", f)
}

func (s *Stack) Peek() (Flight, error) {
	if s.IsEmpty() {
		return Flight{}, errors.New("Stack is Empty")
	} else {
		topElement := len(s.Items) - 1
		flight := s.Items[topElement]
		fmt.Println("Element at the top of the Stack", flight)
		return flight, nil
	}
}

func (s *Stack) IsEmpty() bool {
	return len(s.Items) == 0
}

func StackImplementation() {
	fmt.Println("Go Stack Implementation")
	flights := []Flight{
		{Price: 30},
		{Price: 20},
		{Price: 50},
		{Price: 1000},
	}

	s := Stack{Items: flights}
	s.Push(Flight{Origin: "Mumbai"})
	s.Push(Flight{Destination: "Toronto"})
	fmt.Println("Stack Elements At The Moment", s.Items)
	s.Peek()
	s.Pop()
	s.Pop()
	s.Push(Flight{Origin: "Mumbai"})
	s.Push(Flight{Destination: "Toronto"})
	s.Push(Flight{Origin: "Mumbai"})
	s.Push(Flight{Destination: "Toronto"})
	fmt.Println("Stack Elements At The Moment", s.Items)
}
