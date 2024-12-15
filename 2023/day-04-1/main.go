package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Solve(path string) int {
	total := 0

	file := getFile(path)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Print(text)

		text = text[strings.Index(text, ":")+2:]

		winNums := parseNumbers(text[:strings.Index(text, "|")-1])
		nums := parseNumbers(text[strings.Index(text, "|")+2:])

		gameTotal := 0
		for _, num := range nums {
			for _, winNum := range winNums {
				if num != winNum {
					continue
				}
				if gameTotal == 0 {
					gameTotal += 1
					continue
				}
				gameTotal *= 2
			}
		}

		fmt.Printf(" | score: %d", gameTotal)

		total += gameTotal

		fmt.Println()
	}

	return total
}

func getFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func parseNumbers(str string) []string {
	nums := []string{}
	for {
		nums = append(nums, str[:2])

		if len(str) > 2 {
			str = str[3:]
		} else {
			break
		}
	}

	return nums
}
