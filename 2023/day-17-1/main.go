package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

// Make cache local
var cache = make(Cache)


// TODO: Need to handle going straight for 3+ tiles

func Solve(inputPath string) int {
	input := getInput(inputPath)

	heatLoss, res := move(input, Path{}, NewPostition(0, 0))
	if !res {
		log.Fatal("Could not find a path")
	}

	return heatLoss
}

func move(input Input, path Path, pos Position) (int, bool) {
	var posPrev *Position = nil
	if len(path) > 0 {
		posPrev = &path[len(path)-1]
	}
	path = append(path, pos)
	heatLoss := math.MaxInt

	fmt.Printf("Going to %+v; from %+v\n", pos, posPrev)

	// Out of bound
	if !input.Has(pos) {
		fmt.Println("Out of bounds")
		return heatLoss, false
	}

	// Finish
	if pos.I == len(input)-1 && pos.J == len(input[pos.I])-1 {
		fmt.Println("Finish")
		return input.Get(pos), true
	}

	for i := range input {
		for j := range input[i] {
			posCurr := NewPostition(i, j)
			if path.Has(posCurr) || posCurr == pos {
				fmt.Printf("#")
			} else {
				fmt.Printf("%s", string(input[i][j]))
			}
		}
		fmt.Println()
	}
	fmt.Println()
	// time.Sleep(100 * time.Millisecond)

	for _, posCurr := range []Position{pos.Up(), pos.Right(), pos.Down(), pos.Left()} {
		if posPrev != nil && posCurr == *posPrev {
			continue
		}

		heatLossCurr, ok := cache.Get(pos, posCurr)
        if !ok {
            cache.Add(pos, posCurr, math.MaxInt)
            heatLossCurr, ok = move(input, path, posCurr)
            cache.Add(pos, posCurr, heatLossCurr)
        }
		if ok { // dont need this check as if the result is invalid heat loss will be max int
			if heatLossCurr < heatLoss {
				heatLoss = heatLossCurr
			}
		}
	}

	return heatLoss, true
}

func countHeatLoss(input Input, path Path) int {
	heatLoss := 0
	for _, pos := range path {
		heatLoss += input.Get(pos)
	}

	return heatLoss
}

func getInput(path string) Input {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	input := Input{}

	for scanner.Scan() {
		text := scanner.Text()

		input = append(input, []rune(text))
	}

	return input
}

type Input [][]rune

func (input Input) Get(pos Position) int {
	val := input[pos.I][pos.J]

	return int(val)
}

func (input Input) Has(pos Position) bool {
	if pos.I < 0 || pos.J < 0 {
		return false
	}
	if pos.I >= len(input) || pos.J >= len(input[pos.I]) {
		return false
	}

	return true
}

type Position struct {
	I int
	J int
}

func NewPostition(i, j int) Position {
	return Position{I: i, J: j}
}

func (pos Position) Up() Position {
	return NewPostition(pos.I-1, pos.J)
}

func (pos Position) Right() Position {
	return NewPostition(pos.I, pos.J+1)
}

func (pos Position) Down() Position {
	return NewPostition(pos.I+1, pos.J)
}

func (pos Position) Left() Position {
	return NewPostition(pos.I, pos.J-1)
}

type Path []Position

func (path Path) Has(pos Position) bool {
	for _, it := range path {
		if it == pos {
			return true
		}
	}

	return false
}
