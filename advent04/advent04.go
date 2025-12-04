package advent04

import "advent2024/util"

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

	reachableCount := 0
	for i := range lines {
		for j, v := range grid[i] {
			if !v {
				continue
			}

			nearBy := 0
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
					}
				}
			}

			if nearBy < 4 {
				reachableCount++
			}
		}
	}

	return reachableCount, 0
}
