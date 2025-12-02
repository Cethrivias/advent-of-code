package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	ranges := strings.Split(strings.TrimSpace(string(bytes)), ",")

	part1(ranges)
	part2(ranges)
}

func part1(ranges []string) {
	result := 0

	for _, line := range ranges {
		edges := strings.Split(line, "-")
		lo, err := strconv.Atoi(edges[0])
		if err != nil {
			log.Fatal(err)
		}

		hi, err := strconv.Atoi(edges[1])
		if err != nil {
			log.Fatal(err)
		}

		for ; lo <= hi; lo++ {
			num := strconv.Itoa(lo)
			if len(num)%2 != 0 {
				continue
			}

			if num[:len(num)/2] == num[len(num)/2:] {
				result += lo
			}
		}
	}

	fmt.Printf("Part 1: %d\n", result)
}

func part2(ranges []string) {
	result := 0

	for _, line := range ranges {
		edges := strings.Split(line, "-")
		lo, err := strconv.Atoi(edges[0])
		if err != nil {
			log.Fatal(err)
		}

		hi, err := strconv.Atoi(edges[1])
		if err != nil {
			log.Fatal(err)
		}

		for ; lo <= hi; lo++ {
			num := strconv.Itoa(lo)

			if !checkNumber(num) {
				fmt.Printf("id: %s\n", num)
				result += lo
			}
		}
	}

	fmt.Printf("Part 2: %d\n", result)
}

func checkNumber(num string) (valid bool) {
	valid = true

	for i := 1; len(num)/i >= 2; i++ {
		if len(num)%i != 0 {
			continue
		}

		for j := i; j < len(num); j += i {
			if num[0:i] != num[j:j+i] {
				valid = true
				break
			} else {
				valid = false
			}
		}

		if !valid {
			break
		}
	}

	return valid
}
