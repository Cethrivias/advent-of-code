package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve(path string) int {
	spot := -1

	input := getInput(path)

	seedRanges := []SeedRange{}
	for len(input.seeds) > 0 {
		seedRange := input.seeds[:2]
		input.seeds = input.seeds[2:]

		seedRanges = append(seedRanges, SeedRange{start: seedRange[0], offset: seedRange[1]})
	}

	for _, r := range seedRanges {
		fmt.Printf("Seed range: %+v\n", r)
		for i := r.start; i < r.start+r.offset; i++ {
			seed := i

			// fmt.Printf("Seed %d", seed)
			for _, gmap := range input.maps {
				seed = mapSeed(gmap.ranges, seed)
			}
			// fmt.Printf(" -> %d\n", seed)
			if spot == -1 || spot > seed {
				spot = seed
			}
		}
	}

	return spot
}

func mapSeed(granges []GardenRange, seed int) int {
	for _, grange := range granges {
		if seed < grange.source || seed >= grange.source+grange.offset {
			continue
		}

		offset := seed - grange.source
		seed = grange.target + offset

		break
	}

	return seed
}

func getInput(path string) *Input {
	input := Input{}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	currMap := ""
	for scanner.Scan() {
		text := scanner.Text()
		println(text)

		if text == "" {
			continue
		}

		if strings.HasPrefix(text, "seeds: ") {
			text := text[len("seeds: "):]
			seedStrs := strings.Split(text, " ")

			for _, seedStr := range seedStrs {
				seed, err := strconv.Atoi(seedStr)
				if err != nil {
					log.Fatal(err)
				}
				input.seeds = append(input.seeds, seed)
			}
			continue
		}

		if strings.Contains(text, "map") {
			currMap = strings.Split(text, " ")[0]
			gardenMap := GardenMap{
				name: currMap,
			}
			input.maps = append(input.maps, gardenMap)
			continue
		}

		valStrs := strings.Split(text, " ")
		gardenRange := GardenRange{
			target: toInt(valStrs[0]),
			source: toInt(valStrs[1]),
			offset: toInt(valStrs[2]),
		}

		for i := range input.maps {
			if input.maps[i].name == currMap {
				input.maps[i].ranges = append(input.maps[i].ranges, gardenRange)
			}
		}

	}

	return &input
}

func toInt(val string) int {
	res, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(err)
	}

	return res
}

type Input struct {
	seeds []int
	maps  []GardenMap
}

type GardenMap struct {
	name   string
	ranges []GardenRange
}

type GardenRange struct {
	source int
	target int
	offset int
}

type SeedRange struct {
	start  int
	offset int
}
