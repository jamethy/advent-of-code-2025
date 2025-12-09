package advent09

import (
	"cmp"
	"math"
	"slices"

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

	part1BiggestArea := 0
	for i, a := range points {
		for j := i + 1; j < len(points); j++ {
			b := points[j]
			area := (mathutil.AbsInt(a.x-b.x) + 1) * (mathutil.AbsInt(a.y-b.y) + 1)
			if area > part1BiggestArea {
				part1BiggestArea = area
			}
		}
	}

	//part2BiggestArea := 0
	//for i, a := range points {
	//	for j := i + 1; j < len(points); j++ {
	//		b := points[j]
	//		area := (mathutil.AbsInt(a.x-b.x) + 1) * (mathutil.AbsInt(a.y-b.y) + 1)
	//		if area < part2BiggestArea {
	//			continue
	//		}
	//
	//		noPointsInside := true
	//		for _, c := range points {
	//			if isInside(a, b, c) {
	//				noPointsInside = false
	//				break
	//			}
	//		}
	//
	//		if noPointsInside {
	//			part2BiggestArea = area
	//		}
	//	}
	//}

	horizontalLines := make([][]Point, 0, len(points))
	verticalLines := make([][]Point, 0, len(points))
	for li, pi := range points {
		lj := li + 1
		if lj >= len(points) {
			lj = 0
		}
		pj := points[lj]
		if pj.x == pi.x {
			verticalLines = append(verticalLines, []Point{pi, pj})
		} else {
			horizontalLines = append(horizontalLines, []Point{pi, pj})
		}
	}
	slices.SortFunc(verticalLines, func(a, b []Point) int {
		return cmp.Compare(a[0].x, b[0].x)
	})
	slices.SortFunc(horizontalLines, func(a, b []Point) int {
		return cmp.Compare(a[0].y, b[0].y)
	})

	//a.x == 9 && a.y == 5 && b.x == 2 && b.y == 3
	part2BiggestArea := 0 // 9,5 and 2,3
	for i, a := range points {
		for j := i + 1; j < len(points); j++ {
			b := points[j]
			area := (mathutil.AbsInt(a.x-b.x) + 1) * (mathutil.AbsInt(a.y-b.y) + 1)
			if area < part2BiggestArea {
				continue
			}

			left := mathutil.MinInt(a.x, b.x)
			right := mathutil.MaxInt(a.x, b.x)
			top := mathutil.MinInt(a.y, b.y)
			bottom := mathutil.MaxInt(a.y, b.y)

			pointsInside := false
			for _, c := range points {
				if c.x <= left {
					continue
				}
				if c.x >= right {
					continue
				}
				if c.y <= top {
					continue
				}
				if c.y >= bottom {
					continue
				}
				pointsInside = true
				break
			}
			if pointsInside {
				continue
			}

			cutByLine := false
			for _, hl := range horizontalLines {
				if hl[0].y <= top || hl[0].y >= bottom {
					continue
				}
				lineLeft := mathutil.MinInt(hl[0].x, hl[1].x)
				if lineLeft > left {
					continue
				}
				lineRight := mathutil.MaxInt(hl[0].x, hl[1].x)
				if lineRight < right {
					continue
				}
				cutByLine = true
				break
			}
			if cutByLine {
				continue
			}

			for _, vl := range verticalLines {
				if vl[0].x <= left || vl[0].x >= right {
					continue
				}
				lineTop := mathutil.MinInt(vl[0].y, vl[1].y)
				if lineTop < top {
					continue
				}
				lineBottom := mathutil.MaxInt(vl[0].y, vl[1].y)
				if lineBottom > bottom {
					continue
				}
				cutByLine = true
				break
			}
			if cutByLine {
				continue
			}

			otherCorners := []Point{
				{a.x, b.y},
				{b.x, a.y},
			}
			otherCornersGood := true
			for _, o := range otherCorners {

				var closestLine []Point
				closestDist := math.MaxInt
				for _, hl := range horizontalLines {
					dist := mathutil.AbsInt(o.y - hl[0].y)
					if dist >= closestDist {
						break // since they are sorted
					}
					if !isBetween(o.x, hl[0].x, hl[1].x) {
						continue
					}
					closestLine = hl
					closestDist = dist
					if dist == 0 {
						break
					}
				}

				if closestDist == 0 {
					continue
				}
				if !isRightOfLine(o, closestLine) {
					otherCornersGood = false
					break
				}

				closestDist = math.MaxInt
				for _, vl := range verticalLines {
					dist := mathutil.AbsInt(o.x - vl[0].x)
					if dist >= closestDist {
						break // since they are sorted
					}
					if !isBetween(o.y, vl[0].y, vl[1].y) {
						continue
					}
					closestLine = vl
					if dist == 0 {
						break
					}
				}

				if closestDist == 0 {
					continue
				}
				if !isRightOfLine(o, closestLine) {
					otherCornersGood = false
					break
				}
			}

			if otherCornersGood {
				part2BiggestArea = area
			}
		}
	}

	return part1BiggestArea, part2BiggestArea
}

func isBetween(v, a, b int) bool {
	if a > b {
		a, b = b, a
	}
	return v >= a && v <= b
}

func isRightOfLine(a Point, line []Point) bool {
	delta := Point{
		x: a.x - line[0].x,
		y: a.y - line[0].y,
	}
	lineRotated90 := Point{
		x: line[0].y - line[1].y,
		y: line[1].x - line[0].x,
	}
	return delta.x*lineRotated90.x+delta.y*lineRotated90.y > 0
}

func isInside(left, right, top, bottom int, c Point) bool {
	if c.x <= left {
		return false
	}
	if c.x >= right {
		return false
	}
	if c.y <= top {
		return false
	}
	if c.y >= bottom {
		return false
	}
	return true
}
