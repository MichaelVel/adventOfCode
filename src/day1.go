package adventofcode

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

var numbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func day1(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		sum += parseNumberFromLine(line, false)
	}

	return fmt.Sprintf("%v", sum)
}

func day1Ext(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		sum += parseNumberFromLine(line, true)
	}

	return fmt.Sprintf("%v", sum)
}

func parseNumberFromLine(line string, checkForWriten bool) int {
	n := ""
	idx := 0

	for i := 0; i < len(line); i += 1 {
		ch := line[i]
		if isDigit(ch) {
			idx = i
			n += string(ch)
			break
		}

		if !checkForWriten {
			continue
		}

		if num, ok := peekWritenNumber(line[i:]); ok {
			idx = i
			n += fmt.Sprintf("%v", num)
			break
		}
	}

	// search in [idx, len(line]; inclusive to check for the case with only
	// one number in the string: 'a1a" = 11.
	for i := len(line) - 1; i >= idx; i -= 1 {
		ch := line[i]
		if isDigit(ch) {
			n += string(ch)
			break
		}

		if !checkForWriten {
			continue
		}

		if num, ok := peekReverseWritenNumber(line[idx : i+1]); ok {
			n += fmt.Sprintf("%v", num)
			break
		}

	}

	number, _ := strconv.ParseInt(n, 10, 64)

	return int(number)
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func peekWritenNumber(slice string) (n int, ok bool) {
	for i := 3; i <= 5; i += 1 {
		if len(slice) < i {
			return 0, false
		}

		n := slice[:i]
		if number, ok := numbers[n]; ok {
			return number, ok
		}
	}

	return 0, false
}

func peekReverseWritenNumber(slice string) (n int, ok bool) {
	for i := 3; i <= 5; i += 1 {
		if len(slice) < i {
			return 0, false
		}

		n := slice[len(slice)-i:]
		if number, ok := numbers[n]; ok {
			return number, ok
		}
	}

	return 0, false
}
