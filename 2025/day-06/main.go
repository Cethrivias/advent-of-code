package day06

import (
	"strconv"
	"strings"
)

func part1(lines []string) (result int, err error) {
	worksheet := [][]int{}
	operations := []string{}

	for i := 0; i < len(lines)-1; i++ {
		numberStrings := strings.Split(lines[i], " ")
		numbers := []int{}
		for _, numberString := range numberStrings {
			if numberString == "" {
				continue
			}

			number, err := strconv.Atoi(numberString)
			if err != nil {
				return result, err
			}

			numbers = append(numbers, number)
		}

		worksheet = append(worksheet, numbers)
	}

	for _, operation := range strings.Split(lines[len(lines)-1], " ") {
		if operation == "" {
			continue
		}

		operations = append(operations, operation)
	}

	for i, op := range operations {
		switch op {
		case "+":
			problem := 0
			for j := range len(worksheet) {
				problem += worksheet[j][i]
			}
			result += problem
		case "*":
			problem := 1
			for j := range len(worksheet) {
				problem *= worksheet[j][i]
			}
			result += problem
		}
	}

	return result, err
}

func part2(lines []string) (result int, err error) {
	rows := len(lines)
	cols := len(lines[0])

	numbers := []int{}
	operator := ' '

	// Going top-to-bottom right-to-left through the worksheet
	for col := cols - 1; col >= 0; col-- {
		// Scanning the column.
		// Collecting digits characters into a "buffer"
		// and searching for the operator for the current problem
		numberBuf := []rune{}
		for row := range rows {
			val := rune(lines[row][col])
			// Skipping empty spaces
			if val == ' ' {
				continue
			}

			// Operator always going to be last for each problem with how I'm parsing the input
			// It's not important for my solution, but maybe the calculation can be triggered at this point
			// to optimize the solution a bit
			if val == '*' || val == '+' {
				operator = val
				continue
			}

			numberBuf = append(numberBuf, val)
		}

		// At this point the column is done.
		// If there is something in the buffer, parse it into an int and continue to the next column
		if len(numberBuf) != 0 {
			number, err := strconv.Atoi(string(numberBuf))
			if err != nil {
				return result, err
			}
			numbers = append(numbers, number)
			continue
		}

		// If the buffer is empty it means we scanned an empty column
		// So the problem was fully scanned. We can do the calculation and reset the numbers
		switch operator {
		case '+':
			problem := 0
			for _, number := range numbers {
				problem += number
			}
			result += problem
		case '*':
			problem := 1
			for _, number := range numbers {
				problem *= number
			}
			result += problem
		}
		numbers = []int{}
	}

	// After we scan the last column there is no "empty-column" to trigger the calculation as in the main loop
	// So I run the same code code again.
	// I think this duplication can be fixed if we trigger calculation as soon as we find an operator. But I can't be bothered
	switch operator {
	case '+':
		problem := 0
		for _, number := range numbers {
			problem += number
		}
		result += problem
	case '*':
		problem := 1
		for _, number := range numbers {
			problem *= number
		}
		result += problem
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
