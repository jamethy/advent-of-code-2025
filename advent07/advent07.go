package advent07

import (
	"maps"
	"strings"

	"advent2024/util"
)

func Solution(inputFile string) (part1, part2 any) {
	lines := util.ReadFile(inputFile)
	startingLine := lines[0]
	lines = lines[1:]

	beams := make(map[int]int)
	beams[strings.Index(startingLine, "S")] = 1

	splits := 0
	for _, line := range lines {
		previousBeams := maps.Clone(beams)
		for beam, cnt := range previousBeams {
			c := line[beam]
			if c == '^' {
				splits++
				delete(beams, beam)
				beams[beam-1] += cnt
				beams[beam+1] += cnt
			}
		}
	}

	totalBeams := 0
	for _, cnt := range beams {
		totalBeams += cnt
	}

	return splits, totalBeams
}
