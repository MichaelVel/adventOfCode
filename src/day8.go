package adventofcode

import (
	"bufio"
	"fmt"
	"strings"
	"unicode"
)

type Network struct {
	nodes map[string][2]string
}

func (n *Network) traverse(
	instructions string,
	from string,
	to func(string) bool,
) int {
	nInstr := len(instructions)
	steps := 0

	next := from
	for i := 0; to(next); i++ {
		nextInstr := instructions[i%nInstr]
		if nextInstr == 'R' {
			next = n.nodes[next][1]
		} else if nextInstr == 'L' {
			next = n.nodes[next][0]
		} else {
			break
		}
		steps++
	}

	return steps
}

func (n *Network) traverseAll(instructions string) int {
	aNodes := n.getNodes('A')
	loopDistances := Map(aNodes, func(node string) int {
		f := func(s string) bool { return s[2] != 'Z' }
		return n.traverse(instructions, node, f)
	})

	return LCM(loopDistances...)
}

func (n *Network) getNodes(letter byte) []string {
	nodes := []string{}
	for node := range n.nodes {
		if node[2] == letter {
			nodes = append(nodes, node)
		}
	}
	return nodes
}

func (n *Network) addNode(line string) {
	chuncks := strings.Split(line, " = ")
	node := chuncks[0]
	childs := strings.FieldsFunc(chuncks[1], func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})

	n.nodes[node] = [2]string{childs[0], childs[1]}
}

func day8(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))

	scanner.Scan()
	instructions := scanner.Text()
	scanner.Scan()

	net := Network{map[string][2]string{}}

	for scanner.Scan() {
		net.addNode(scanner.Text())
	}

	steps := net.traverse(instructions, "AAA", func(s string) bool {
		return s != "ZZZ"
	})

	return fmt.Sprintf("%v", steps)
}

func day8Ext(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))

	scanner.Scan()
	instructions := scanner.Text()
	scanner.Scan()

	net := Network{map[string][2]string{}}

	for scanner.Scan() {
		net.addNode(scanner.Text())
	}

	steps := net.traverseAll(instructions)

	return fmt.Sprintf("%v", steps)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
func LCM(integers ...int) int {
	result := 1
	for _, i := range integers {
		result = result * i / GCD(result, i)
	}
	return result
}
