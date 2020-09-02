package Educative_Algorithms

import (
	"reflect"
	"testing"
)

func TestInsertToArray1(t *testing.T) {
	result := InsertToArray([]int{3, 5, 7, 11, 13, 8, 9, 6}, 4, 8)
	temp := []int{3, 5, 7, 8, 11, 13, 9, 6}
	if !reflect.DeepEqual(result, temp) {
		t.Errorf("Expected sorted array, got %v", result)
	}
}

func TestInsertToArray2(t *testing.T) {
	result := InsertToArray([]int{3, 5, 7, 11, 13, 8, 9, 6}, 0, 8)
	temp := []int{3, 5, 7, 11, 13, 8, 9, 6}
	if !reflect.DeepEqual(result, temp) {
		t.Errorf("Expected sorted array, got %v", result)
	}
}

func TestInsertionSort1(t *testing.T) {
	result := InsertionSort([]int{5, 2, 47, 7, 23, 13, 41, 29, 53, 11})
	temp := []int{2, 5, 7, 11, 13, 23, 29, 41, 47, 53}
	if !reflect.DeepEqual(result, temp) {
		t.Errorf("Expected sorted array, got %v", result)
	}
}
