package Educative_Algorithms

import (
	"testing"
)

func TestSwapPositive(t *testing.T) {
	result := ArraySwap([]int{1, 2, 3, 4, 5}, 0, 1)
	if result[0] != 2 {
		t.Errorf("wrong output, expected 2, got %v", result)
	}
}

func TestSwapOutOfBound(t *testing.T) {
	result := ArraySwap([]int{1, 2, 3, 4, 5}, 0, 6)
	if result[0] != 1 {
		t.Errorf("wrong output, expected 1, got %v", result)
	}
}

func TestSwapEmpty(t *testing.T) {
	result := ArraySwap([]int{}, 0, 1)
	if len(result) != 0 {
		t.Errorf("wrong output, expected blank, got %v", result)
	}
}
