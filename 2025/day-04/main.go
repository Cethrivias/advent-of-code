package day04

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

func part1(lines [][]rune) (result int, err error) {
	cols := len(lines[0])
	rows := len(lines)
	for x := range cols {
		for y := range rows {
			if lines[x][y] != '@' {
				continue
			}
			currPos := Coords{x, y}
			rolls := 0
			for _, dir := range directions {
				nextPos := move(currPos, dir)
				if withinBounds(lines, nextPos) && lines[nextPos.x][nextPos.y] == '@' {
					rolls++
				}
				if rolls >= 4 {
					break
				}
			}
			if rolls < 4 {
				result++
			}
		}
	}

	return result, err
}

func part2(lines [][]rune) (result int, err error) {
	cols := len(lines[0])
	rows := len(lines)
	done := false
	for !done {
		done = true

		for x := range cols {
			for y := range rows {
				if lines[x][y] != '@' {
					continue
				}
				currPos := Coords{x, y}
				rolls := 0
				for _, dir := range directions {
					nextPos := move(currPos, dir)
					if withinBounds(lines, nextPos) && lines[nextPos.x][nextPos.y] == '@' {
						rolls++
					}
					if rolls >= 4 {
						break
					}
				}
				if rolls < 4 {
					lines[x][y] = '.'
					result++
					done = false
				}
			}
		}

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
