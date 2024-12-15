package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Input []Hand

type Hand struct {
	value    string
	score    int
	handType HandType
}

type HandType int

var (
	FiveOfAKind  HandType = 0
	FourOfAKind  HandType = 1
	FullHouse    HandType = 2
	ThreeOfAKind HandType = 3
	TwoPairs     HandType = 4
	Pair         HandType = 5
	HighCard     HandType = 6
)

var labels = []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}

func Solve(path string) int {
	total := 0

	input := getInput(path)

	for i := range input {
		input[i].handType = getHandType(input[i].value)
	}

	sort.Slice(input, func(i, j int) bool {
		return compareHands(input[i], input[j])
	})

	for i, hand := range input {
		total += (i + 1) * hand.score
	}

	return total
}

func getInput(path string) Input {
	input := Input{}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		hand := Hand{}
		test := strings.Split(scanner.Text(), " ")

		hand.value = test[0]

		score, err := strconv.Atoi(test[1])
		if err != nil {
			log.Fatal(err)
		}

		hand.score = score

		input = append(input, hand)
	}

	return input
}

func getLabelPower(label rune) int {
	for power, currlabel := range labels {
		if label == currlabel {
			return power
		}
	}

	log.Fatalf("Could not find the power of label %s", string(label))
	return -1
}

func compareHands(handA Hand, handB Hand) bool {
	if handA.handType == handB.handType {
		for i, labelA := range handA.value {
			powerA := getLabelPower(labelA)
			powerB := getLabelPower(rune(handB.value[i]))

			if powerA == powerB {
				continue
			}

			return powerA > powerB
		}
	}
	return handA.handType > handB.handType
}

func getHandType(hand string) HandType {
	handType := HighCard
	labelsCounts := map[rune]int{}

	for _, rune := range hand {
		labelsCounts[rune]++
	}

	for label, labelCount := range labelsCounts {
		currHandType := func(labelsCount map[rune]int, labelCount int) HandType {
			if labelCount == 5 {
				return FiveOfAKind
			}

			if labelCount == 4 {
				return FourOfAKind
			}

			if labelCount == 3 {
				for _, labelCountB := range labelsCounts {
					if labelCountB == 2 {
						return FullHouse
					}
				}

				return ThreeOfAKind
			}

			if labelCount == 2 {
				for labelB, labelCountB := range labelsCounts {
					if labelCountB == 2 && label != labelB {
						return TwoPairs
					}
				}
				return Pair
			}

			return HighCard
		}(labelsCounts, labelCount)

		if currHandType < handType {
			handType = currHandType
		}
	}

	return handType
}
