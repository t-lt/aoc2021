package utils

import (
	"bufio"
	"encoding/csv"
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

func Get2DSliceFromFile(p string) [][]byte {
	f, err := os.Open(p)
	if err != nil {
		log.Default().Fatalf("Could not find file with path %s", p)
	}
	defer f.Close()

	var s [][]byte
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		parsed := []byte(scan.Text())
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

func ParseMultiColumnsFile(path string, comma rune) [][]string {
	f, e := os.Open(path)
	if e != nil {
		log.Fatalf("File %s not found", path)
	}
	r := csv.NewReader(f)
	r.Comma = comma
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalf("Error parsing csv : %v", e)
	}
	return records
}
