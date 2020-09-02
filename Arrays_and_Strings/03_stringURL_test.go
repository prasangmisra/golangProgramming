package Arrays_and_Strings

import (
	"testing"
)

func TestStringURL1(t *testing.T) {
	result := StringURL("abc def")
	if result != "abc%20def" {
		t.Errorf("Expected abc20def, got %v", result)
	}
}

func TestStringURL2(t *testing.T) {
	result := StringURL("abcdef")
	if result != "abcdef" {
		t.Errorf("Expected abcdef, got %v", result)
	}
}

func TestStringURL3(t *testing.T) {
	result := StringURL("")
	if result != "" {
		t.Errorf("Expected blank, got %v", result)
	}
}

func TestStringURL4(t *testing.T) {
	result := StringURL(" ")
	if result != " " {
		t.Errorf("Expected 20, got %v", result)
	}
}
