package Arrays_and_Strings

import "fmt"

func RotateMatrix(input [][]int) bool {
	//Checking if empty matrix
	if len(input) == 0 {
		fmt.Println("Length is empty")
		return false
	}
	//Checking if it is n*n square matrix
	for key := range input {
		if len(input) != len(input[key]) {
			fmt.Println("Length is not same")
			return false
		}
	}

	//Now the magic begins
	//Lets print the matrix first
	for key := range input {
		for key1 := range input[key] {
			fmt.Print(input[key][key1], "\t")
		}
		fmt.Println("")
	}
	fmt.Println("")
	//var result [][]int
	//Step1: Transpose Matrix
	temp := 0
	for key := 0; key < len(input); key++ {
		for key1 := key; key1 < len(input); key1++ {
			if key != key1 {
				temp = input[key][key1]
				input[key][key1] = input[key1][key]
				input[key1][key] = temp
			}
		}
	}
	//Step2: Inverse the row
	for key := range input {
		for key1 := 0; key1 < len(input)/2; key1++ {
			temp = input[key][key1]
			input[key][key1] = input[key][len(input)-1-key1]
			input[key][len(input)-1-key1] = temp
		}
	}
	for key := range input {
		for key1 := range input[key] {
			fmt.Print(input[key][key1], "\t")
		}
		fmt.Println("")
	}
	fmt.Println("")

	return true
}
