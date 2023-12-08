package adventofcode

import (
	"bufio"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	id      int
	winners []int
	numbers []int
}

func newCard(line string) *Card {
	card := &Card{}

	idx := card.parseId(line)
	fields := strings.Split(line[idx:], "|")

	if len(fields) < 2 {
		return card
	}

	card.winners = Map(strings.Fields(fields[0]), parseInt)
	card.numbers = Map(strings.Fields(fields[1]), parseInt)

	return card
}

func (c *Card) parseId(line string) int {
	idx := 0
	for ; idx < len(line) && !isDigit(line[idx]); idx++ {
	}
	start := idx
	for ; idx < len(line) && line[idx] != ':'; idx++ {
	}
	end := idx
	c.id = parseInt(line[start:end])
	return idx + 1
}

func (c *Card) winning_numbers() []int {
	arr := []int{}
	for _, n := range c.numbers {
		if slices.Contains(c.winners, n) {
			arr = append(arr, n)
		}
	}
	return arr
}

func day4(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	sum := 0

	for scanner.Scan() {
		card := newCard(scanner.Text())
		n := len(card.winning_numbers()) - 1
		if n >= 0 {

			sum += int(math.Pow(2, float64(n)))
		}
	}

	return fmt.Sprintf("%v", sum)
}

func day4Ext(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	cards := []*Card{}
	counts := []int{}

	for scanner.Scan() {
		cards = append(cards, newCard(scanner.Text()))
		counts = append(counts, 1)
	}

	for i, card := range cards {
		for n := len(card.winning_numbers()); n > 0; n-- {
			if i+n < len(counts) {
				ncopies := counts[i]
				counts[i+n] += ncopies
			}
		}
	}

	sum := 0
	for _, n := range counts {
		sum += n
	}

	return fmt.Sprintf("%v", sum)
}

func Map[I any, O any](vs []I, f func(I) O) []O {
	vsm := make([]O, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func parseInt(val string) int {
	if number, err := strconv.ParseInt(val, 10, 64); err == nil {
		return int(number)
	}
	return 0
}
