package adventofcode

import (
	"fmt"
	"os"
)

type SolutionFunc func(string) string

func defaultFunction(input string) string {
	return input
}

func Main(filename, ex string) {
	buff, err := os.ReadFile(filename)
	if err != nil {
		return
	}

	msg := fmt.Sprintf("exercise %s not implemented", ex)
	answer := defaultFunction(msg)
	if fn, ok := Solutions[ex]; ok {
		answer = fn(string(buff))
	}

	println(answer)
}

var Solutions = map[string]SolutionFunc{
	"1":    day1,
	"1ext": day1Ext,
	"2":    day2,
	"2ext": day2Ext,
	"3":    day3,
	"4":    day4,
	"4ext": day4Ext,
	"5":    day5,
	"5ext": day5Ext,
	"6":    day6,
	"6ext": day6Ext,
	"7":    day7,
	"7ext": day7Ext,
	"8":    day8,
	"8ext": day8Ext,
}
