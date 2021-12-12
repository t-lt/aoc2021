package aoc_3_1

import (
	"aoc2021/utils"
	"strconv"
)

func FindMostCommonBits(path string) string {
	s := utils.Get2DSliceFromFile(path)
	var r string
	var znum, onum int
	for i, j := 0, 0; j < len(s[0]); {
		switch rune(s[i][j]) {
		case '0':
			znum++
		case '1':
			onum++
		}
		if i == len(s)-1 {
			i=0
			j++
			var toAppend string
			if znum > onum {
				toAppend = "0"
			} else {
				toAppend = "1"
			}
			r += toAppend
			znum = 0
			onum = 0
		} else {
			i++
		}
	}
	return r
}

func FindLeastCommonBits(mstcommonbits string) uint {
	var sr string
	for _, v := range mstcommonbits {
		if rune(v) == '0' {
			sr += "1"
		} else {
			sr += "0"
		}
	}
	r, _ := strconv.ParseUint(sr, 2, 64)
	return uint(r)
}
func Aoc_3_1(path string) uint {
	str := FindMostCommonBits(path)
	lc := FindLeastCommonBits(str)
	mstcommon, _ := strconv.ParseUint(str, 2, 64)
	return uint(mstcommon) * lc 
}
