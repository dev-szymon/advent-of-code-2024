package day5

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type puzzle struct {
	rules  [][2]int
	manual [][]int
}

func NewSolution(filename string) *puzzle {
	p := &puzzle{
		rules:  [][2]int{},
		manual: [][]int{},
	}

	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("[day5] error opening input file: %+v\n", err)
	}

	parts := strings.Split(string(file), "\n\n")
	rules := strings.Fields(parts[0])
	for _, r := range rules {
		rule := [2]int{}
		numbers := strings.Split(r, "|")
		for i, number := range numbers {
			if i == 2 {
				log.Fatal("[day5] expected 2 numbers in a rule")
			}
			no, err := strconv.Atoi(number)
			if err != nil {
				log.Fatalf("[day5] error converting rule to number: %+v\n", err)
			}
			rule[i] = no
		}
		p.rules = append(p.rules, rule)
	}

	for _, printedPages := range strings.Fields(parts[1]) {
		pages := []int{}
		for _, p := range strings.Split(printedPages, ",") {
			page, err := strconv.Atoi(p)
			if err != nil {
				log.Fatalf("[day5] error converting rule to number: %+v\n", err)
			}
			pages = append(pages, page)
		}
		p.manual = append(p.manual, pages)
	}

	return p
}

func (p *puzzle) Part1() (string, error) {
	total := 0
	for _, pages := range p.manual {
		order := map[int]int{}
		for i, page := range pages {
			order[page] = i
		}

		valid := true
		for _, rule := range p.rules {
			lo := rule[0]
			hi := rule[1]

			loIndex, loOk := order[lo]
			hiIndex, hiOk := order[hi]
			if !loOk || !hiOk {
				continue
			}
			if loIndex > hiIndex {
				valid = false
			}
		}
		if valid {
			total += pages[len(pages)/2]
		}
	}

	return fmt.Sprintf("%d", total), nil
}

func (p *puzzle) Part2() (string, error) {
	total := 0
	for _, pages := range p.manual {
		order := map[int]int{}
		for i, page := range pages {
			order[page] = i
		}

		valid := true
		for _, rule := range p.rules {
			lo := rule[0]
			hi := rule[1]

			loIndex, loOk := order[lo]
			hiIndex, hiOk := order[hi]
			if !loOk || !hiOk {
				continue
			}
			if loIndex > hiIndex {
				valid = false
			}
		}

		slices.SortFunc(pages, func(a, b int) int {
			for _, rule := range p.rules {
				lo := rule[0]
				hi := rule[1]

				if lo == a && hi == b {
					return -1
				} else if lo == b && hi == a {
					return 1
				} else {
					continue
				}
			}
			return 0
		})
		if !valid {
			total += pages[len(pages)/2]
		}
	}

	return fmt.Sprintf("%d", total), nil
}
