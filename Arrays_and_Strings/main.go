package Arrays_and_Strings

import (
	"fmt"
)

var allFunctions = [4]string{
	"1: Unique String",
	"2: String Permutation",
	"3: String URLify",
	"4: Palindrome Permutation",
}

func GetOptionsList() {
optionLoop:
	for {
		fmt.Println("\n\nSelect a Code:")
		for _, value := range allFunctions {
			fmt.Println(value)
		}
		var option string
		fmt.Scanln(&option)
		switch option {
		case "1":
			FindUniqueInput()
			break
		case "2":
			FindPermutationInput()
			break
		case "3":
			StringURLInput()
			break
		case "4":
			PalindromePermutationInput()
			break
		default:
			fmt.Println("Going back")
			break optionLoop
		}
	}
}
