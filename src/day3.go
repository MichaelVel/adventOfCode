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

func (c *Coord) adjacentCoords() []Coord {
	return []Coord{
		Coord{c.x - 1, c.y - 1},
		Coord{c.x - 1, c.y},
		Coord{c.x - 1, c.y + 1},
		Coord{c.x, c.y - 1},
		Coord{c.x, c.y + 1},
		Coord{c.x + 1, c.y - 1},
		Coord{c.x + 1, c.y},
		Coord{c.x + 1, c.y + 1},
	}
}

func (c *Coord) crossCoords() []Coord {
	return []Coord{
		Coord{c.x - 1, c.y},
		Coord{c.x, c.y - 1},
		Coord{c.x, c.y + 1},
		Coord{c.x + 1, c.y},
	}
}

type Number struct {
	value int
	pairs []Coord
}

func (n *Number) adjacentCoords() []Coord {
	adjacents := []Coord{}
	for _, pair := range n.pairs {
		adjacents = append(adjacents, pair.adjacentCoords()...)
	}

	return adjacents
}

func (n *Number) coordInRange(c Coord) bool {
	for _, innerCord := range n.pairs {
		if c == innerCord {
			return true
		}
	}
	return false
}

func (n *Number) String() string {
	// Simple hack to hash the type, given that the coords of the number
	// always are unique.
	return fmt.Sprintf("%v", n.pairs)
}

type Symbol struct {
	value byte
	Coord
}

func (s *Symbol) gearRatio(nums []Number) (int, bool) {
	if s.value != '*' {
		return 0, false
	}
	adjacentNumbers := map[string]int{}
	for _, coord := range s.adjacentCoords() {
		for _, num := range nums {
			_, present := adjacentNumbers[num.String()]
			if num.coordInRange(coord) && !present {
				adjacentNumbers[num.String()] = num.value
			}
		}
	}

	if len(adjacentNumbers) != 2 {
		return 0, false
	}

	ratio := 1
	for _, val := range adjacentNumbers {
		ratio *= val
	}

	return ratio, true

}

type Engine struct {
	raw     []string
	numbers []Number
	symbols []Symbol
}

func newEngine(input string) *Engine {
	engine := &Engine{
		raw:     []string{},
		numbers: []Number{},
		symbols: []Symbol{},
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
	e.symbols = append(e.symbols, Symbol{e.raw[y][x], Coord{x, y}})

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

func (e *Engine) gearsRatios() int {
	sumRatios := 0
	partNums := e.partNumbers()
	for _, s := range e.symbols {
		if ratio, ok := s.gearRatio(partNums); ok {
			sumRatios += ratio
		}
	}
	return sumRatios
}

func (e *Engine) isNumberAdjacentToSymbol(num Number) bool {
	coords := Map(e.symbols, func(s Symbol) Coord { return s.Coord })
	for _, adjacent := range num.adjacentCoords() {
		if slices.Contains(coords, adjacent) {
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

func day3Ext(input string) string {
	engine := newEngine(input)
	return fmt.Sprintf("%v", engine.gearsRatios())
}

func isSymbol(ch byte) bool {
	return !isDigit(ch) && ch != '.'
}
