package main

import (
	"fmt"

	arrstr "github.com/prasangmisra/golangProgramming/Arrays_and_Strings"
)

var allFunctions = [1]string{
	"1: Arrays and Strings",
}

func main() {
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
			arrstr.GetOptionsList()
			break
		default:
			fmt.Println("Nothing selected, bye")
			break optionLoop
		}
	}
}
