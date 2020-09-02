package Arrays_and_Strings

import (
	"fmt"
	"strconv"
)

// Implement a method to perform basic string compression using the counts
// of repeated characters. For example, the string aabcccccaaa would become a2blc5a3. If the
// "compressed "string would not become smaller than the original string, your method should return
// the original string. You can assume the string has only uppercase and lowercase letters (a - z),

func StringCompression(input string) string {
	if input == "" {
		return input
	}
	compressedString := string(input[0])
	fmt.Println(compressedString)
	currentString := input[0]
	currentCount := 1
	for key := range input {
		fmt.Println(key, "__", input[key])
		if key != 0 {
			if input[key] == currentString {
				fmt.Println("Same character found, incrementing count")
				currentCount++
				fmt.Println(currentCount)
			} else {
				fmt.Println("Different char found, updating compressedString")
				compressedString += strconv.Itoa(currentCount) + string(input[key])
				fmt.Println(compressedString)
				currentString = input[key]
				currentCount = 1
			}

		}
	}
	compressedString += strconv.Itoa(currentCount)
	fmt.Println(compressedString)
	if len(compressedString) >= len(input) {
		return input
	}
	return compressedString
}
