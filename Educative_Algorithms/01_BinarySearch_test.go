package Educative_Algorithms

import "testing"

func TestBinarySearchPositive(t *testing.T) {
	result := BinarySearch(31, []int{1, 2, 9, 20, 31, 45, 63, 70, 100, 101})
	if result != true {
		t.Errorf("Wrong output, expected true, got %v", result)
	}
}

func TestBinarySearchNegative(t *testing.T) {
	result := BinarySearch(3, []int{1, 2, 9, 20, 31, 45, 63, 70})
	if result != false {
		t.Errorf("Wrong output, expected false, got %v", result)
	}
}
