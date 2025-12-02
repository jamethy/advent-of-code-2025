package advent02

import (
	"strconv"
	"strings"

	"advent2024/util"
)

func Solution(inputFile string) (part1, part2 any) {
	lines := util.ReadFile(inputFile)

	part1Total := 0
	part2Total := 0
	for _, r := range strings.Split(lines[0], ",") {
		numRange := util.ParseIntList(r, "-")

		for n := numRange[0]; n <= numRange[1]; n++ {
			s := strconv.Itoa(n)
			if s[0:len(s)/2] == s[len(s)/2:] {
				part1Total += n
				part2Total += n
				continue
			}
			for i := 1; i < (len(s)+1)/2; i++ {
				if len(s)%i != 0 {
					continue
				}
				sub := s[0:i]
				if strings.Repeat(sub, len(s)/i-1) == s[i:] {
					part2Total += n
					break
				}
			}
		}
	}

	return part1Total, part2Total
}
