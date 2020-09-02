package Educative_Algorithms

import (
	"fmt"
)

func SelectionSort(input []int) []int {
	//Iterate over the array, select the lowest in every iteration
	lowest := 0
	if len(input) == 0 || len(input) == 1 {
		return input
	}
	for key := 0; key < len(input); key++ {
		lowest = key
		fmt.Println("\nLowest is", lowest)
		fmt.Println(input)
		fmt.Println("Pivot on", input[key])
		for i := key + 1; i < len(input); i++ {
			if input[i] < input[lowest] {
				fmt.Println(input[i], "is less than", input[lowest])
				lowest = i
			}
		}
		input = ArraySwap(input, key, lowest)

		fmt.Println(input)
	}
	return input
}
