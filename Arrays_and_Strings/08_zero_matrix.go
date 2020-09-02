package Arrays_and_Strings

import "fmt"

func ZeroMatrix(input [][]int) [][]int {
	var rows, cols []int
	for key := range input {
		for key1, value := range input[key] {
			fmt.Print(input[key][key1])
			if value == 0 {
				if !contains(rows, key) {
					rows = append(rows, key)
				}
				if !contains(cols, key1) {
					cols = append(cols, key1)
				}
			}
		}
		fmt.Println()
	}
	for key := range rows {
		for key1 := range input[key] {
			input[rows[key]][key1] = 0
		}
	}

	for key := range cols {
		for key1 := range input {
			input[key1][cols[key]] = 0
		}
	}
	fmt.Println()
	for key := range input {
		for _, value := range input[key] {
			fmt.Print(value)
		}
		fmt.Println()
	}
	return input
}

func contains(array []int, key int) bool {
	for _, value := range array {
		if value == key {
			return true
		}
	}
	return false
}
