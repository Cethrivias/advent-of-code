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

	input := getInput(path)

	i := 0
	currentNode := "AAA"
	for {
		total++

		if input.Instructions[i] == "L" {
			currentNode = input.AdjMap[currentNode][0]
		} else {
			currentNode = input.AdjMap[currentNode][1]
		}

		// time.Sleep(time.Second * 1)
		if currentNode == "ZZZ" {
			break
		}

		i++
		if i >= len(input.Instructions) {
			i = 0
		}
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
	fmt.Printf("%+v\n", input)

	return &input
}

type Input struct {
	Instructions []string
	AdjMap       map[string][2]string
}
