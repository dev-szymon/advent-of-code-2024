package main

import (
	"fmt"
	"log"

	"github.com/dev-szymon/advent-of-code/day1"
	"github.com/dev-szymon/advent-of-code/day2"
	"github.com/dev-szymon/advent-of-code/day3"
	"github.com/dev-szymon/advent-of-code/day4"
	"github.com/dev-szymon/advent-of-code/day5"
	"github.com/dev-szymon/advent-of-code/day6"
)

type Solution interface {
	Part1() (string, error)
	Part2() (string, error)
}

func main() {
	solutions := map[int]Solution{
		1: day1.NewSolution("day1/input.txt"),
		2: day2.NewSolution("day2/input.txt"),
		3: day3.NewSolution("day3/input.txt"),
		4: day4.NewSolution("day4/input.txt"),
		5: day5.NewSolution("day5/input.txt"),
		6: day6.NewSolution("day6/input_test.txt"),
	}

	for day := 1; day <= len(solutions); day++ {
		puzzle := solutions[day]
		part1, err := puzzle.Part1()
		if err != nil {
			log.Fatalf("error solving day %d part 1: %+v\n", day, err)
		}
		part2, err := puzzle.Part2()
		if err != nil {
			log.Fatalf("error solving day %d part 2: %+v\n", day, err)
		}

		fmt.Printf("\n[day%d][part1] %s\n", day, part1)
		fmt.Printf("[day%d][part2] %s\n", day, part2)
	}

}
