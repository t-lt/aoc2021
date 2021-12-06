package aoc_1_1

import "aoc2021/utils"

func GetNumberOfIncrease(p string) int {
	var counter int
	s := utils.GetIntSliceFromFile(p)
	for i, t := range s {
		if i == 0 {
			continue
		}
		if s[i-1] < t {
			counter++
		}
	}
	return counter
}
