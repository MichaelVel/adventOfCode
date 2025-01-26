package adventofcode

import (
	"bufio"
	"fmt"
	"strings"
)

type Direction int

const (
	_ Direction = iota
	North
	South
	East
	West
)

func (d Direction) opossite() Direction {
	switch d {
	case North:
		return South
	case South:
		return North
	case East:
		return West
	case West:
		return East
	default:
		return 0
	}
}

type Pipe struct {
	entry, exit Direction
	Coord
}

func newPipe(coord Coord, from Direction, symbol byte) Pipe {
	p := Pipe{entry: from, Coord: coord}

	switch symbol {
	case '|':
		p.exit = IfElse(from == North, South, North)
	case '-':
		p.exit = IfElse(from == East, West, East)
	case 'L':
		p.exit = IfElse(from == North, East, North)
	case 'J':
		p.exit = IfElse(from == North, West, North)
	case '7':
		p.exit = IfElse(from == South, West, South)
	case 'F':
		p.exit = IfElse(from == South, East, South)
	}
	return p
}

func (p Pipe) nextCoord() Coord {
	coord := Coord{x: p.Coord.x, y: p.Coord.y}
	switch p.exit {
	case North:
		coord.y -= 1
	case South:
		coord.y += 1
	case East:
		coord.x += 1
	case West:
		coord.x -= 1
	}
	return coord
}

func (p Pipe) equal(o Pipe) bool {
	return p.Coord == o.Coord
}

type PipeField struct {
	grid  []string
	start Coord
}

func newPipeField(input string) *PipeField {
	pf := &PipeField{
		grid: []string{},
	}

	scanner := bufio.NewScanner(strings.NewReader(input))
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		if startPos := strings.Index(line, "S"); startPos != -1 {
			pf.start = Coord{x: startPos, y: i}
		}

		pf.grid = append(pf.grid, line)
	}

	return pf
}

func (p *PipeField) startLoop() (Pipe, Pipe) {
	adjacentCoords := Filter(p.start.crossCoords(), func(c Coord) bool {
		return p.coordInGrid(c)
	})

	pipes := []Pipe{}
	for _, coord := range adjacentCoords {
		symbol := p.grid[coord.y][coord.x]
		var from Direction
		if symbol == '.' {
			continue
		}

		switch {
		case coord.x == p.start.x && coord.y < p.start.y:
			from = South
		case coord.x == p.start.x && coord.y > p.start.y:
			from = North
		case coord.x < p.start.x && coord.y == p.start.y:
			from = East
		case coord.x > p.start.x && coord.y == p.start.y:
			from = West
		}

		if p.connectionIsPosible(from, symbol) {
			pipe := newPipe(coord, from, symbol)
			pipes = append(pipes, pipe)
		}
	}

	return pipes[0], pipes[1]

}

func (p *PipeField) coordInGrid(coord Coord) bool {
	lx, ly := len(p.grid[0]), len(p.grid)
	return 0 <= coord.x && coord.x < lx &&
		0 <= coord.y && coord.y < ly
}

func (p *PipeField) valueOfCoord(coord Coord) byte {
	return p.grid[coord.y][coord.x]
}

func (p *PipeField) connectionIsPosible(from Direction, symbol byte) bool {
	switch symbol {
	case '|':
		return from == South || from == North
	case '-':
		return from == West || from == East
	case 'L':
		return from == East || from == North
	case 'J':
		return from == West || from == North
	case '7':
		return from == West || from == South
	case 'F':
		return from == East || from == South
	}
	return false
}

func (p *PipeField) nextPipe(pipe Pipe) Pipe {
	nc := pipe.nextCoord()
	return newPipe(nc, pipe.exit.opossite(), p.valueOfCoord(nc))
}

func day10(input string) string {
	pf := newPipeField(input)
	n := 1

	firstPath, secondPath := pf.startLoop()
	for ; !firstPath.equal(secondPath); n++ {
		firstPath = pf.nextPipe(firstPath)
		secondPath = pf.nextPipe(secondPath)
	}

	return fmt.Sprintf("%v", n)
}
