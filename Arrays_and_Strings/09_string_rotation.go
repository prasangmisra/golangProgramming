package Arrays_and_Strings

import "fmt"

func IsSubstring(first, second string) bool {
	if len(first) != len(second) {
		return false
	}
	if len(first) < 2 || len(second) < 2 {
		return true
	}

	for i := 0; i < len(first); i++ {
		if first == second {
			return true
		}
		second = RotateStringBy1Bit(second)
	}
	return false
}

func RotateStringBy1Bit(input string) string {
	input = input[1:] + input[:1]
	fmt.Println(input)
	return input
}