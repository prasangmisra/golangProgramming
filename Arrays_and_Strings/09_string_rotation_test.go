package Arrays_and_Strings

import (
	"testing"
)

func TestStringRotation(t *testing.T) {
	result := IsSubstring("waterbottle", "erbottlewat")
	if !result {
		t.Errorf("Expected true, got %v", result)
	}
}
