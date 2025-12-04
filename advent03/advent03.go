package advent03

import (
	"advent2024/util"
	"advent2024/util/mathutil"
)

func Solution(inputFile string) (part1, part2 any) {
	lines := util.ReadFile(inputFile)

	part1Joltage, part2Joltage := 0, 0
	for _, line := range lines {
		nums := util.ParseIntList(line, "")

		part1ForLine := largestRemainingNum(nums, 2)
		part1Joltage += part1ForLine

		part2ForLine := largestRemainingNum(nums, 12)
		part2Joltage += part2ForLine
	}

	return part1Joltage, part2Joltage
}

func largestRemainingNum(nums []int, size int) int {
	if size == 0 {
		return 0
	}
	maxNum := 0
	for i, n := range nums {
		if i > len(nums)-size {
			break
		}
		if n > maxNum {
			maxNum = n
		}
	}
	maxValue := 0
	for i, n := range nums {
		if i > len(nums)-size {
			break
		}
		if n != maxNum {
			continue
		}
		maxNum = n
		maxRemaining := largestRemainingNum(nums[i+1:], size-1)
		potentialMax := maxNum*mathutil.IntPow(10, size-1) + maxRemaining
		if potentialMax > maxValue {
			maxValue = potentialMax
		}
	}
	return maxValue
}
