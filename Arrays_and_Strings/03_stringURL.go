package Arrays_and_Strings

import (
	"bufio"
	"os"
	"strings"
)

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
