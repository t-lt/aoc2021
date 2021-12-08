package aoc_2_2

import (
	"testing"
)

func TestAoc_2_2(t * testing.T) {
    e := 900
	r := Aoc_2_2("aoc_2_2_testdata/testdata")
	if r != e {
		t.Errorf("Expected %d, got %d", e, r)
	}	
}