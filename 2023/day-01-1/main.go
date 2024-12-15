package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	total := 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		numbers := make([]int, 0)

		for _, r := range text {
			if r < 48 || r > 57 {
				continue
			}
			numbers = append(numbers, int(r-'0'))
		}

		num, err := strconv.Atoi(strconv.Itoa(numbers[0]) + strconv.Itoa(numbers[len(numbers)-1]))
		if err != nil {
			log.Fatal(err.Error())
		}

		total += num
	}

	fmt.Printf("Total: %d\n", total)
}
