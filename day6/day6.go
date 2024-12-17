package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type matrix [][]rune

func (m matrix) isOutOfBounds(l location) bool {
	y, x := l[0], l[1]
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
	start  step
}

func NewSolution(filename string) *puzzle {
	p := &puzzle{
		matrix: [][]rune{},
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("[day6] error opening input file: %+v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := []rune{}
		for x, object := range scanner.Text() {
			if object == '^' {

				p.start = step{y: len(p.matrix), x: x, direction: top}
			}
			row = append(row, object)
		}
		p.matrix = append(p.matrix, row)
	}

	return p
}

const (
	top = iota
	right
	bottom
	left
)

type location = [2]int

type step struct {
	y         int
	x         int
	direction int
}

func (s *step) next() step {
	nextY, nextX := s.y+directions[s.direction][0], s.x+directions[s.direction][1]
	return step{
		y:         nextY,
		x:         nextX,
		direction: s.direction,
	}
}

func (s *step) turn(prevDirection int) step {
	nextDirection := nextDirections[prevDirection]
	return step{
		y:         s.y,
		x:         s.x,
		direction: nextDirection,
	}
}

var directions = map[int][2]int{
	top:    {-1, 0},
	right:  {0, 1},
	bottom: {1, 0},
	left:   {0, -1},
}
var nextDirections = map[int]int{
	top:    right,
	right:  bottom,
	bottom: left,
	left:   top,
}

func findWayOut(path []step, matrix [][]rune) []step {
	currStep := path[len(path)-1]
	y, x := currStep.y, currStep.x
	if y < 0 || y > len(matrix)-1 || x < 0 || x > len(matrix[y])-1 {
		return path[:len(path)-1]
	}
	object := matrix[y][x]
	if object == '#' {
		poppedSteps := path[:len(path)-1]
		nextStep := poppedSteps[len(poppedSteps)-1].turn(currStep.direction)
		path = append(poppedSteps, nextStep)

		return findWayOut(path, matrix)
	}

	nextStep := currStep.next()
	path = append(path, nextStep)
	return findWayOut(path, matrix)
}

func checkLoop(steps []step, matrix matrix, additionalObstacle location, seen map[step]int) bool {
	currStep := steps[len(steps)-1]
	if matrix.isOutOfBounds(location{currStep.y, currStep.x}) || matrix.isOutOfBounds(additionalObstacle) {
		return false
	}

	if matrix[additionalObstacle[0]][additionalObstacle[1]] == '#' || matrix[additionalObstacle[0]][additionalObstacle[1]] == '^' {
		return false
	}

	count, ok := seen[currStep]
	if ok && count > 1 {
		return true
	}

	currY, currX := currStep.y, currStep.x
	if matrix[currY][currX] == '#' || (currY == additionalObstacle[0] && currX == additionalObstacle[1]) {
		poppedSteps := steps[:len(steps)-1]
		nextStep := poppedSteps[len(poppedSteps)-1].turn(currStep.direction)
		steps = append(poppedSteps, nextStep)
		seen[currStep]++
		return checkLoop(steps, matrix, additionalObstacle, seen)
	}

	nextY, nextX := currY+directions[currStep.direction][0], currX+directions[currStep.direction][1]
	nextStep := step{
		y:         nextY,
		x:         nextX,
		direction: currStep.direction,
	}
	steps = append(steps, nextStep)
	seen[currStep]++
	return checkLoop(steps, matrix, additionalObstacle, seen)
}

func (p *puzzle) Part1() (string, error) {
	path := findWayOut([]step{{y: p.start.y, x: p.start.x, direction: top}}, p.matrix)

	uniqueLocations := map[location]bool{}
	for _, step := range path {
		uniqueLocations[location{step.y, step.x}] = true
	}

	return fmt.Sprintf("%d", len(uniqueLocations)), nil
}

func (p *puzzle) Part2() (string, error) {
	path := findWayOut([]step{{y: p.start.y, x: p.start.x, direction: top}}, p.matrix)

	uniqueLocations := map[location]bool{}
	for _, step := range path {
		uniqueLocations[location{step.y, step.x}] = true
	}

	total := 0
	for newObstacle := range uniqueLocations {
		isLoop := checkLoop([]step{{y: p.start.y, x: p.start.x, direction: top}}, p.matrix, newObstacle, map[step]int{})
		if isLoop {
			total++
		}
	}

	return fmt.Sprintf("%d", total), nil
}
