package Educative_Algorithms

import (
	"reflect"
	"testing"
)

func Test_Selection_Sort_1(t *testing.T) {
	result := SelectionSort([]int{7, 6, 5, 4, 3, 2, 1})
	temp := []int{1, 2, 3, 4, 5, 6, 7}
	if !reflect.DeepEqual(result, temp) {
		t.Errorf("Expected sorted array, got %v", result)
	}
}

func Test_Selection_Sort_2(t *testing.T) {
	result := SelectionSort([]int{4, 3, 2, 1})
	temp := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(result, temp) {
		t.Errorf("Expected sorted array, got %v", result)
	}
}

func Test_Selection_Sort_3(t *testing.T) {
	result := SelectionSort([]int{4, 3, 2, 1, 0, -99, -1})
	temp := []int{-99, -1, 0, 1, 2, 3, 4}
	if !reflect.DeepEqual(result, temp) {
		t.Errorf("Expected sorted array, got %v", result)
	}
}

func Test_Selection_Sort_4(t *testing.T) {
	result := SelectionSort([]int{1, 2, 3, 4})
	temp := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(result, temp) {
		t.Errorf("Expected sorted array, got %v", result)
	}
}
func Test_Selection_Sort_5(t *testing.T) {
	result := SelectionSort([]int{2})
	temp := []int{2}
	if !reflect.DeepEqual(result, temp) {
		t.Errorf("Expected sorted array, got %v", result)
	}
}

func Test_Selection_Sort_6(t *testing.T) {
	result := SelectionSort([]int{})
	temp := []int{}
	if !reflect.DeepEqual(result, temp) {
		t.Errorf("Expected sorted array, got %v", result)
	}
}
