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

	beams := make(map[int]bool)
	beams[strings.Index(startingLine, "S")] = true

	splits := 0
	for _, line := range lines {
		previousBeams := maps.Clone(beams)
		for beam := range previousBeams {
			c := line[beam]
			if c == '^' {
				splits++
				delete(beams, beam)
				beams[beam-1] = true
				beams[beam+1] = true
			}
		}
	}

	return splits, 0
}
