package gochallenges

import (
	"fmt"
	"sort"
)

type Flight struct {
	Origin      string
	Destination string
	Price       int
}

// Define a struct of type Flight
type ByPrice []Flight

// Overrides the sort method by implementing the three defined functions
func (p ByPrice) Len() int {
	return len(p)
}

func (p ByPrice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p ByPrice) Less(i, j int) bool {
	return p[i].Price > p[j].Price
}

func (p ByPrice) Print(data interface{}) {
	fmt.Println(data)
}

func SortByPrice(flights []Flight) []Flight {
	sort.Sort(ByPrice(flights))
	return flights
}

func C3() {
	flights := []Flight{
		{Price: 30},
		{Price: 20},
		{Price: 50},
		{Price: 1000},
	}

	sort.Sort(ByPrice(flights))
	fmt.Println(flights)
}
