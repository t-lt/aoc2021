package aoc_1_1

import "testing"


func TestIncreaseDepth(t *testing.T) {
	p := "part1_testdata/testdata"

	r := GetNumberOfIncrease(p)
	if r != 7 {
		t.Errorf("Expected 7 , got %d", r)
	}
}
