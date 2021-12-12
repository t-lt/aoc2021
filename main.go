package main

import (
	"aoc2021/aoc_3_1"
	"fmt"
	"strconv"
)

func main() {
	mstcommonstr := aoc_3_1.FindMostCommonBits("input/day3")
	println(mstcommonstr)
	mstcommon, _ := strconv.ParseUint(mstcommonstr, 2, 64)
	leastcommon := aoc_3_1.FindLeastCommonBits(mstcommonstr)
	fmt.Printf("%d", uint(mstcommon)*leastcommon)
}
