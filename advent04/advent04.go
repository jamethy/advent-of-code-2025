package advent04

import (
	"advent2024/util"
)

type Point struct {
	row, col int
}

func Solution(inputFile string) (part1, part2 any) {
	lines := util.ReadFile(inputFile)
	width := len(lines[0])

	grid := make([][]bool, len(lines))
	for i, line := range lines {
		gridLine := make([]bool, width)
		for j, cell := range line {
			if cell == '@' {
				gridLine[j] = true
			}
		}
		grid[i] = gridLine
	}

	part1Count := 0
	part2Count := 0
	reachable := make(map[Point]struct{})
	for {
		for i := range lines {
			for j, v := range grid[i] {
				if !v {
					continue
				}

				nearBy := 0
			nearByLoop:
				for ii := i - 1; ii <= i+1; ii++ {
					for jj := j - 1; jj <= j+1; jj++ {
						if ii == i && jj == j {
							continue
						}
						if ii < 0 || jj < 0 || ii >= len(grid) || jj >= width {
							continue
						}
						if grid[ii][jj] {
							nearBy++
							if nearBy > 3 {
								break nearByLoop
							}
						}
					}
				}

				if nearBy < 4 {
					reachable[Point{i, j}] = struct{}{}
				}
			}
		}

		if part1Count == 0 {
			part1Count = len(reachable)
		}
		part2Count += len(reachable)
		if len(reachable) == 0 {
			break
		}
		for p := range reachable {
			grid[p.row][p.col] = false
		}
		reachable = make(map[Point]struct{})
	}

	return part1Count, part2Count
}
