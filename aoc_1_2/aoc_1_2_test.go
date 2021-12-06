package aoc_1_2

import "testing"

func TestIncreasingDepthSlidingWindow(t *testing.T) {
	r := GetIncreasingDepthSlidingWindow("aoc_1_2_testdata/testdata")
	e := 5
	if r != e {
		t.Errorf("Expected %d, got %d", e, r)
	}
}
