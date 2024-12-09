package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type puzzle struct {
	left  []int
	right []int
}

func NewSolution(input string) *puzzle {
	puzzle := &puzzle{
		left:  []int{},
		right: []int{},
	}

	file, err := os.Open(input)
	if err != nil {
		log.Fatalf("[day1] error opening input file: %+v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		locations := strings.Fields(scanner.Text())
		leftLocation, err := strconv.Atoi(locations[0])
		if err != nil {
			log.Fatalf("[day1] error converting location to int: %s\n", locations[0])
		}
		puzzle.left = append(puzzle.left, leftLocation)

		rightLocation, err := strconv.Atoi(locations[1])
		if err != nil {
			log.Fatalf("[day1] error converting location to int: %s\n", locations[1])
		}
		puzzle.right = append(puzzle.right, rightLocation)
	}

	return puzzle
}

func absDiff(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func (p *puzzle) Part1() (string, error) {
	// left := p.left[0:]
	right := p.right[0:]

	left := make([]int, len(p.left))
	copy(left, p.left)

	slices.Sort(left)
	slices.Sort(right)

	total := 0
	for i, l := range left {
		total += absDiff(l, right[i])
	}
	return fmt.Sprintf("%d", total), nil
}

func (p *puzzle) Part2() (string, error) {
	rightOccurances := map[int]int{}
	for _, l := range p.left {
		for _, r := range p.right {
			if l == r {
				rightOccurances[l]++
			}
		}
	}

	similarityScore := 0
	for num, occurances := range rightOccurances {
		similarityScore += (num * occurances)
	}

	return fmt.Sprintf("%d", similarityScore), nil
}
