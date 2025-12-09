package advent09

import (
	"advent2024/util"
	"advent2024/util/mathutil"
)

type Point struct {
	x, y int
}

func Solution(inputFile string) (part1, part2 any) {
	lines := util.ReadFile(inputFile)
	points := make([]Point, len(lines))
	for i, line := range lines {
		nums := util.ParseIntList(line, ",")
		points[i] = Point{nums[0], nums[1]}
	}

	biggestArea := 0
	for i, a := range points {
		for j := i + 1; j < len(points); j++ {
			b := points[j]
			area := (mathutil.AbsInt(a.x-b.x) + 1) * (mathutil.AbsInt(a.y-b.y) + 1)
			if area > biggestArea {
				biggestArea = area
			}
		}
	}

	return biggestArea, 0
}
