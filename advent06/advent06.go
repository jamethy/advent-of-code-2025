package advent06

import (
	"strings"

	"advent2024/util"
)

func Solution(inputFile string) (part1, part2 any) {
	lines := util.ReadFile(inputFile)
	grid := make([][]int, len(lines)-1)
	operators := make([]string, 0)
	for i, line := range lines {
		if i == len(lines)-1 {
			parts := strings.Split(line, " ")
			for _, p := range parts {
				if p != "" {
					operators = append(operators, p)
				}
			}
			break
		}
		grid[i] = util.ParseIntList(line, " ")
	}

	part1Total := 0
	for col, op := range operators {
		colTotal := 0
		if op == "*" {
			colTotal = 1
		}
		for _, row := range grid {
			if op == "+" {
				colTotal += row[col]
			} else {
				colTotal *= row[col]
			}
		}
		part1Total += colTotal
	}

	return part1Total, 0
}
