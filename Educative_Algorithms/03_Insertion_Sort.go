package Educative_Algorithms

import (
	"fmt"
)

func InsertionSort(input []int) []int {
	var n = len(input)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if input[j-1] > input[j] {
				input[j-1], input[j] = input[j], input[j-1]
			}
			j = j - 1
		}
	}
	return input
}

func InsertToArray(input []int, index, value int) []int {
	fmt.Println(index, value)
	result := make([]int, 0)
	temp := make([]int, 0)
	//First of all remove the element
	for key := range input {
		if input[key] != value {
			temp = append(temp, input[key])
		}
	}
	input = temp

	for key := index; key >= 0; key-- {
		fmt.Println("Comparing ", input[key])
		if value < input[key] {
			//insert the value
			fmt.Println("Smaller,go left", key)
		} else {
			fmt.Println("Found the place", key)
			fmt.Println(input[:key+1])
			result = append(result, input[:key+1]...)
			result = append(result, value)
			result = append(result, input[key+1:]...)
			break
		}
		if key == 0 {
			result = append(result, value)
			result = append(result, input[key:]...)
			fmt.Println("Gotcha")
		}
		//fmt.Println(result)
	}
	return result
}
