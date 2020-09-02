package Arrays_and_Strings

import (
	"testing"
)

func TestPalindromePermutationPositiveAllEven(t *testing.T) {
	result := PalindromePermutation("aabbcc")
	if !result {
		t.Errorf("Expected true, received %v", result)
	}
}

func TestPalindromePermutationPositiveOneOdd(t *testing.T) {
	result := PalindromePermutation("aabbccd")
	if !result {
		t.Errorf("Expected true, received %v", result)
	}
}

func TestPalindromePermutationNegativeTwoOdd(t *testing.T) {
	result := PalindromePermutation("abbccd")
	if result {
		t.Errorf("Expected false, received %v", result)
	}
}
func TestPalindromePermutationPositiveEmpty(t *testing.T) {
	result := PalindromePermutation("")
	if !result {
		t.Errorf("Expected true, received %v", result)
	}
}
func TestPalindromePermutationPositiveSingle(t *testing.T) {
	result := PalindromePermutation("a")
	if !result {
		t.Errorf("Expected true, received %v", result)
	}
}
