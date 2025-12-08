package advent08

import (
	"cmp"
	"maps"
	"slices"

	"advent2024/util"
	"advent2024/util/mathutil"
)

type Point struct {
	x, y, z int
}

func (pt Point) DistanceSquared(p Point) int {
	return mathutil.IntPow(pt.x-p.x, 2) + mathutil.IntPow(pt.y-p.y, 2) + mathutil.IntPow(pt.z-p.z, 2)
}

type Connection struct {
	p1, p2   Point
	distance int
}

func Solution(inputFile string, iterations int) (part1, part2 any) {
	lines := util.ReadFile(inputFile)

	points := make([]Point, len(lines))
	for i, line := range lines {
		nums := util.ParseIntList(line, ",")
		points[i] = Point{nums[0], nums[1], nums[2]}
	}

	connections := make([]Connection, 0, len(points)*len(points)/2)
	for i, a := range points {
		for j, b := range points {
			if i == j {
				break
			}
			connections = append(connections, Connection{
				p1:       a,
				p2:       b,
				distance: a.DistanceSquared(b),
			})
		}
	}

	slices.SortFunc(connections, func(a, b Connection) int {
		return cmp.Compare(a.distance, b.distance)
	})

	circuits := make(map[Point]int)
	circuit := 1

	joinCircuit := func(conn Connection) {
		c1, c1ok := circuits[conn.p1]
		c2, c2ok := circuits[conn.p2]
		if c1ok && c2ok {
			// move them
			for p, c := range circuits {
				if c == c2 {
					circuits[p] = c1
				}
			}
		} else if c1ok {
			circuits[conn.p2] = c1
		} else if c2ok {
			circuits[conn.p1] = c2
		} else {
			circuits[conn.p1] = circuit
			circuits[conn.p2] = circuit
			circuit++
		}
	}

	for i, conn := range connections {
		if i == iterations {
			break
		}
		joinCircuit(conn)
	}

	circuitCounts := make(map[int]int)
	for _, c := range circuits {
		circuitCounts[c]++
	}
	ordered := slices.SortedFunc(maps.Values(circuitCounts), func(a int, b int) int {
		return -cmp.Compare(a, b)
	})
	part1 = ordered[0] * ordered[1] * ordered[2]

	for i := iterations; i < len(connections); i++ {
		conn := connections[i]
		joinCircuit(conn)

		if len(circuits) < len(points) {
			continue
		}

		uniqueCircuits := make(map[int]struct{})
		for _, c := range circuits {
			uniqueCircuits[c] = struct{}{}
			if len(uniqueCircuits) > 1 {
				break
			}
		}
		if len(uniqueCircuits) == 1 {
			part2 = conn.p1.x * conn.p2.x
			break
		}
	}

	return part1, part2
}
