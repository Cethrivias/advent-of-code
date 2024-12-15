package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Result: 71274
func main() {
	total := 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		power, res := parseGame(text)
		if res {
			total += power
		}
	}

	fmt.Printf("Total: %d\n", total)
}

func parseGame(s string) (int, bool) {
	maxRed := 1
	maxGreen := 1
	maxBlue := 1
	idx := strings.Index(s, ":")
	s = s[idx+1:]

	for _, hand := range strings.Split(s, ";") {
		hand = strings.TrimSpace(hand)
		red, green, blue, res := parseHand(hand)
		if !res {
			return 0, false
		}
		if red > maxRed {
			maxRed = red
		}
		if green > maxGreen {
			maxGreen = green
		}
		if blue > maxBlue {
			maxBlue = blue
		}
	}

	return maxRed * maxGreen * maxBlue, true
}

func parseHand(game string) (int, int, int, bool) {
	red := 1
	green := 1
	blue := 1

	for _, ballz := range strings.Split(game, ",") {
		ballz = strings.TrimSpace(ballz)        // trimming ballz kekw
		splitBallz := strings.Split(ballz, " ") // kekw^2
		color := splitBallz[1]

		count, err := strconv.Atoi(splitBallz[0])
		if err != nil {
			log.Fatal(err.Error())
		}

		switch color {
		case "red":
			red = count
		case "green":
			green = count
		case "blue":
			blue = count
		default:
			log.Fatalf("Color %s does not exist", color)
		}
	}

	return red, green, blue, true
}
