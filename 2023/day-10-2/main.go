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

	loop := Loop{}
	directions := getDirections(input.Map, input.Start)
	for _, direction := range directions {
		currLoop, ok := move(input.Map, direction, Loop{})
		if ok {
			total = len(currLoop) / 2
			if len(currLoop)%2 != 0 {
				total++
			}
			loop = currLoop
			break
		}
	}

	// Part 2

	input.Map.MarkLoop(input.Start, loop)

	// Enlarge map
	mapEnchanced := enchance(input.Map)
	println("ENCHANCE!")
	printMap(mapEnchanced)

	// go through tiles 1 by 1
	// skip if loop
	// mark as regular if next to border or have other regular nearby
	//    check if has enclosed nearby then mark them as regular too
	// mark everything else as enclosed

	for i := range mapEnchanced {
		for j := range mapEnchanced[i] {
			// println()
			// printMap(mapEnchanced, loop)
			// time.Sleep(250 * time.Millisecond)
			if mapEnchanced[i][j].Type == LoopTile {
				continue
			}
			if isRegular(mapEnchanced, Coords{i, j}) {
				mapEnchanced[i][j].Type = RegularTile
				spreadRegularity(mapEnchanced, Coords{i, j})
				continue
			}
			mapEnchanced[i][j].Type = EnclosedTile
		}
	}

	println("Marked map:")
	printMap(mapEnchanced)

	// Ensmall map
	println("ENSMALL!")
	input.Map = ensmall(mapEnchanced)
	printMap(input.Map)

	// Counting enclosed tiles
	total = 0
	for i := range input.Map {
		for j := range input.Map[i] {
			if input.Map[i][j].Type == EnclosedTile {
				total++
			}
		}
	}

	return total
}

func enchance(m Map) Map {
	mapEnchanced := make([][]Tile, len(m)*2)

	for i := range m {
		mapEnchanced[i*2] = make([]Tile, len(m[i])*2)
		for j := range m[i] {
			mapEnchanced[i*2][j*2] = m[i][j]

			newTile := NewTile('.', NewCoords(i*2, j*2+1), UnknownTile)
			if m[i][j].Type == LoopTile && strings.Contains("SFL-", string(m[i][j].Value)) {
				newTile.Value = '-'
				newTile.Type = LoopTile
			}
			mapEnchanced[i*2][j*2+1] = newTile
		}

		mapEnchanced[i*2+1] = make([]Tile, len(m[i])*2)
		for j := range m[i] {
			mapEnchanced[i*2+1][j*2] = m[i][j]

			newTile := NewTile('.', NewCoords(i*2+1, j), UnknownTile)
			if m[i][j].Type == LoopTile && strings.Contains("SF7|", string(m[i][j].Value)) {
				newTile.Value = '|'
				newTile.Type = LoopTile
			}
			mapEnchanced[i*2+1][j*2] = newTile
		}
	}

	return mapEnchanced
}

func ensmall(m Map) Map {
	mapEnsmalled := make(Map, 0)
	for i := range m {
		if i%2 != 0 {
			continue
		}
		mapEnsmalled = append(mapEnsmalled, make([]Tile, 0))
		for j := range m[i] {
			if j%2 != 0 {
				continue
			}
			mapEnsmalled[i/2] = append(mapEnsmalled[i/2], m[i][j])
		}
	}

	return mapEnsmalled
}

func isLoop(m Map, l Loop, c Coords) bool {
	if m[c.i][c.j].Value == 'S' {
		return true
	}
	for _, cLoop := range l {
		if c == cLoop {
			return true
		}
	}
	return false
}

func isRegular(m Map, c Coords) bool {
	for _, cCurr := range []Coords{c.Top(), c.Right(), c.Bottom(), c.Left()} {
		if !m.withinBounds(cCurr) || m[cCurr.i][cCurr.j].Type == RegularTile {
			return true
		}
	}
	return false
}

func spreadRegularity(m Map, c Coords) {
	for _, cCurr := range []Coords{c.Top(), c.Right(), c.Bottom(), c.Left()} {
		if m.withinBounds(cCurr) && m[cCurr.i][cCurr.j].Type == EnclosedTile {
			m[cCurr.i][cCurr.j].Type = RegularTile
			spreadRegularity(m, cCurr)
		}
	}
}

func move(m Map, c Coords, loop Loop) (Loop, bool) {
	if m[c.i][c.j].Value == 'S' {
		return loop, true
	}
	for _, cVisited := range loop {
		if c.i == cVisited.i && c.j == cVisited.j {
			return loop, false
		}
	}

	loop = append(loop, c)

	directions := getDirections(m, c)
	for _, direction := range directions {
		loop, ok := move(m, direction, loop)
		if ok {
			return loop, ok
		}
	}

	return loop, false
}

func getDirections(m Map, c Coords) []Coords {
	res := []Coords{}

	switch m[c.i][c.j].Value {
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
		fmt.Printf("Unknown symbol: '%s'", string(m[c.i][c.j].Value))
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
		tiles := []Tile{}
		for j, val := range text {
			tiles = append(tiles, NewTile(val, NewCoords(i, j), UnknownTile))
		}

		input.Map = append(input.Map, tiles)

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

type Map [][]Tile

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

func (m Map) MarkLoop(s Coords, l Loop) Map {
	for _, c := range l {
		m[c.i][c.j].Type = LoopTile
	}
	m[s.i][s.j].Type = LoopTile

	return m
}

func (m Map) hasConnection(c Coords, pipes string) (Coords, bool) {
	return c, m.withinBounds(c) && strings.Contains(pipes, string(m[c.i][c.j].Value))
}

func (m Map) withinBounds(c Coords) bool {
	return c.i >= 0 && c.i < len(m) && c.j >= 0 && c.j < len(m[c.i])
}

type Coords struct {
	i int
	j int
}

func NewCoords(i int, j int) Coords {
	return Coords{i: i, j: j}
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

type Loop []Coords

type Tile struct {
	Type   TileType
	Coords Coords
	Value  rune
}

func NewTile(val rune, c Coords, t TileType) Tile {
	return Tile{Coords: c, Type: t, Value: val}
}

type TileType string

var (
	UnknownTile  TileType = "unknown_tile"
	RegularTile  TileType = "regular_tile"
	EnclosedTile TileType = "enclosed_tile"
	LoopTile     TileType = "loop_tile"
)
