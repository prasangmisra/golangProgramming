package Arrays_and_Strings

import (
	"testing"
)

func TestZeroMatrix(t *testing.T) {
	result := ZeroMatrix([][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 0, 1}})
	if result[0][0] != 0 {
		t.Errorf("Expected all 0s, got %v", result)
	}
}
