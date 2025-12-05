package advent05

import (
	"cmp"
	"slices"
	"strconv"
	"strings"

	"advent2024/util"
	"advent2024/util/mathutil"
)

type Range struct {
	start   int
	end     int
	invalid bool
}

func (r Range) Contains(i int) bool {
	return i >= r.start && i <= r.end
}

func (r Range) Overlaps(other Range) bool {
	other.start--
	other.end++
	return r.Contains(other.start) || r.Contains(other.end) || other.Contains(r.start) || other.Contains(r.end)
}

func Solution(inputFile string) (part1, part2 any) {
	chunks := util.ReadFileSplitBy(inputFile, "\n\n")
	rangeLines := strings.Split(chunks[0], "\n")
	ingredientLines := strings.Split(chunks[1], "\n")

	ranges := make([]Range, 0)
	for i := range rangeLines {
		slices.SortFunc(ranges, func(a, b Range) int {
			return cmp.Compare(a.start, b.start)
		})

		nums := util.ParseIntList(rangeLines[i], "-")
		rng := Range{
			start: nums[0],
			end:   nums[1],
		}

		for otherI, other := range ranges {
			if other.invalid {
				continue
			}
			if rng.Overlaps(other) {
				ranges[otherI].invalid = true
				rng.start = mathutil.MinInt(rng.start, other.start)
				rng.end = mathutil.MaxInt(rng.end, other.end)
				continue
			}
		}
		ranges = append(ranges, rng)
	}

	part1Count := 0
	for _, ing := range ingredientLines {
		ingNum, _ := strconv.Atoi(ing)
		for _, rng := range ranges {
			if rng.invalid {
				continue
			}
			if rng.Contains(ingNum) {
				part1Count++
				break
			}
		}
	}

	part2Count := 0
	for _, rng := range ranges {
		if rng.invalid {
			continue
		}
		part2Count += rng.end - rng.start + 1
	}

	return part1Count, part2Count
}
