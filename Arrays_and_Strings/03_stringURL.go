package Arrays_and_Strings

import (
	"bufio"
	"os"
	"strings"
)

// URLify: Write a method to replace all spaces in a string with '%20'. You may assume that the string
// has sufficient space at the end to hold the additional characters, and that you are given the "true"
// length of the string. (Note: If implementing in Java, please use a character array so that you can
// perform this operation in place.)
func StringURLInput() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		StringURL(scanner.Text())
		break
	}
	if err := scanner.Err(); err != nil {
		os.Exit(1)
	}
}
func StringURL(input string) string {
	if input == " " {
		return input
	}
	return strings.ReplaceAll(input, " ", "%20")
}
