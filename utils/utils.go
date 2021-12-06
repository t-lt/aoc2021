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
