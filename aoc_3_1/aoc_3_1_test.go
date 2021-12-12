package aoc_3_1

import (
	"strconv"
	"testing"
)

func TestFindMostCommonBits(t *testing.T) {
	var e int64 = 22
	s := FindMostCommonBits("aoc_3_1_testdata/testdata")
	r, _ := strconv.ParseInt(s, 2, 64)
	if e != r {
		t.Errorf("Expected %d, got %d", e, r)
	}
}

func TestFindLeastCommonBits(t *testing.T) {
	mstComBits  := "10110"
	var e uint = 9
	r := FindLeastCommonBits(mstComBits)
	if e != r {
		t.Errorf("Expected %d, got %d", e, r)
	}
}

func TestAoc_3_1( t *testing.T){
	var e uint = 198
	r := Aoc_3_1("aoc_3_1_testdata/testdata")
	if e != r {
		t.Errorf("Expected %d, got %d", e, r)
	}
}
