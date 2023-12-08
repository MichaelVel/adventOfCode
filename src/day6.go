package adventofcode

import (
	"bufio"
	"fmt"
	"strings"
)

func distance(timePushed, totalTime int) int {
	return timePushed * (totalTime - timePushed)
}

func nWinsForTime(time, record int) int {
	count := 0
	for i := 0; i <= time; i++ {
		if distance(i, time) > record {
			count++
		}
	}
	return count
}

func day6(input string) string {
	mult := 1

	scanner := bufio.NewScanner(strings.NewReader(input))

	scanner.Scan()
	line := strings.Split(scanner.Text(), ":")[1]
	times := Map(strings.Fields(line), parseInt)

	scanner.Scan()
	line = strings.Split(scanner.Text(), ":")[1]
	records := Map(strings.Fields(line), parseInt)

	for i, time := range times {
		record := records[i]
		mult *= nWinsForTime(time, record)
	}

	return fmt.Sprintf("%v", mult)
}

func day6Ext(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))

	scanner.Scan()
	line := strings.Split(scanner.Text(), ":")[1]
	time := parseInt(strings.ReplaceAll(line, " ", ""))

	scanner.Scan()
	line = strings.Split(scanner.Text(), ":")[1]
	record := parseInt(strings.ReplaceAll(line, " ", ""))

	return fmt.Sprintf("%v", nWinsForTime(time, record))
}
