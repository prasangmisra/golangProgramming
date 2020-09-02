package Arrays_and_Strings

import (
	"testing"
)

func TestOneWayPositive(t *testing.T) {
	result := OneWay("abcd", "abcd")
	if !result {
		t.Errorf("Expected true, received %v", result)
	}
}

func TestOneWayNegativeSameLength1(t *testing.T) {
	result := OneWay("abcd", "abce")
	if !result {
		t.Errorf("Expected true, received %v", result)
	}
}

func TestOneWayNegativeSameLength2(t *testing.T) {
	result := OneWay("abcd", "abde")
	if result {
		t.Errorf("Expected true, received %v", result)
	}
}

func TestOneWayNegativeDifferentLength3(t *testing.T) {
	result := OneWay("eabcd", "abcd")
	if !result {
		t.Errorf("Expected true, received %v", result)
	}
}
