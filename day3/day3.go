package day3

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type puzzle struct {
	commands []string
}

func NewSolution(filename string) *puzzle {
	p := &puzzle{
		commands: []string{},
	}

	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("[day3] error opening input file: %+v\n", err)
	}

	re, err := regexp.Compile(`(do\(\)|don't\(\)|mul\([0-9]+,[0-9]+\))`)
	if err != nil {
		log.Fatalf("[day3] failed to compile regexp: %+v\n", err)
	}

	validCommands := re.FindAllString(string(file), -1)
	p.commands = append(p.commands, validCommands...)

	return p
}

func parseMulCommand(command string) ([2]int, error) {
	values := strings.Split(strings.TrimSuffix(strings.TrimPrefix(command, "mul("), ")"), ",")
	c := [2]int{}
	for i, v := range values {
		value, err := strconv.Atoi(v)
		if err != nil {
			return [2]int{}, err
		}
		c[i] = value
	}

	if len(c) != 2 {
		return [2]int{}, fmt.Errorf("invalid mul() command")
	}
	return c, nil
}

func (p *puzzle) Part1() (string, error) {
	total := 0
	multipliers := [][2]int{}
	for _, command := range p.commands {
		if strings.Contains(command, "mul(") {
			m, err := parseMulCommand(command)
			if err != nil {
				return "", err
			}
			multipliers = append(multipliers, m)
		}
	}
	for _, m := range multipliers {
		total += m[0] * m[1]
	}

	return fmt.Sprintf("%d", total), nil
}

func (p *puzzle) Part2() (string, error) {
	total := 0
	enabled := true
	multipliers := [][2]int{}
	for _, command := range p.commands {
		if command == "do()" {
			enabled = true
		} else if command == "don't()" {
			enabled = false

		} else if enabled {
			m, err := parseMulCommand(command)
			if err != nil {
				log.Fatalf("Error parsing mul() command: %+v", err)
			}
			multipliers = append(multipliers, m)

		}
	}
	for _, m := range multipliers {
		total += m[0] * m[1]
	}

	return fmt.Sprintf("%d", total), nil
}
