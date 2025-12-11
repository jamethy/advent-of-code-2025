package advent10

import (
	"strings"

	"advent2024/util"
	"advent2024/util/bitutil"
	"advent2024/util/mathutil"
)

func Solution(inputFile string) (part1, part2 any) {
	lines := util.ReadFile(inputFile)

	part1Sum := 0
	for _, line := range lines {
		chunks := strings.Split(line, " ")

		desiredLightStateChunk, chunks := chunks[0], chunks[1:]
		lightCount := uint(len(desiredLightStateChunk) - 2)
		desiredLightState := uint(0)
		for i, c := range desiredLightStateChunk {
			if c == '#' {
				desiredLightState = bitutil.SetBit(desiredLightState, uint(i-1), true)
			}
		}

		chunks = chunks[:len(chunks)-1] // get rid of joltage requirements
		buttons := make([]uint, 0, len(chunks))
		for _, c := range chunks {
			button := uint(0)
			nums := util.ParseIntList(c[1:len(c)-1], ",")
			for _, num := range nums {
				button = bitutil.SetBit(button, uint(num), true)
			}
			buttons = append(buttons, button)
		}

		options := uint(mathutil.IntPow(2, len(buttons)) - 1)
		minimumButtons := 1000
		for pressBits := uint(1); pressBits < options; pressBits++ {
			if pressButtons(buttons, pressBits, lightCount) == desiredLightState {
				buttonsPressed := 0
				for i := 0; i < len(buttons); i++ {
					if bitutil.IsBitSet(pressBits, uint(i)) {
						buttonsPressed++
					}
				}
				if buttonsPressed < minimumButtons {
					minimumButtons = buttonsPressed
				}
			}
		}

		part1Sum += minimumButtons
	}

	return part1Sum, 0
}

func pressButtons(buttons []uint, pressBits, lightCount uint) uint {
	value := uint(0)
	for i, button := range buttons {
		if !bitutil.IsBitSet(pressBits, uint(i)) {
			continue
		}
		value = bitutil.LimitedXOR(value, button, lightCount)
	}
	return value
}
