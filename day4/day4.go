package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

type puzzle struct {
	matrix [][]rune
}

func NewSolution(filename string) *puzzle {
	p := &puzzle{
		matrix: [][]rune{},
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening input file: %+v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		p.matrix = append(p.matrix, []rune{})
		for _, r := range scanner.Text() {
			p.matrix[row] = append(p.matrix[row], r)
		}
		row++
	}

	return p
}

const (
	top = iota
	topRight
	right
	bottomRight
	bottom
	bottomLeft
	left
	topLeft
)

var directions map[int][2]int = map[int][2]int{
	top:         {-1, 0},
	topRight:    {-1, 1},
	right:       {0, 1},
	bottomRight: {1, 1},
	bottom:      {1, 0},
	bottomLeft:  {1, -1},
	left:        {0, -1},
	topLeft:     {-1, -1},
}

var validConfiguration []map[int]rune = []map[int]rune{
	{
		topLeft:     'M',
		topRight:    'M',
		bottomLeft:  'S',
		bottomRight: 'S',
	},
	{
		topLeft:     'S',
		topRight:    'S',
		bottomLeft:  'M',
		bottomRight: 'M',
	},
	{
		topLeft:     'S',
		topRight:    'M',
		bottomLeft:  'S',
		bottomRight: 'M',
	},
	{
		topLeft:     'M',
		topRight:    'S',
		bottomLeft:  'M',
		bottomRight: 'S',
	},
}

func isExpectedLetter(matrix [][]rune, location [2]int, expectedLetter rune) bool {
	y, x := location[0], location[1]
	if y < 0 || x < 0 || y > len(matrix)-1 || x > len(matrix[y])-1 {
		return false
	}

	return expectedLetter == matrix[y][x]
}

func walkTwo(wg *sync.WaitGroup, foundCh chan interface{}, matrix [][]rune, location [2]int) {
	defer wg.Done()
	ok := isExpectedLetter(matrix, location, 'A')
	if !ok {
		return
	}

	y, x := location[0], location[1]
	for _, configuration := range validConfiguration {
		totalCorrect := 0
		for direction, expectedLetter := range configuration {
			nextY, nextX := y+directions[direction][0], x+directions[direction][1]
			ok := isExpectedLetter(matrix, [2]int{nextY, nextX}, expectedLetter)
			if !ok {
				continue
			}
			totalCorrect++
		}
		if totalCorrect != 4 {
			continue
		}
		foundCh <- struct{}{}
		return
	}
}

const searchWord = "XMAS"

func walkOne(wg *sync.WaitGroup, foundCh chan interface{}, matrix [][]rune, location [2]int, direction int, correctLetter int) {
	defer wg.Done()
	ok := isExpectedLetter(matrix, location, rune(searchWord[correctLetter]))
	if !ok {
		return
	}

	if correctLetter == len(searchWord)-1 {
		foundCh <- struct{}{}
		return
	}

	y, x := location[0], location[1]
	nextY, nextX := y+directions[direction][0], x+directions[direction][1]
	wg.Add(1)
	walkOne(wg, foundCh, matrix, [2]int{nextY, nextX}, direction, correctLetter+1)
}

func (p *puzzle) Part1() (string, error) {
	total := 0

	var wg sync.WaitGroup
	foundCh := make(chan interface{})
	doneCh := make(chan interface{})

	for y, row := range p.matrix {
		for x := range row {
			for direction := range directions {
				wg.Add(1)
				go walkOne(&wg, foundCh, p.matrix, [2]int{y, x}, direction, 0)
			}
		}
	}

	go func() {
		for {
			select {
			case <-foundCh:
				total++
			case <-doneCh:
				return
			}
		}
	}()

	wg.Wait()
	doneCh <- struct{}{}

	return fmt.Sprintf("%d", total), nil
}

func (p *puzzle) Part2() (string, error) {
	total := 0

	var wg sync.WaitGroup
	foundCh := make(chan interface{})
	doneCh := make(chan interface{})

	for y, row := range p.matrix {
		for x := range row {
			wg.Add(1)
			go walkTwo(&wg, foundCh, p.matrix, [2]int{y, x})
		}
	}

	go func() {
		for {
			select {
			case <-foundCh:
				total++
			case <-doneCh:
				return
			}
		}
	}()

	wg.Wait()
	doneCh <- struct{}{}

	return fmt.Sprintf("%d", total), nil
}
