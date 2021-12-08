package aoc_2_1

import "testing"

func TestAoc2_1(t *testing.T){
	e :=150
	r := Aoc_2_1("aoc_2_1_testdata/testdata")
	if e != r {
		t.Errorf("Expected %d, got %d", e, r)
	}

}