package day05

import (
	"strconv"
	"strings"
)

func part1(lines []string) (result int, err error) {
	freshIDRanges := [][2]int{}
	i := 0

	// Parsing fresh ID ranges
	for ; len(lines[i]) != 0; i++ {
		idRangeStrings := strings.Split(lines[i], "-")
		idRange := [2]int{}
		for j := range idRange {
			idRange[j], err = strconv.Atoi(idRangeStrings[j])
			if err != nil {
				return result, err
			}
		}

		freshIDRanges = append(freshIDRanges, idRange)
	}

	// Skipping the empty line
	i++

	// Checking the IDs
	for ; i < len(lines); i++ {
		id, err := strconv.Atoi(lines[i])
		if err != nil {
			return result, err
		}

		for _, idRange := range freshIDRanges {
			if id < idRange[0] || id > idRange[1] {
				continue
			}

			result++
			break
		}
	}

	return result, err
}

func part2(lines []string) (result int, err error) {
	idRanges := [][2]int{}

	// Parsing fresh ID ranges
	for i := 0; len(lines[i]) != 0; i++ {
		newRangeStrings := strings.Split(lines[i], "-")
		newRange := [2]int{}
		for j := range newRange {
			newRange[j], err = strconv.Atoi(newRangeStrings[j])
			if err != nil {
				return result, err
			}
		}

		// Checking if the new range overlaps with any of the existing ranges
		// And shifting IDs to avoid overlaps
		for j, existingRange := range idRanges {
			// If the new range completely covers an existing range
			// we need to remove that existing range.
			// It's a dirty/hacky way of doing it, but this will result in adding 0 to the total count
			if newRange[0] < existingRange[0] && newRange[1] > existingRange[1] {
				idRanges[j][0] = 0
				idRanges[j][1] = -1
				continue
			}

			// If lower bound of the new range overlaps with an existing range
			// we need to bump it up to be above the upper bound of the existing range
			if newRange[0] >= existingRange[0] && newRange[0] <= existingRange[1] {
				newRange[0] = existingRange[1] + 1
			}

			// If upper bound of the new range overlaps with an existing range
			// we need to bump it down to be below the lower bound of the existing range
			if newRange[1] >= existingRange[0] && newRange[1] <= existingRange[1] {
				newRange[1] = existingRange[0] - 1
			}
		}

		// if the new range is inverted it means that it's completely covered by existing ranges
		// and should be ignored
		if newRange[1]-newRange[0] < 0 {
			continue
		}

		idRanges = append(idRanges, newRange)
	}

	for _, idRange := range idRanges {
		result += idRange[1] - idRange[0] + 1
	}

	return result, err
}

func move(pos Coords, dir Coords) Coords {
	return Coords{pos.x + dir.x, pos.y + dir.y}
}

func withinBounds(lines [][]rune, pos Coords) bool {
	if pos.x < 0 || pos.x >= len(lines[0]) {
		return false
	}

	if pos.y < 0 || pos.y >= len(lines) {
		return false
	}

	return true
}

type Coords struct {
	x int
	y int
}

var directions = []Coords{
	{-1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
	{1, 0},
	{1, -1},
	{0, -1},
	{-1, -1},
}
