package day8

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type matrix [][]string

type location struct {
	y int
	x int
}

func (m matrix) isOutOfBounds(l location) bool {
	y, x := l.y, l.x
	if y < 0 || y >= len(m) {
		return true
	}

	if x < 0 || x >= len(m[y]) {
		return true
	}
	return false
}

type puzzle struct {
	matrix matrix
}

func NewSolution(filename string) *puzzle {
	p := &puzzle{
		matrix: [][]string{},
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("[day8] error opening input file: %+v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := []string{}
		for _, object := range scanner.Text() {
			row = append(row, string(object))
		}
		p.matrix = append(p.matrix, row)
	}

	return p
}

func findAntinodes(node location, nodes []location, m matrix) []location {
	antinodes := []location{}
	for _, n := range nodes {
		if n.y == node.y && n.x == node.x {
			continue
		}
		offsetY := n.y - node.y
		offsetX := n.x - node.x

		a := location{y: node.y + offsetY*2, x: node.x + offsetX*2}
		if !m.isOutOfBounds(a) {
			antinodes = append(antinodes, a)
		}

	}
	return antinodes
}

func findRowOfAntinodes(node location, nodes []location, m matrix) []location {
	antinodes := []location{}
	if len(nodes) > 1 {
		antinodes = append(antinodes, nodes...)
	}
	for _, n := range nodes {
		if n.y == node.y && n.x == node.x {
			continue
		}
		offsetY := n.y - node.y
		offsetX := n.x - node.x

		forwardMultiplier := 2
		for {
			l := location{y: node.y + offsetY*forwardMultiplier, x: node.x + offsetX*forwardMultiplier}
			if m.isOutOfBounds(l) {
				break
			}

			antinodes = append(antinodes, l)
			forwardMultiplier++
		}

		backwardMultiplier := 2
		for {
			l := location{y: node.y + offsetY*backwardMultiplier*-1, x: node.x + offsetX*backwardMultiplier*-1}
			if m.isOutOfBounds(l) {
				break
			}

			antinodes = append(antinodes, l)
			backwardMultiplier++
		}
	}
	return antinodes
}

func (p *puzzle) Part1() (string, error) {
	frequencyNodes := map[string][]location{}
	for y, row := range p.matrix {
		for x, frequency := range row {
			if frequency == "." {
				continue
			}

			_, ok := frequencyNodes[frequency]
			if !ok {
				frequencyNodes[frequency] = []location{}
			}

			frequencyNodes[frequency] = append(frequencyNodes[frequency], location{y: y, x: x})
		}
	}

	antinodeLocations := map[string][]location{}
	for frequency, nodes := range frequencyNodes {
		_, ok := antinodeLocations[frequency]
		if !ok {
			antinodeLocations[frequency] = []location{}
		}

		for _, a := range nodes {
			antinodes := findAntinodes(a, nodes, p.matrix)
			antinodeLocations[frequency] = append(antinodeLocations[frequency], antinodes...)
		}
	}

	uniqueAntinodes := map[location]interface{}{}
	for _, antinodes := range antinodeLocations {
		for _, a := range antinodes {
			uniqueAntinodes[a] = struct{}{}
		}
	}

	total := len(uniqueAntinodes)
	return fmt.Sprintf("%d", total), nil
}

func (p *puzzle) Part2() (string, error) {
	frequencyNodes := map[string][]location{}
	for y, row := range p.matrix {
		for x, frequency := range row {
			if frequency == "." {
				continue
			}

			_, ok := frequencyNodes[frequency]
			if !ok {
				frequencyNodes[frequency] = []location{}
			}

			frequencyNodes[frequency] = append(frequencyNodes[frequency], location{y: y, x: x})
		}
	}

	antinodeLocations := map[string][]location{}
	for frequency, nodes := range frequencyNodes {
		_, ok := antinodeLocations[frequency]
		if !ok {
			antinodeLocations[frequency] = []location{}
		}

		for _, a := range nodes {
			antinodes := findRowOfAntinodes(a, nodes, p.matrix)
			antinodeLocations[frequency] = append(antinodeLocations[frequency], antinodes...)
		}
	}

	uniqueAntinodes := map[location]interface{}{}
	for _, antinodes := range antinodeLocations {
		for _, a := range antinodes {
			uniqueAntinodes[a] = struct{}{}
		}
	}

	total := len(uniqueAntinodes)
	return fmt.Sprintf("%d", total), nil
}
