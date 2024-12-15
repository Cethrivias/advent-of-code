package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Solve(path string) (total int) {
	input := getInput(path)

	for _, line := range input.Map {
		println(line)
	}

	directions := getDirections(input.Map, input.Start)
	fmt.Printf("Directions: %+v\n", directions)
	for _, direction := range directions {
		println("Route start:", direction.i, direction.j)
		path, ok := move(input.Map, direction, []Coords{})
		if ok {
			total = len(path) / 2
			if len(path)%2 != 0 {
				total++
			}
			break
		}
	}

	return total
}

func move(m []string, c Coords, path []Coords) ([]Coords, bool) {
	fmt.Printf("Moving to: [%d:%d]%s\n", c.i, c.j, string(m[c.i][c.j]))
	if m[c.i][c.j] == 'S' {
		println("Finished loop")
		return path, true
	}
	for _, cVisited := range path {
		if c.i == cVisited.i && c.j == cVisited.j {
			println("Been there")
			return path, false
		}
	}

	path = append(path, c)

	directions := getDirections(m, c)
	for _, direction := range directions {
		path, ok := move(m, direction, path)
		if ok {
			return path, ok
		}
	}

	return path, false
}

func getDirections(m Map, c Coords) []Coords {
	res := []Coords{}

	switch m[c.i][c.j] {
	case 'S':
		if c, ok := m.TopConnected(c); ok {
			res = append(res, c)
		}
		if c, ok := m.RightConnected(c); ok {
			res = append(res, c)
		}
		if c, ok := m.BottomConnected(c); ok {
			res = append(res, c)
		}
		if c, ok := m.LeftConnected(c); ok {
			res = append(res, c)
		}
	case '-':
		if c, ok := m.RightConnected(c); ok {
			res = append(res, c)
		}
		if c, ok := m.LeftConnected(c); ok {
			res = append(res, c)
		}
	case '|':
		if c, ok := m.TopConnected(c); ok {
			res = append(res, c)
		}
		if c, ok := m.BottomConnected(c); ok {
			res = append(res, c)
		}
	case 'L':
		if c, ok := m.TopConnected(c); ok {
			res = append(res, c)
		}
		if c, ok := m.RightConnected(c); ok {
			res = append(res, c)
		}
	case 'F':
		if c, ok := m.RightConnected(c); ok {
			res = append(res, c)
		}
		if c, ok := m.BottomConnected(c); ok {
			res = append(res, c)
		}
	case '7':
		if c, ok := m.BottomConnected(c); ok {
			res = append(res, c)
		}
		if c, ok := m.LeftConnected(c); ok {
			res = append(res, c)
		}
	case 'J':
		if c, ok := m.TopConnected(c); ok {
			res = append(res, c)
		}
		if c, ok := m.LeftConnected(c); ok {
			res = append(res, c)
		}
	default:
		fmt.Printf("Unknown symbol: '%s'", string(m[c.i][c.j]))
	}

	return res
}

func getInput(path string) *Input {
	input := Input{}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	i := 0
	for scanner.Scan() {
		text := scanner.Text()

		input.Map = append(input.Map, text)

		j := strings.Index(text, "S")
		if j != -1 {
			input.Start = Coords{i: i, j: j}
		}
		i++
	}

	return &input
}

type Input struct {
	Map   Map
	Start Coords
}

type Map []string

func (m Map) TopConnected(c Coords) (Coords, bool) {
	return m.hasConnection(c.Top(), "|F7S")
}

func (m Map) RightConnected(c Coords) (Coords, bool) {
	return m.hasConnection(c.Right(), "-J7S")
}

func (m Map) BottomConnected(c Coords) (Coords, bool) {
	return m.hasConnection(c.Bottom(), "|JLS")
}

func (m Map) LeftConnected(c Coords) (Coords, bool) {
	return m.hasConnection(c.Left(), "-FLS")
}

func (m Map) hasConnection(c Coords, pipes string) (Coords, bool) {
	return c, m.withinBounds(c) && strings.Contains(pipes, string(m[c.i][c.j]))
}

func (m Map) withinBounds(c Coords) bool {
	return c.i >= 0 && c.i < len(m) && c.j >= 0 && c.j < len(m[c.i])
}

type Coords struct {
	i int
	j int
}

func (c *Coords) Top() Coords {
	return Coords{i: c.i - 1, j: c.j}
}
func (c *Coords) Right() Coords {
	return Coords{i: c.i, j: c.j + 1}
}
func (c *Coords) Bottom() Coords {
	return Coords{i: c.i + 1, j: c.j}
}
func (c *Coords) Left() Coords {
	return Coords{i: c.i, j: c.j - 1}
}
