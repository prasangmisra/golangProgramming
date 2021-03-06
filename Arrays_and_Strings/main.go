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
			FindUnique()
			break
		case "2":
			FindPermutation()
			break
		case "3":
			StringURL()
			break
		default:
			fmt.Println("Going back")
			break optionLoop
		}
	}
}
