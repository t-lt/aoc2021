package utils

import (
	"reflect"
	"testing"
)

func TestGetIntSliceFromFile(t *testing.T) {
	e := []int{1, 2, 3, 4, 5, 6}
	r := GetIntSliceFromFile("utils_testdata/inttestdata")
	if !reflect.DeepEqual(e, r) {
		t.Errorf("Expected %v, got %v", e, r)
	}
}

func TestSumOfIntSlice(t *testing.T) {
	i := []int{1, 2, 3, 4}
	e := 10
	r := SumOfIntSlice(i)
	if r != e {
		t.Errorf("Expected %d, got %d", e, r)
	}

}
func TestSumOfIntSliceRecursive(t *testing.T) {
	i := []int{1, 2, 3, 4}
	e := 10
	r := SumOfIntSliceRecursive(i)
	if r != e {
		t.Errorf("Expected %d, got %d", e, r)
	}

}

func TestParseMultiColumnsFile(t *testing.T) {
	e := [][]string{{"forward", "5"}, {"down", "5"}}
	r := ParseMultiColumnsFile("utils_testdata/csvtestdata", ' ')
	if !reflect.DeepEqual(e, r) {
		t.Errorf("Expected % v, got %v", e, r)
	}
}

func TestGet2dSliceFromFile(t *testing.T) {
	e := [][]byte{{48, 48, 49, 48, 48}, {49, 49, 49, 49, 48}}
	r := Get2DSliceFromFile("utils_testdata/2dbyteslice")
	if !reflect.DeepEqual(e, r) {
		t.Errorf("Expected %v, got %v", e, r)
	}
}
