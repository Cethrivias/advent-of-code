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

func Solve(path string) (total int) {
	m := getInput(path)

	galaxies := findGalaxies(m)
	galaxyPairs := pairGalaxies(galaxies)

	for _, pair := range galaxyPairs {
		total += findPath(pair)
	}

	return total
}

func findPath(pair []Coords) int {
	return int(math.Abs(float64(pair[0].I-pair[1].I)) + math.Abs(float64(pair[0].J-pair[1].J)))
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

func getInput(path string) Map {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	m := Map{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		m = append(m, []rune(text))
		if !strings.Contains(text, "#") {
			emptyRow := make([]rune, len(text))
			for i := range emptyRow {
				emptyRow[i] = '.'
			}
			m = append(m, emptyRow)
		}
	}

	filledColumns := make([]bool, len(m[0]))

	for i := range m {
		for j := range m[i] {
			if m[i][j] == '#' {
				filledColumns[j] = true
			}
		}
	}

	for i := range m {
		for j := len(m[i]) - 1; j >= 0; j-- {
			if filledColumns[j] {
				continue
			}

			tmp := append([]rune{}, m[i][:j]...)
			tmp = append(tmp, '.')
			tmp = append(tmp, m[i][j:]...)

			m[i] = tmp
		}
	}

	return m
}

type Map [][]rune

type Coords struct {
	I int
	J int
}

func (c *Coords) Hash() string {
	return fmt.Sprint(c.I) + ":" + fmt.Sprint(c.J)
}
