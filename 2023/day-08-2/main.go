package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

func Solve(path string) int {
	total := 0

	input := getInput(path)

	currNodes := []string{}
	for node := range input.AdjMap {
		if strings.HasSuffix(node, "A") {
			currNodes = append(currNodes, node)
		}
	}

	sort.Slice(currNodes, func(i, j int) bool {
		return currNodes[i] > currNodes[j]
	})

	paths := make([]int, len(currNodes))
	doneNodes := make([]bool, len(currNodes))

	i := 0
	for {
		done := true

		for j := range currNodes {
			if doneNodes[j] {
				continue
			}

			if input.Instructions[i] == "L" {
				currNodes[j] = input.AdjMap[currNodes[j]][0]
			} else {
				currNodes[j] = input.AdjMap[currNodes[j]][1]
			}

			paths[j]++

			if !strings.HasSuffix(currNodes[j], "Z") {
				done = false
			} else {
				doneNodes[j] = true
			}
		}

		i++
		if i >= len(input.Instructions) {
			i = 0
		}
		if done {
			break
		}
	}

	totals := []int{}
	for _, val := range paths {
		totals = append(totals, val)
	}

	biggestPath := 0
	for _, val := range paths {
		if val > biggestPath {
			biggestPath = val
		}
	}

	total = biggestPath
	for {
		done := true

		for _, val := range paths {
			if total%val != 0 {
				done = false
				break
			}
		}

		if done {
			break
		}

		total += biggestPath
	}

	return total
}

func getInput(path string) *Input {
	input := Input{AdjMap: map[string][2]string{}}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	input.Instructions = strings.Split(scanner.Text(), "")

	scanner.Scan()

	for scanner.Scan() {
		text := scanner.Text()
		kvPair := strings.Split(text, " = ")

		key := kvPair[0]
		nodesStr := kvPair[1][1 : len(kvPair[1])-1]
		nodes := strings.Split(nodesStr, ", ")
		input.AdjMap[key] = [2]string{
			nodes[0],
			nodes[1],
		}

	}

	return &input
}

type Input struct {
	Instructions []string
	AdjMap       map[string][2]string
}
