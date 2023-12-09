package adventofcode

import (
	"bufio"
	"fmt"
	"strings"
)

func nextNumber(numbers []int) int {
	last := 0
	if l := len(numbers); l != 0 {
		last = numbers[l-1]
	}

	if allEqual(numbers) {
		return last
	}

	newSeq := []int{}
	for i := 0; i < len(numbers)-1; i++ {
		newSeq = append(newSeq, numbers[i+1]-numbers[i])
	}

	return last + nextNumber(newSeq)
}

func prevNumber(numbers []int) int {
	first := 0
	if l := len(numbers); l != 0 {
		first = numbers[0]
	}

	if allEqual(numbers) {
		return first
	}

	newSeq := []int{}
	for i := 0; i < len(numbers)-1; i++ {
		newSeq = append(newSeq, numbers[i+1]-numbers[i])
	}

	prev := prevNumber(newSeq)

	return first - prev
}

func day9(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		nums := Map(strings.Fields(line), parseInt)
		sum += nextNumber(nums)
	}

	return fmt.Sprintf("%v", sum)
}

func day9Ext(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		nums := Map(strings.Fields(line), parseInt)
		sum += prevNumber(nums)
	}

	return fmt.Sprintf("%v", sum)
}

func allEqual[T comparable](arr []T) bool {
	for _, v := range arr {
		if v != arr[0] {
			return false
		}
	}
	return true
}
