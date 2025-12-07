package day06

import (
	"slices"
)

func part1(lines [][]rune) (result int, err error) {
	for x := 1; x < len(lines); x++ {
		for y := 0; y < len(lines[x]); y++ {
			if lines[x-1][y] != 'S' && lines[x-1][y] != '|' {
				continue
			}

			if lines[x][y] != '^' {
				lines[x][y] = '|'
				continue
			}

			result++

			lines[x][y-1] = '|'
			lines[x][y+1] = '|'
		}
	}

	return result, err
}

var cache [141][141]int = [141][141]int{}

func part2(lines [][]rune) (result int, err error) {
	y := slices.Index(lines[0], 'S')
	result = countTimelines(lines, 0, y)

	return result, err
}

func countTimelines(lines [][]rune, x int, y int) int {
	x_o := x
	y_o := y
	result := cache[x_o][y_o]
	if result != 0 {
		return result
	}

	for {
		if x == len(lines) {
			return 1
		}
		if lines[x][y] != '^' {
			x++
			continue
		}

		result = countTimelines(lines, x, y-1) + countTimelines(lines, x, y+1)
		cache[x_o][y_o] = result

		return result
	}
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
