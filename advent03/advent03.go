package advent03

import (
	"strconv"

	"advent2024/util"
)

func Solution(inputFile string) (part1, part2 any) {
	lines := util.ReadFile(inputFile)

	part1Jotalge := 0
	for _, line := range lines {
		currentMax := 0
		maxFirstChar := '0'
		for i, char := range line {
			if i >= len(line)-1 {
				break
			}
			if char < maxFirstChar {
				continue
			}
			maxFirstChar = char
			maxSecondChar := '0'
			for _, char2 := range line[i+1:] {
				if char2 < maxSecondChar {
					continue
				}
				maxSecondChar = char2
			}
			v, _ := strconv.Atoi(string(maxFirstChar) + string(maxSecondChar))
			if v > currentMax {
				currentMax = v
			}
		}
		part1Jotalge += currentMax
	}

	return part1Jotalge, 0
}
