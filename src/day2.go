package adventofcode

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Set struct {
	blue  int
	red   int
	green int
}

type Game struct {
	id   int
	sets []*Set
}

func newGame(line string) *Game {
	game := &Game{}
	idx := game.parseId(line)
	for _, text := range strings.Split(line[idx:], ";") {
		game.addSet(text)
	}

	return game
}

func (g *Game) parseId(line string) int {
	idx := 0
	for ; idx < len(line) && !isDigit(line[idx]); idx++ {
	}
	start := idx
	for ; idx < len(line) && line[idx] != ':'; idx++ {
	}
	end := idx
	if number, err := strconv.ParseInt(line[start:end], 10, 64); err == nil {
		g.id = int(number)
	}
	return idx + 1
}

func (g *Game) addSet(segment string) {
	set := &Set{}

	for _, text := range strings.Split(segment, ",") {
		fields := strings.Fields(text)
		if len(fields) < 2 {
			return
		}

		number, _ := strconv.ParseInt(fields[0], 10, 64)
		switch fields[1] {
		case "red":
			set.red = int(number)
		case "green":
			set.green = int(number)
		case "blue":
			set.blue = int(number)
		}
	}

	g.sets = append(g.sets, set)
}

func (g *Game) isValid(red, green, blue int) bool {
	for _, set := range g.sets {
		if set.red > red || set.green > green || set.blue > blue {
			return false
		}
	}

	return true
}

func (g *Game) minCubsPossible() (red, green, blue int) {
	for _, set := range g.sets {
		if set.red > red {
			red = set.red
		}
		if set.green > green {
			green = set.green
		}
		if set.blue > blue {
			blue = set.blue
		}
	}
	return red, green, blue
}

func day2(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	sum := 0

	for scanner.Scan() {
		game := newGame(scanner.Text())

		if game.isValid(12, 13, 14) {
			sum += game.id
		}

	}

	return fmt.Sprintf("%v", sum)
}

func day2Ext(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	sum := 0

	for scanner.Scan() {
		game := newGame(scanner.Text())
		red, green, blue := game.minCubsPossible()
		sum += red * green * blue
	}

	return fmt.Sprintf("%v", sum)
}
