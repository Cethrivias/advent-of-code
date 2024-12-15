package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

// Result Part 1: 560670
// Result Part 2: 91622824
func main() {
	numbers := map[string]int{}
	total := 0
	totalGear := 0

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	lines := [][]rune{}

	for scanner.Scan() {
		lines = append(lines, []rune(scanner.Text()))
	}

	for i, line := range lines {
		for j, rune := range line {
			if rune != '.' && !unicode.IsDigit(rune) {
				adjNumbers := findAdjacentNumbers(lines, i, j)
				for key, num := range adjNumbers {
					numbers[key] = num
				}

				// PART 2
				if rune == '*' && len(adjNumbers) == 2 {
					gear := 1
					for _, num := range adjNumbers {
						gear *= num
					}

					totalGear += gear
				}
			}
		}
	}

	for _, num := range numbers {
		total += num
	}

	fmt.Printf("Part 1: %d\n", total)
	fmt.Printf("Part 2: %d\n", totalGear)
}

func findAdjacentNumbers(lines [][]rune, i int, j int) map[string]int {
	numbers := map[string]int{}

	for iAdj := max(i-1, 0); iAdj <= min(i+1, len(lines)-1); iAdj++ {
		for jAdj := max(j-1, 0); jAdj <= min(j+1, len(lines[iAdj])-1); jAdj++ {
			if unicode.IsDigit(lines[iAdj][jAdj]) {
				num := findNumber(lines, iAdj, jAdj)
				numbers[num.Hash()] = num.value
			}
		}
	}

	return numbers
}

func findNumber(lines [][]rune, i int, j int) Number {
	line := lines[i]
	numRunes := []rune{line[j]}
	coords := NewCoords(i, j)

	for x := j - 1; x >= 0 && unicode.IsDigit(line[x]); x-- {
		numRunes = append([]rune{line[x]}, numRunes...)
		coords.j = x
	}
	for x := j + 1; x < len(line) && unicode.IsDigit(line[x]); x++ {
		numRunes = append(numRunes, line[x])
	}

	num, err := strconv.Atoi(string(numRunes))
	if err != nil {
		log.Fatal(err)
	}

	return NewNumber(num, coords)
}

func NewNumber(value int, pos Coords) Number {
	return Number{position: pos, value: value}
}

func NewCoords(i int, j int) Coords {
	return Coords{i: i, j: j}
}

type Number struct {
	position Coords
	value    int
}

func (n *Number) Hash() string {
	return fmt.Sprintf("%d:%d", n.position.i, n.position.j)
}

type Coords struct {
	i int
	j int
}
