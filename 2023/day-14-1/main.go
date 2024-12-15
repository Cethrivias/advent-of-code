package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Solve(path string) (total int) {
	platform := getInput(path)

	platform.Dump()
	fmt.Println()

	for i := 0; i < len(platform); i++ {
		for j := range platform[i] {
			if platform[i][j] != 'O' {
				continue
			}

			c := NewCoords(i, j)
			platform, c = moveUp(platform, c)

			total += len(platform) - c.I
		}
	}

	platform.Dump()
	fmt.Println()

	return total
}

func moveUp(platform Platform, c Coords) (Platform, Coords) {
	for {
		cNew := c.Up()
		if cNew.I < 0 || platform[cNew.I][cNew.J] != '.' {
			break
		}
		platform.Swap(c, cNew)

		c = cNew
	}

	return platform, c
}

func getInput(path string) Platform {
	platform := Platform{}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		platform = append(platform, []rune(scanner.Text()))
	}

	return platform
}

type Platform [][]rune

func (p Platform) Swap(cA Coords, cB Coords) {
	p[cA.I][cA.J], p[cB.I][cB.J] = p[cB.I][cB.J], p[cA.I][cA.J]
}

type Coords struct {
	I int
	J int
}

func (c Coords) Up() Coords {
	return NewCoords(c.I-1, c.J)
}

func NewCoords(i int, j int) Coords {
	return Coords{I: i, J: j}
}

func (p Platform) Dump() {
	for i, row := range p {
		fmt.Printf("%s %d\n", string(row), len(p)-i)
	}
}
