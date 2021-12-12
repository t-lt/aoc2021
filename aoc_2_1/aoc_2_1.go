package aoc_2_1

import (
	"aoc2021/utils"
	"aoc2021/command"
	"log"
	"strconv"
)


type enginePostion struct {
	X, Y int
}
//first column = command, second command = distance 
func Aoc_2_1(path string) int{
	ur := utils.ParseMultiColumnsFile(path, ' ')
	var postion enginePostion
	
	for i, v := range ur {
      c := v[0]
	  d, e := strconv.Atoi(v[1])
	  
	  if e != nil {
		 log.Fatalf("Problem converting value to int at index %d", i)
	  }

	  switch c {
	    case  command.Forward:
		  postion.X += d 
		case command.Down:
			postion.Y += d
		case command.Up:
			postion.Y -=d	  
	  }
	}
	
	result := postion.X * postion.Y
    return result
}