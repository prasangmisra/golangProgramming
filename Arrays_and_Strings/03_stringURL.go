package Arrays_and_Strings

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StringURL() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(s)
		s = strings.ReplaceAll(s, " ", "%20")
		fmt.Println(s)
	}
	if err := scanner.Err(); err != nil {
		os.Exit(1)
	}
}
