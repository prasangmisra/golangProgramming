package Arrays_and_Strings

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PalindromePermutation() {
	fmt.Println("\nEnter a string to check if its a permutation of palindrome")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := strings.ToLower(scanner.Text())
		fmt.Println(s)
		m := make(map[rune]int)
		for _, value := range s {
			_, prs := m[value]
			if prs {
				m[value] = m[value] + 1
			} else {
				m[value] = 1
			}
		}
		fmt.Println(m)
		odds := 0
		for key, value := range m {
			fmt.Println(key, " -- ", value)
			if value%2 != 0 {
				odds++
			}
		}
		if odds > 1 {
			fmt.Println("No Palindrome")
		} else {
			fmt.Println("Palindrome")
		}
		break
	}
	if err := scanner.Err(); err != nil {
		os.Exit(1)
	}
}
