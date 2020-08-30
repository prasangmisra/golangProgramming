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
		s = strings.ReplaceAll(s, " ", "%20")
		fmt.Println(s)
		break
	}
	if err := scanner.Err(); err != nil {
		os.Exit(1)
	}
}
