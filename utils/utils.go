package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func GetIntSliceFromFile(p string) []int {
	f, err := os.Open(p)
	if err != nil {
		log.Default().Fatalf("Could not find file with path %s", p)
	}
	defer f.Close()

	var s []int
	scan := bufio.NewScanner(f)

	for scan.Scan() {
		parsed, e := strconv.Atoi(scan.Text())
		if e != nil {
			log.Default().Fatalln("Parsed a string inconvertible to int")
		}
		s = append(s, parsed)
	}
	return s
}

func SumOfIntSlice(s []int) int {
	var r int
	for _, e := range s {
		r += e
	}
	return r
}

func SumOfIntSliceRecursive(s []int) int {
	return add(s, len(s))
}

func add(s []int, i int) int {
	if i == 0 {
		return 0
	}
	return add(s, i-1) + s[i-1]
}
