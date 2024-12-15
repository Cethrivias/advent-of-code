package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Solve(path string) (total int) {
	platform := getInput(path)

	states := []Platform{platform.Copy()}
	for i := 0; i < 1_000_000_000; i++ {
		platform = cycle(platform)
		idx := reverseIndex(states, platform)
		if idx != -1 {
			remainingCycles := 1_000_000_000 - idx
			loopLength := i - idx + 1
			finalStateIdx := remainingCycles % loopLength
			platform = states[idx+finalStateIdx]

			break
		}
		states = append(states, platform.Copy())
	}

	total = platform.TotalLoad()

	return total
}

func reverseIndex(states []Platform, p Platform) int {
	for i := len(states) - 1; i >= 0; i-- {
		if states[i].Equal(p) {
			return i
		}
	}

	return -1
}

func cycle(p Platform) Platform {
	for i := 0; i < len(p); i++ {
		for j := range p[i] {
			if p[i][j] != 'O' {
				continue
			}

			p = move(p, NewCoords(i, j), func(c Coords) Coords { return c.North() })
		}
	}
	for i := 0; i < len(p); i++ {
		for j := range p[i] {
			if p[i][j] != 'O' {
				continue
			}

			p = move(p, NewCoords(i, j), func(c Coords) Coords { return c.West() })
		}
	}
	for i := len(p) - 1; i >= 0; i-- {
		for j := range p[i] {
			if p[i][j] != 'O' {
				continue
			}

			p = move(p, NewCoords(i, j), func(c Coords) Coords { return c.South() })
		}
	}
	for i := 0; i < len(p); i++ {
		for j := len(p[i]) - 1; j >= 0; j-- {
			if p[i][j] != 'O' {
				continue
			}

			p = move(p, NewCoords(i, j), func(c Coords) Coords { return c.East() })
		}
	}

	return p
}

func move(platform Platform, c Coords, mover func(c Coords) Coords) Platform {
	for {
		cNew := mover(c)
		if cNew.I < 0 || cNew.I >= len(platform) {
			break
		}
		if cNew.J < 0 || cNew.J >= len(platform[cNew.I]) {
			break
		}
		if platform[cNew.I][cNew.J] != '.' {
			break
		}

		platform.Swap(c, cNew)

		c = cNew
	}

	return platform
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

func (p Platform) Copy() Platform {
	pNew := make(Platform, len(p))
	for i := range p {
		pNew[i] = make([]rune, len(p[i]))
		copy(pNew[i], p[i])
	}

	return pNew
}

func (p Platform) Dump() {
	for i, row := range p {
		fmt.Printf("%s %d\n", string(row), len(p)-i)
	}
}

func (p Platform) Equal(other Platform) bool {
	for i := range p {
		for j := range p[i] {
			if p[i][j] != other[i][j] {
				return false
			}
		}
	}

	return true
}

func (p Platform) TotalLoad() (total int) {
	for i := range p {
		for j := range p[i] {
			if p[i][j] == 'O' {
				total += len(p) - i
			}
		}
	}

	return total
}

func (p Platform) Swap(cA Coords, cB Coords) {
	p[cA.I][cA.J], p[cB.I][cB.J] = p[cB.I][cB.J], p[cA.I][cA.J]
}

type Coords struct {
	I int
	J int
}

func (c Coords) North() Coords {
	return NewCoords(c.I-1, c.J)
}

func (c Coords) West() Coords {
	return NewCoords(c.I, c.J-1)
}

func (c Coords) South() Coords {
	return NewCoords(c.I+1, c.J)
}

func (c Coords) East() Coords {
	return NewCoords(c.I, c.J+1)
}

func NewCoords(i int, j int) Coords {
	return Coords{I: i, J: j}
}
