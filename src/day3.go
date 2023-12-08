package adventofcode

import (
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Coord struct {
	x, y int
}

type Number struct {
	value int
	pairs []Coord
}

func (n *Number) adjacentCoords() []Coord {
	adjacents := []Coord{}
	for _, pair := range n.pairs {
		adjacents = append(adjacents, Coord{pair.x - 1, pair.y - 1})
		adjacents = append(adjacents, Coord{pair.x - 1, pair.y})
		adjacents = append(adjacents, Coord{pair.x - 1, pair.y + 1})
		adjacents = append(adjacents, Coord{pair.x, pair.y - 1})
		adjacents = append(adjacents, Coord{pair.x, pair.y + 1})
		adjacents = append(adjacents, Coord{pair.x + 1, pair.y - 1})
		adjacents = append(adjacents, Coord{pair.x + 1, pair.y})
		adjacents = append(adjacents, Coord{pair.x + 1, pair.y + 1})
	}

	return adjacents
}

type Engine struct {
	raw     []string
	numbers []Number
	symbols []Coord
}

func newEngine(input string) *Engine {
	engine := &Engine{
		raw:     []string{},
		numbers: []Number{},
		symbols: []Coord{},
	}

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		engine.raw = append(engine.raw, scanner.Text())
	}

	for y := 0; y < len(engine.raw); y++ {
		row := engine.raw[y]
		for x := 0; x < len(row); x++ {
			cell := row[x]
			if isDigit(cell) {
				x = engine.addNumber(x, y)
			} else if isSymbol(cell) {
				engine.addSymbol(x, y)
			}

		}
	}

	return engine
}

func (e *Engine) addSymbol(x, y int) {
	e.symbols = append(e.symbols, Coord{x, y})

}

func (e *Engine) addNumber(x, y int) int {
	number := Number{
		value: 0,
		pairs: []Coord{},
	}

	row := e.raw[y]
	lx := x
	for ; lx < len(row) && isDigit(row[lx]); lx++ {
		number.pairs = append(number.pairs, struct{ x, y int }{lx, y})
	}

	if n, err := strconv.ParseInt(row[x:lx], 10, 64); err == nil {
		number.value = int(n)
	}

	e.numbers = append(e.numbers, number)

	return lx - 1
}

func (e *Engine) partNumbers() []Number {
	arr := []Number{}
	for _, num := range e.numbers {
		if e.isNumberAdjacentToSymbol(num) {
			arr = append(arr, num)
		}
	}
	return arr
}

func (e *Engine) isNumberAdjacentToSymbol(num Number) bool {
	for _, adjacent := range num.adjacentCoords() {
		if slices.Contains(e.symbols, adjacent) {
			return true
		}
	}
	return false
}

func day3(input string) string {
	sum := 0
	engine := newEngine(input)
	for _, partNumber := range engine.partNumbers() {
		sum += partNumber.value
	}

	return fmt.Sprintf("%v", sum)
}

func isSymbol(ch byte) bool {
	return !isDigit(ch) && ch != '.'
}
