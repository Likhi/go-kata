package kata

import (
	"testing"
)

func TestRangeBitCount(t *testing.T) {
	testValue27 := RangeBitCount(2, 7)
	if testValue27 != 11 {
		t.Errorf("Result for a,b = 2,7 should be 11 but got %v", testValue27)
	}

	test01 := RangeBitCount(0, 1)
	if test01 != 1 {
		t.Errorf("Result for a,b = 0,1 should 1 but got %v", test01)
	}

	test44 := RangeBitCount(4, 4)
	if test44 != 1 {
		t.Errorf("Result for a,b = 4,4 should 1 but got %v", test44)
	}

}

func TestBlackOrWhiteKey(t *testing.T) {
	if BlackOrWhiteKey(1) != "white" {
		t.Errorf("1 should be white")
	}
	if BlackOrWhiteKey(5) != "black" {
		t.Errorf("5 should be black")
	}
}
