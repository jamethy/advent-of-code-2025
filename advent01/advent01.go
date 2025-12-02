package advent01

import (
	"fmt"
	"strconv"

	"advent2024/util"
)

func Solution(inputFile string) (part1, part2 any) {
	lines := util.ReadFile(inputFile)

	location := 50
	zeroCount := 0
	zeroCountPart2 := 0

	for _, line := range lines {
		dir := line[0]
		num, _ := strconv.Atoi(line[1:])
		if dir == 'L' {
			if location == 0 && num != 0 {
				zeroCountPart2--
			}
			location -= num

			for location < 0 {
				location += 100
				zeroCountPart2++
			}
			if location == 0 {
				zeroCountPart2++
			}
		} else {
			location += num
			for location >= 100 {
				location -= 100
				zeroCountPart2++
			}
		}
		if location == 0 {
			zeroCount++
		}
	}

	return zeroCount, zeroCountPart2
}
