package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type puzzle struct {
	matrix [][]rune
	start  [2]int
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
				p.start = [2]int{len(p.matrix), x}
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

func walk(path [][2]int, matrix [][]rune, direction int) [][2]int {
	y, x := path[len(path)-1][0], path[len(path)-1][1]
	if y < 0 || y > len(matrix)-1 || x < 0 || x > len(matrix[y])-1 {
		return path[:len(path)-1]
	}
	object := matrix[y][x]
	nextDirection := direction
	if object == '#' {
		nextDirection = nextDirections[direction]
		stepBackPath := path[:len(path)-1]
		nextY, nextX := stepBackPath[len(stepBackPath)-1][0]+directions[nextDirection][0], stepBackPath[len(stepBackPath)-1][1]+directions[nextDirection][1]
		stepBackPath = append(stepBackPath, [2]int{nextY, nextX})

		return walk(stepBackPath, matrix, nextDirection)
	}

	nextY, nextX := y+directions[nextDirection][0], x+directions[nextDirection][1]
	nextPath := append(path, [2]int{nextY, nextX})
	return walk(nextPath, matrix, nextDirection)
}

func walkTwo(path [][2]int, matrix [][]rune, direction int, newObstacles [][2]int, turns [][2]int) [][2]int {
	y, x := path[len(path)-1][0], path[len(path)-1][1]
	if y < 0 || y > len(matrix)-1 || x < 0 || x > len(matrix[y])-1 {
		return newObstacles
	}
	object := matrix[y][x]
	nextDirection := direction
	if object == '#' {
		nextDirection = nextDirections[direction]
		stepBackPath := path[:len(path)-1]
		nextY, nextX := stepBackPath[len(stepBackPath)-1][0]+directions[nextDirection][0], stepBackPath[len(stepBackPath)-1][1]+directions[nextDirection][1]
		nextStep := [2]int{nextY, nextX}
		stepBackPath = append(stepBackPath, nextStep)

		nextTurns := [][2]int{}
		if len(turns) == 2 {
			var obstacleY, obstacleX int
			if nextDirection == left || nextDirection == right {
				obstacleY = nextY
				if nextDirection == left {
					obstacleX = turns[0][1] - 2
				} else {
					obstacleX = turns[0][1] + 1
				}
			} else {
				obstacleX = nextX
				if nextDirection == top {
					obstacleY = turns[0][1] - 1
				} else {
					obstacleY = turns[0][1] + 1
				}
			}

			newObstacles = append(newObstacles, [2]int{obstacleY, obstacleX})
		} else {
			nextTurns = append(turns, [2]int{nextY, nextX})
		}
		return walkTwo(stepBackPath, matrix, nextDirection, newObstacles, nextTurns)
	}

	nextY, nextX := y+directions[nextDirection][0], x+directions[nextDirection][1]
	nextPath := append(path, [2]int{nextY, nextX})

	return walkTwo(nextPath, matrix, nextDirection, newObstacles, turns)
}

func (p *puzzle) Part1() (string, error) {
	path := walk([][2]int{p.start}, p.matrix, top)

	uniqueLocations := map[[2]int]bool{}
	for _, step := range path {
		uniqueLocations[step] = true
	}

	return fmt.Sprintf("%d", len(uniqueLocations)), nil
}

func (p *puzzle) Part2() (string, error) {
	newObstacles := walkTwo([][2]int{p.start}, p.matrix, top, [][2]int{}, [][2]int{})

	obs := map[[2]int]bool{}
	for _, o := range newObstacles {
		obs[o] = true
	}
	fmt.Printf("%+v\n", newObstacles)
	for y, row := range p.matrix {
		for x, r := range row {
			_, ok := obs[[2]int{y, x}]
			if !ok {
				fmt.Printf("%s", string(r))
			} else {
				fmt.Printf("%s", "O")
			}
		}
		fmt.Printf("\n")
	}
	return fmt.Sprintf("%d", len(newObstacles)), nil
}
