package Arrays_and_Strings

import (
	"fmt"
	"sort"
	"strings"
)

func FindPermutation() bool {
	result := true
	fmt.Println("Enter 2 strings to check if we have a permutation")
	var firstString, secondString string
	fmt.Scanln(&firstString)
	fmt.Scanln(&secondString)
	firstString = SortString(firstString)
	secondString = SortString(secondString)
	fmt.Println(firstString)
	fmt.Println(secondString)
	if len(firstString) == len(secondString) {
		for key := range firstString {
			if firstString[key] != secondString[key] {
				result = false
				break
			}
		}
	} else {
		result = false
	}
	fmt.Println(result)
	return result
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
