package Arrays_and_Strings

import "testing"

func TestStringPermutationPass(t *testing.T) {
	result := FindUnique("abcd")
	if result != true {
		t.Errorf("Wrong output, expected true, got %v", result)
	}
}

func TestStringPermutationFail(t *testing.T) {
	result := FindUnique("abca")
	if result != false {
		t.Errorf("Wrong output, expected false, got %v", result)
	}
}

func TestStringPermutationEmpty(t *testing.T) {
	result := FindUnique("")
	if result != true {
		t.Errorf("Wrong output, expected true, got %v", result)
	}
}
