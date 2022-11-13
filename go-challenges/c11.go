package gochallenges

import (
	"fmt"

	"github.com/hbollon/go-edlib"
)

// ClosestMatch
func ClosestMatch() {

	res, err := edlib.StringsSimilarity("string1", "string2", edlib.Levenshtein)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Similarity: %f", res)
	}

	strList := []string{"Batman", "Joker", "Batmn trilogy", "ssn2"}
	res1, err := edlib.FuzzySearchThreshold("password", strList, 0.7, edlib.Levenshtein)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Result for 'testnig': %s", res1)
	}

}
