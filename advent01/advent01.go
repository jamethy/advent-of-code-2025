package advent01

import (
	"strconv"

	"advent2024/util"
)

func Solution(inputFile string) (part1, part2 any) {
	lines := util.ReadFile(inputFile)

	location := 50
	zeroCount := 0

	for _, line := range lines {
		dir := line[0]
		num, _ := strconv.Atoi(line[1:])
		if dir == 'L' {
			location -= num
		} else {
			location += num
		}
		for location < 0 {
			location += 100
		}
		location = location % 100
		if location == 0 {
			zeroCount++
		}
	}

	return zeroCount, 0
}
