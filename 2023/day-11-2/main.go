package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

func Solve(path string, multiplier int) (total int) {
	input := getInput(path)
	for _, row := range input.Map {
		println(string(row))
	}
	fmt.Printf("Empty columns: %+v\n", input.EmptyColumns)
	fmt.Printf("Empty rows: %+v\n", input.EmptyRows)

	galaxies := findGalaxies(input.Map)
	galaxyPairs := pairGalaxies(galaxies)

	for _, pair := range galaxyPairs {
		total += findPath(input, pair, multiplier)
	}

	return total
}

func findPath(input *Input, pair []Coords, multiplier int) int {
	path := 0
	for _, i := range input.EmptyRows {
		if i > min(pair[0].I, pair[1].I) && i < max(pair[0].I, pair[1].I) {
			path += multiplier - 1
		}
	}
	for _, j := range input.EmptyColumns {
		if j > min(pair[0].J, pair[1].J) && j < max(pair[0].J, pair[1].J) {
			path += multiplier - 1
		}
	}
	path += int(math.Abs(float64(pair[0].I-pair[1].I)) + math.Abs(float64(pair[0].J-pair[1].J)))
	fmt.Printf("Pair: %+v; path: %d\n", pair, path)

	return path
}

func pairGalaxies(c []Coords) map[string][]Coords {
	pairs := map[string][]Coords{}

	for _, cA := range c {
		hashes := []string{cA.Hash()}
		for _, cB := range c {
			if cA == cB {
				continue
			}

			hashes = append(hashes, cB.Hash())

			sort.Slice(hashes, func(i, j int) bool {
				return hashes[i] > hashes[j]
			})

			pairs[strings.Join(hashes, "|")] = []Coords{cA, cB}
		}
	}

	return pairs
}

func findGalaxies(m Map) []Coords {
	coords := []Coords{}

	for i := range m {
		for j := range m[i] {
			if m[i][j] == '#' {
				coords = append(coords, Coords{I: i, J: j})
			}
		}
	}

	return coords
}

func getInput(path string) *Input {
	input := Input{Map: Map{}, EmptyRows: []int{}, EmptyColumns: []int{}}
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	i := 0
	for scanner.Scan() {
		text := scanner.Text()
		input.Map = append(input.Map, []rune(text))
		if !strings.Contains(text, "#") {
			input.EmptyRows = append(input.EmptyRows, i)
		}

		i++
	}

	filledColumns := make([]bool, len(input.Map[0]))

	for i := range input.Map {
		for j := range input.Map[i] {
			if input.Map[i][j] == '#' {
				filledColumns[j] = true
			}
		}
	}

	for i, filled := range filledColumns {
		if !filled {
			input.EmptyColumns = append(input.EmptyColumns, i)
		}
	}

	return &input
}

type Input struct {
	Map          Map
	EmptyRows    []int
	EmptyColumns []int
}

type Map [][]rune

type Coords struct {
	I int
	J int
}

func (c *Coords) Hash() string {
	return fmt.Sprint(c.I) + ":" + fmt.Sprint(c.J)
}
