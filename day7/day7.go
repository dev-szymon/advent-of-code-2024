package day7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dev-szymon/advent-of-code/utils"
)

type puzzle struct {
	equations []equation
}

type equation struct {
	output   int
	operands []int
}

func calculatePermutations(values []int, outputs []int) []int {
	if len(values) == 0 {
		return outputs
	}

	if len(outputs) == 0 {
		return calculatePermutations(values[1:], []int{values[0]})
	}

	nextOutputs := []int{}
	for _, o := range outputs {
		nextOutputs = append(nextOutputs, o*values[0])
		nextOutputs = append(nextOutputs, o+values[0])
	}
	return calculatePermutations(values[1:], nextOutputs)
}

func withConcatenation(values []int, outputs []int) []int {
	if len(values) == 0 {
		return outputs
	}

	if len(outputs) == 0 {
		c := fmt.Sprintf("%d%d", values[0], values[1])
		n := utils.MustMakeInt(c)
		return withConcatenation(values[2:], []int{values[0] * values[1], values[0] + values[1], n})
	}

	nextOutputs := []int{}
	for _, o := range outputs {
		nextOutputs = append(nextOutputs, o*values[0])
		nextOutputs = append(nextOutputs, o+values[0])

		concatenatedValue := fmt.Sprintf("%d%d", o, values[0])
		n := utils.MustMakeInt(concatenatedValue)
		nextOutputs = append(nextOutputs, n)
	}
	return withConcatenation(values[1:], nextOutputs)
}

func NewSolution(filename string) *puzzle {
	p := &puzzle{
		equations: []equation{},
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("[day7] error opening input file: %+v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")
		e := equation{
			output:   utils.MustMakeInt(parts[0]),
			operands: []int{},
		}
		for _, o := range strings.Fields(parts[1]) {
			e.operands = append(e.operands, utils.MustMakeInt(o))
		}

		p.equations = append(p.equations, e)
	}

	return p
}

func (p *puzzle) Part1() (string, error) {
	total := 0
	for _, e := range p.equations {
		permutations := calculatePermutations(e.operands, []int{})
		for _, p := range permutations {
			if p == e.output {
				total += e.output
				break
			}
		}
	}

	return fmt.Sprintf("%d", total), nil
}

func (p *puzzle) Part2() (string, error) {
	total := 0
	for _, e := range p.equations {
		permutations := withConcatenation(e.operands, []int{})
		for _, p := range permutations {
			if p == e.output {
				total += e.output
				break
			}
		}
	}

	return fmt.Sprintf("%d", total), nil
}
