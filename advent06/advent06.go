package advent06

import (
	"strconv"
	"strings"

	"advent2024/util"
	"advent2024/util/mathutil"
)

func Solution(inputFile string) (part1, part2 any) {
	lines := util.ReadFile(inputFile)
	part1Total := calculateSum1(lines)

	maxLength := 0
	for _, line := range lines {
		maxLength = mathutil.MaxInt(maxLength, len(line))
	}

	part2Lines := make([]string, maxLength)
	for _, line := range lines {
		for j, c := range line {
			part2Lines[j] += strings.TrimSpace(string(c))
		}
	}

	part2Str := strings.Join(part2Lines, "\n")
	problems := strings.Split(part2Str, "\n\n")
	part2Total := 0
	for _, prob := range problems {
		probLines := strings.Split(prob, "\n")
		op := probLines[0][len(probLines[0])-1]
		probLines[0] = probLines[0][:len(probLines[0])-1]
		total := 0
		if op == '*' {
			total = 1
		}
		for _, l := range probLines {
			n, _ := strconv.Atoi(l)
			if op == '+' {
				total += n
			} else {
				total *= n
			}
		}
		part2Total += total
	}

	return part1Total, part2Total
}

func calculateSum1(lines []string) int {
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

	total := 0
	for col, op := range operators {
		colTotal := 0
		if op == "*" {
			colTotal = 1
		}
		for _, row := range grid {
			if col >= len(row) {
				continue
			}
			if op == "+" {
				colTotal += row[col]
			} else {
				colTotal *= row[col]
			}
		}
		total += colTotal
	}

	return total
}
