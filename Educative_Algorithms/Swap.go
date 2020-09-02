package Educative_Algorithms

func ArraySwap(array []int, firstIndex int, secondIndex int) []int {
	if firstIndex <= len(array) && secondIndex <= len(array) {
		temp := array[firstIndex]
		array[firstIndex] = array[secondIndex]
		array[secondIndex] = temp
	}
	return array
}
