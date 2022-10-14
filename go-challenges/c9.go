package gochallenges

import "fmt"

type Node struct {
	Value string
	Next  *Node
}

type SinglyLinkedList struct {
	Count int
	Head  *Node
	Tail  *Node
}

// Yet to be done
func (l *SinglyLinkedList) Append(value string) {
	if l.Size() == 0 {
		node := &Node{Value: value, Next: nil}
		l.Count = 1
		l.Head = node
		l.Tail = node
	} else {
		l.Count++
	}

}

func (l *SinglyLinkedList) Size() int {
	return l.Count
}

func (l *SinglyLinkedList) Print() {
	current := l.Head
	fmt.Printf("%+v\n", current.Value)
	for current.Next != nil {
		fmt.Printf("%+v\n", current.Value)
		current = current.Next
	}
}

func SingleLinkedList() {
	fmt.Println("Singly Linked List Challenge")

	var singleLinkedList SinglyLinkedList

	values := []string{"First", "Second", "Third"}
	for _, value := range values {
		singleLinkedList.Append(value)
	}
	singleLinkedList.Print()
}
