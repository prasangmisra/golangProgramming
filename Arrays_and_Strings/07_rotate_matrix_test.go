package Arrays_and_Strings

import "testing"

func TestRotateMatrixNegativeEmpty(t *testing.T) {
	result := RotateMatrix([][]int{})
	if result {
		t.Errorf("Expected false, got %v", result)
	}
}

func TestRotateMatrixNegativeNotSquareRow(t *testing.T) {
	result := RotateMatrix([][]int{{1, 2, 3}, {4, 5, 6}})
	if result {
		t.Errorf("Expected false, got %v", result)
	}
}

func TestRotateMatrixNegativeNotSquareCol(t *testing.T) {
	result := RotateMatrix([][]int{{1, 2}, {3, 4}, {5, 6}})
	if result {
		t.Errorf("Expected false, got %v", result)
	}
}

func TestRotateMatrixPositive(t *testing.T) {
	result := RotateMatrix([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}})
	if !result {
		t.Errorf("Expected false, got %v", result)
	}
}
