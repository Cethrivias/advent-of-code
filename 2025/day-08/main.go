package day08

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

type Box struct {
	x         float64
	y         float64
	z         float64
	circuitID int
}

type Wire struct {
	boxA     *Box
	boxB     *Box
	distance float64
}

func distance(a Box, b Box) float64 {
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2) + math.Pow(a.z-b.z, 2))
}

func part1(lines []string, connections int) (result int, err error) {
	boxes := []*Box{}

	for _, line := range lines {
		coords := strings.Split(line, ",")
		x, err := strconv.ParseFloat(coords[0], 64)
		if err != nil {
			return result, err
		}
		y, err := strconv.ParseFloat(coords[1], 64)
		if err != nil {
			return result, err
		}
		z, err := strconv.ParseFloat(coords[2], 64)
		if err != nil {
			return result, err
		}

		boxes = append(boxes, &Box{x: x, y: y, z: z})
	}

	wires := []Wire{}
	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			wire := Wire{boxA: boxes[i], boxB: boxes[j], distance: distance(*boxes[i], *boxes[j])}
			wires = append(wires, wire)
		}
	}

	slices.SortFunc(wires, func(a, b Wire) int {
		return int(a.distance - b.distance)
	})

	circuitID := 1
	for i := 0; i < connections; i++ {
		if wires[i].boxA.circuitID == 0 && wires[i].boxB.circuitID == 0 {
			wires[i].boxA.circuitID = circuitID
			wires[i].boxB.circuitID = circuitID
			circuitID++
			continue
		}

		if wires[i].boxA.circuitID != 0 && wires[i].boxB.circuitID != 0 {
			tmpID := wires[i].boxB.circuitID

			for _, vox := range boxes {
				if vox.circuitID == tmpID {
					vox.circuitID = wires[i].boxA.circuitID
				}
			}
			continue
		}

		if wires[i].boxA.circuitID == 0 {
			wires[i].boxA.circuitID = wires[i].boxB.circuitID

			continue
		}

		if wires[i].boxB.circuitID == 0 {
			wires[i].boxB.circuitID = wires[i].boxA.circuitID

			continue
		}
	}

	circuits := make([]int, circuitID, circuitID)

	for _, box := range boxes {
		if box.circuitID == 0 {
			continue
		}
		circuits[box.circuitID]++
	}

	slices.SortFunc(circuits, func(a int, b int) int {
		return b - a
	})

	result = circuits[0] * circuits[1] * circuits[2]

	return result, err
}

func part2(lines []string) (result int, err error) {
	boxes := []*Box{}

	// Parsing input into boxes
	for _, line := range lines {
		coords := strings.Split(line, ",")
		x, err := strconv.ParseFloat(coords[0], 64)
		if err != nil {
			return result, err
		}
		y, err := strconv.ParseFloat(coords[1], 64)
		if err != nil {
			return result, err
		}
		z, err := strconv.ParseFloat(coords[2], 64)
		if err != nil {
			return result, err
		}

		// MaxInt means that the box is not a part of any circuit
		boxes = append(boxes, &Box{x: x, y: y, z: z, circuitID: math.MaxInt})
	}

	// All possible wire combinations
	wires := []Wire{}
	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			wire := Wire{boxA: boxes[i], boxB: boxes[j], distance: distance(*boxes[i], *boxes[j])}
			wires = append(wires, wire)
		}
	}

	// Sorted from shortest to longest
	slices.SortFunc(wires, func(a, b Wire) int {
		return int(a.distance - b.distance)
	})

	// Connecting boxes into circuits
	circuits := [][]*Box{}
	for i := 0; true; i++ {
		// Boxes already in the same circuit
		if wires[i].boxA.circuitID == wires[i].boxB.circuitID && wires[i].boxA.circuitID != math.MaxInt {
			continue
		}

		// Both boxes are not a part of any circuit
		if wires[i].boxA.circuitID == math.MaxInt && wires[i].boxB.circuitID == math.MaxInt {
			// Starting a new circuit. Using the slice index as a `circuitID`
			circuit := []*Box{wires[i].boxA, wires[i].boxB}
			circuits = append(circuits, circuit)
			wires[i].boxA.circuitID = len(circuits) - 1
			wires[i].boxB.circuitID = len(circuits) - 1
			continue
		}

		// Both boxes belong to different circuits.
		// Moving all boxes from the second circuit into the first circuit
		if wires[i].boxA.circuitID != math.MaxInt && wires[i].boxB.circuitID != math.MaxInt {
			circuitAID := wires[i].boxA.circuitID
			circuitBID := wires[i].boxB.circuitID

			circuitA := circuits[circuitAID]
			circuitB := circuits[circuitBID]

			for _, box := range circuitB {
				box.circuitID = wires[i].boxA.circuitID
				circuitA = append(circuitA, box)
			}
			circuits[circuitAID] = circuitA
			circuits[circuitBID] = []*Box{}

			// If the first circuit contains all the boxes, compute the result
			if len(circuitA) == len(boxes) {
				return int(wires[i].boxA.x) * int(wires[i].boxB.x), err
			}

			continue
		}

		// If only the first box does not belong to a circuit add it to the circuit of the second box
		if wires[i].boxA.circuitID == math.MaxInt {
			wires[i].boxA.circuitID = wires[i].boxB.circuitID
			circuits[wires[i].boxB.circuitID] = append(circuits[wires[i].boxB.circuitID], wires[i].boxA)

			// If the circuit contains all the boxes, compute the result
			if len(circuits[wires[i].boxB.circuitID]) == len(boxes) {
				return int(wires[i].boxA.x) * int(wires[i].boxB.x), err
			}

			continue
		}

		// If only the second box does not belong to a circuit add it to the circuit of the first box
		if wires[i].boxB.circuitID == math.MaxInt {
			wires[i].boxB.circuitID = wires[i].boxA.circuitID
			circuits[wires[i].boxA.circuitID] = append(circuits[wires[i].boxA.circuitID], wires[i].boxB)

			// If the circuit contains all the boxes, compute the result
			if len(circuits[wires[i].boxA.circuitID]) == len(boxes) {
				return int(wires[i].boxA.x) * int(wires[i].boxB.x), err
			}

			continue
		}
	}

	return result, err
}
