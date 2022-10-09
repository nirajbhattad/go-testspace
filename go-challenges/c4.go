package gochallenges

import "fmt"

func FilterUnique(developers []Developer) []string {
	var uniqueSet []string
	unique := make(map[string]bool)

	for _, developer := range developers {
		if _, ok := unique[developer.Name]; ok {
			continue
		} else {
			unique[developer.Name] = true
			uniqueSet = append(uniqueSet, developer.Name)
		}
	}

	// TODO Implement
	return uniqueSet
}

func FindUnique() {
	developers := []Developer{
		{Name: "Elliot"},
		{Name: "Alan"},
		{Name: "Jennifer"},
		{Name: "Graham"},
		{Name: "Paul"},
		{Name: "Alan"},
	}

	unique := FilterUnique(developers)
	fmt.Println(unique)
}
