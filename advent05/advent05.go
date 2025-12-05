package advent05

import (
	"strconv"
	"strings"

	"advent2024/util"
)

func Solution(inputFile string) (part1, part2 any) {
	chunks := util.ReadFileSplitBy(inputFile, "\n\n")
	rangeLines := strings.Split(chunks[0], "\n")
	ingredientLines := strings.Split(chunks[1], "\n")

	ranges := make([][]int, len(rangeLines))
	for i := range rangeLines {
		ranges[i] = util.ParseIntList(rangeLines[i], "-")
	}

	part1Count := 0
	for _, ing := range ingredientLines {
		ingNum, _ := strconv.Atoi(ing)
		for _, rng := range ranges {
			if ingNum >= rng[0] && ingNum <= rng[1] {
				part1Count++
				break
			}
		}
	}

	return part1Count, 0
}
