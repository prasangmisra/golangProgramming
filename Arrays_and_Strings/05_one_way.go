package Arrays_and_Strings

import (
	"fmt"
	"strings"
)

func OneWay(first, second string) bool {
	result := false
	//Check if they are one replacement
	switch len(first) - len(second) {
	case 0:
		result = checkReplace(first, second)
		break
	case 1:
		result = checkInsertRemove(first, second)
		break
	case -1:
		result = checkInsertRemove(second, first)
	default:
		break
	}
	return result
}

func checkInsertRemove(big, small string) bool {
	if strings.Contains(big, small) {
		fmt.Println("Small contains big")
		return true
	}
	fmt.Println("Small does not contain big")
	return false
}
func checkReplace(first, second string) bool {
	fmt.Println("Inside checkReplace")
	if first == second {
		return true
	}
	count := 0
	for key := range first {
		if first[key] != second[key] {
			count++
			if count > 1 {
				return false
			}
		}
	}
	return true
}
