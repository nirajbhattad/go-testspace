package gochallenges

import (
	"fmt"
)

func GetMinMax(flights []Flight) (int, int, error) {
	var min, max int
	min, max = 100000, 0
	for _, flight := range flights {
		if flight.Price < min {
			min = flight.Price
		}

		if flight.Price > max {
			max = flight.Price
		}
	}

	return min, max, nil
}

func MinMax() {
	flights := []Flight{
		{Price: 30},
		{Price: 20},
		{Price: 50},
		{Price: 1000},
	}
	min, max, err := GetMinMax(flights)
	fmt.Println("Getting the Minimum and Maximum Flight Prices", min, max, err)
}
