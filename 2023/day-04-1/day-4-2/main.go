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
	cardCounts := map[int]int{}
	currCard := 1

	file := getFile(path)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Print(text)

		cardCounts[currCard]++
		text = text[strings.Index(text, ":")+2:]

		winNums := parseNumbers(text[:strings.Index(text, "|")-1])
		nums := parseNumbers(text[strings.Index(text, "|")+2:])

		offset := 0
		for _, num := range nums {
			for _, winNum := range winNums {
				if num != winNum {
					continue
				}
				offset++
			}
		}

		for j := currCard + 1; j <= currCard+offset; j++ {
			cardCounts[j] += 1 * cardCounts[currCard]
		}

		fmt.Printf(" | x%d", cardCounts[currCard])

		currCard++

		fmt.Println()
	}

	for _, count := range cardCounts {
		total += count
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
