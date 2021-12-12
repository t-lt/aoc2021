package aoc_2_2

import (
	"aoc2021/command"
	"aoc2021/utils"
	"log"
	"strconv"
)

type enginePostion struct {
	X,Y int
	Aim int
}

func Aoc_2_2(path string) int {
	var position enginePostion
	p := utils.ParseMultiColumnsFile(path , ' ')
	for i, ur := range p {
       d, err := strconv.Atoi(ur[1])
	   if err != nil {
		   log.Fatalf("Can't convert to int at index %d", i)
	   }
	   switch ur[0] {
	   case command.Up:
		position.Aim -= d
       case command.Down:
		position.Aim += d
	   case command.Forward:
		 position.X +=d
		 position.Y += position.Aim * d
	   }
	}
    return position.X * position.Y
}