package main

import (
	"fmt"

	arrstr "./Arrays_and_Strings"
)

var allFunctions = [3]string{
	"1: String Unique",
	"2: String Permutation",
	"3: String URLify",
}

func main() {

	fmt.Println("Select a Code:")
	for _, value := range allFunctions {
		fmt.Println(value)
	}
	var option string
	fmt.Scanln(&option)
	switch option {
	case "1":
		arrstr.FindUnique()
		break
	case "2":
		arrstr.FindPermutation()
		break
	case "3":
		arrstr.StringURL()
		break
	default:
		fmt.Println("Nothing selected, bye")
	}
}
