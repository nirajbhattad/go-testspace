package gochallenges

import "fmt"

type Employee interface {
	Language() string
	Age() int
}

type Engineer struct {
	Name string
}

func (e Engineer) Language() string {
	return e.Name + " programs in Go"
}

func (e Engineer) Age() int {
	return 25
}

func OverrideInterfaceFunction() {
	var programmers []Employee
	elliot := Engineer{Name: "Elliot"}
	programmers = append(programmers, elliot)
	fmt.Println(programmers)
}
