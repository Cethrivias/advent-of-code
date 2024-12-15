package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var limits = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

// Result: 2149
func main() {
	total := 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		gameId := parseGameId(text)
		if handleBallz(text) {
			total += gameId
		}

		fmt.Printf("gameId: %d, game: %s\n", gameId, text)

	}

	fmt.Printf("Total: %d\n", total)
}

func parseGameId(s string) int {
	idx := strings.Index(s, ":")
	s = s[:idx]
	s = strings.Split(s, " ")[1]

	gameId, err := strconv.Atoi(s)

	if err != nil {
		log.Fatal(err.Error())
	}

	return gameId
}

func handleBallz(s string) bool {
	idx := strings.Index(s, ":")
	s = s[idx+1:]

	for _, hand := range strings.Split(s, ";") {
		hand = strings.TrimSpace(hand)
		for _, ballz := range strings.Split(hand, ",") {
			ballz = strings.TrimSpace(ballz)        // trimming ballz kekw
			splitBallz := strings.Split(ballz, " ") // kekw^2

			limit, exists := limits[splitBallz[1]]
			if !exists {
				log.Fatalf("Color %s does not exist", splitBallz[1])

			}
			count, err := strconv.Atoi(splitBallz[0])
			if err != nil {
				log.Fatal(err.Error())
			}

			if count > limit {
				return false
			}
		}
	}

	return true
}
