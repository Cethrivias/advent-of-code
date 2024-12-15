package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Solve(path string) int {
	input := getInput(path)

	seedRanges := []SeedRange{}
	for len(input.seeds) > 0 {
		seedRange := input.seeds[:2]
		input.seeds = input.seeds[2:]

		seedRanges = append(seedRanges, SeedRange{start: seedRange[0], offset: seedRange[1]})
	}

	locationRanges := input.maps[len(input.maps)-1].ranges
	sort.Slice(locationRanges, func(i, j int) bool {
		return locationRanges[i].target < locationRanges[j].target
	})

	for spot := 0; spot < locationRanges[0].target; spot++ {
		if checkIfSpotHasSeed(input.maps, spot, seedRanges) {
			return spot
		}
	}

	for _, location := range locationRanges {
		if checkIfSpotHasSeed(input.maps, location.target, seedRanges) {
			return location.target
		}
	}

	return -1
}

func checkIfSpotHasSeed(maps []GardenMap, spot int, seedRanges []SeedRange) bool {
	for i := len(maps) - 1; i >= 0; i-- {
		spot = mapLocation(maps[i].ranges, spot)
	}

	for _, seed := range seedRanges {
		if spot >= seed.start && spot < seed.start+seed.offset {
			return true
		}
	}

	return false
}

func mapLocation(granges []GardenRange, spot int) int {
	for _, grange := range granges {
		if spot < grange.target || spot >= grange.target+grange.offset {
			continue
		}

		offset := spot - grange.target
		spot = grange.source + offset

		break
	}

	return spot
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
