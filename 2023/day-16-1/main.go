package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Solve(path string) (total int) {
	grid := getInput(path)

	move(grid, []Coords{}, Visited{}, NewCoords(0, 0))

	for _, row := range grid {
		for _, tile := range row {
			if tile.Energized {
				total++
			}

		}
	}

	return total
}

func move(g Grid, p []Coords, v Visited, c Coords) {
	if !g.Exists(c) {
		return
	}

	cNext := getNextCoords(g, p, v, c)

	g.GetTile(c).Energized = true
	p = append(p, c)

	for _, cCurr := range cNext {
		v.Add(c, cCurr)
		move(g, p, v, cCurr)
	}
}

func getNextCoords(grid Grid, path []Coords, visited Visited, c Coords) []Coords {
	cNext := []Coords{}

	cPrev := NewCoords(0, -1)
	if len(path) != 0 {
		cPrev = path[len(path)-1]
	}

	iPrevOffset := c.I - cPrev.I
	jPrevOffset := c.J - cPrev.J

	switch grid.GetTile(c).Value {
	case '.':
		cNext = append(cNext, NewCoords(c.I+iPrevOffset, c.J+jPrevOffset))
	case '|':
		directions := []Coords{NewCoords(-1, 0), NewCoords(1, 0)}

		for _, dir := range filter(directions, func(c Coords) bool {
			return !(iPrevOffset*-1 == c.I && jPrevOffset*-1 == c.J)
		}) {
			cNext = append(cNext, NewCoords(c.I+dir.I, c.J+dir.J))
		}
	case '-':
		directions := []Coords{NewCoords(0, -1), NewCoords(0, 1)}

		for _, dir := range filter(directions, func(c Coords) bool {
			return !(iPrevOffset*-1 == c.I && jPrevOffset*-1 == c.J)
		}) {
			cNext = append(cNext, NewCoords(c.I+dir.I, c.J+dir.J))
		}
	case '/':
		cNext = append(cNext, NewCoords(c.I+(jPrevOffset)*-1, c.J+(iPrevOffset)*-1))
	case '\\':
		cNext = append(cNext, NewCoords(c.I+(jPrevOffset), c.J+(iPrevOffset)))
	}

	cNext = filter(cNext, func(cCurr Coords) bool {
		return !visited.Has(c, cCurr)
	})

	return cNext
}

func getInput(path string) Grid {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	grid := Grid{}

	for scanner.Scan() {
		tiles := []Tile{}

		row := scanner.Text()

		for i, val := range row {
			c := NewCoords(len(grid), i)
			tiles = append(tiles, NewTile(val, c, false))
		}

		grid = append(grid, tiles)
	}

	return grid
}

type Grid [][]Tile

func (g Grid) Dump(c Coords) {
	for _, row := range g {
		for _, tile := range row {
			if tile.Coords == c {
				fmt.Print("#")
			} else {
				fmt.Printf("%s", string(tile.Value))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g Grid) Exists(c Coords) bool {
	return c.I >= 0 && c.I < len(g) && c.J >= 0 && c.J < len(g[c.I])
}

func (g Grid) GetTile(c Coords) *Tile {
	return &g[c.I][c.J]
}

type Tile struct {
	Value     rune
	Coords    Coords
	Energized bool
}

func NewTile(v rune, c Coords, e bool) Tile {
	return Tile{Value: v, Coords: c, Energized: e}
}

type Coords struct {
	I int
	J int
}

func NewCoords(i int, j int) Coords {
	return Coords{I: i, J: j}
}

type Visited map[string]bool

func (v Visited) Has(prev, curr Coords) bool {
	_, has := v[v.hash(prev, curr)]

	return has
}

func (v Visited) Add(prev, curr Coords) {
	v[v.hash(prev, curr)] = true
}

func (v Visited) hash(prev, curr Coords) string {
	return fmt.Sprintf("%d:%d|%d:%d", prev.I, prev.J, curr.I, curr.J)
}
