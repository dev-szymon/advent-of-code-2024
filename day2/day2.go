package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type puzzle struct {
	reports [][]int
}

func NewSolution(filename string) *puzzle {
	p := &puzzle{
		reports: [][]int{},
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("[day2] error opening input file: %+v\n", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		report := []int{}
		for _, reportLevel := range strings.Fields(scanner.Text()) {
			level, err := strconv.Atoi(reportLevel)
			if err != nil {
				log.Fatalf("[day2] error converting location to int: %s\n", reportLevel)
			}
			report = append(report, level)
		}
		p.reports = append(p.reports, report)
	}

	return p
}

func isSafe(levels []int) bool {
	diffs := []int{}
	for i := 0; i < len(levels)-1; i++ {
		diffs = append(diffs, levels[i+1]-levels[i])
	}

	isNegative := diffs[0] < 0
	for _, diff := range diffs {
		if isNegative && (diff < -3 || diff > -1) || !isNegative && (diff > 3 || diff < 1) {
			return false
		}
	}
	return true
}

func (p *puzzle) Part1() (string, error) {
	safeReports := 0

	for _, report := range p.reports {
		safe := isSafe(report)

		if safe {
			safeReports++
		}
	}

	return fmt.Sprintf("%d", safeReports), nil
}

func (p *puzzle) Part2() (string, error) {
	safeReports := 0

	for _, report := range p.reports {
		safe := isSafe(report)

		if !safe {
			for i := 0; i < len(report); i++ {
				if i == len(report)-1 {
					removedLevel := []int{}
					removedLevel = append(removedLevel, report[:i]...)

					inSafe := isSafe(removedLevel)
					if inSafe {
						safe = true
					}
				} else if i == 0 {
					removedLevel := []int{}
					removedLevel = append(removedLevel, report[i+1:]...)

					inSafe := isSafe(removedLevel)
					if inSafe {
						safe = true
					}
				} else {
					removedLevel := []int{}
					removedLevel = append(removedLevel, report[:i]...)
					removedLevel = append(removedLevel, report[i+1:]...)

					inSafe := isSafe(removedLevel)
					if inSafe {
						safe = true
					}
				}
			}
		}

		if safe {
			safeReports++
		}
	}

	return fmt.Sprintf("%d", safeReports), nil
}
