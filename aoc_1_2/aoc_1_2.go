package aoc_1_2

import "aoc2021/utils"

func GetIncreasingDepthSlidingWindow(path string) int {
	var counter int
    var precedingResult int 

	s := utils.GetIntSliceFromFile(path)
	for i := 0; i <= len(s)-3; i++ {
		r := utils.SumOfIntSlice(s[i:i+3])
		if i != 0 && r > precedingResult {
			counter++
		}
		precedingResult = r
	}
	return counter

}
