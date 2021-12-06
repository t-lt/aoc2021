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
