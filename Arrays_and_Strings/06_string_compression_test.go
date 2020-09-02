package Arrays_and_Strings

import (
	"testing"
)

func TestStringCompressionPositive(t *testing.T) {
	result := StringCompression("aabcccccaaa")
	if result != "a2b1c5a3" {
		t.Errorf("Expected a2blc5a3, got %v", result)
	}
}

func TestStringCompressionPositive1(t *testing.T) {
	result := StringCompression("aabcd")
	if result != "aabcd" {
		t.Errorf("Expected abcd, got %v", result)
	}
}
