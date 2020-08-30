package Arrays_and_Strings

import "testing"

func TestStringUniquePass(t *testing.T) {
	result := FindUnique("abcd")
	if result != true {
		t.Errorf("Wrong output, expected true, got %v", result)
	}
}

func TestStringUniqueFail(t *testing.T) {
	result := FindUnique("abca")
	if result != false {
		t.Errorf("Wrong output, expected false, got %v", result)
	}
}

func TestStringUniqueEmpty(t *testing.T) {
	result := FindUnique("")
	if result != true {
		t.Errorf("Wrong output, expected true, got %v", result)
	}
}
