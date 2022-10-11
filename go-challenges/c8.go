package gochallenges

import (
	"fmt"
)

func checkPermutations(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	charCount := make(map[string]int)
	for i := range str1 {
		charCount[string(str1[i])]++
		charCount[string(str2[i])]--
	}

	for _, count := range charCount {
		if count != 0 {
			return false
		}
	}

	return true
}

func CheckPermutations() {
	fmt.Println("Check Permutations Challenge")

	str1 := "adcmea"
	str2 := "medaca"

	isPermutation := checkPermutations(str1, str2)
	fmt.Println(isPermutation)
}
