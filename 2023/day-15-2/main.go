package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strings"
)

func Solve(path string) (total int) {
	input := getInput(path)

	boxes := make(Boxes, 256)

	for i := range input {
		lens := NewLens("", -1)

		if strings.HasSuffix(input[i], "-") {
			lens.Label = input[i][:len(input[i])-1]
		} else {
			tmp := strings.Split(input[i], "=")

			lens.Label = tmp[0]
			lens.FocalLength = partInt(tmp[1])
		}

		if lens.FocalLength == -1 {
			boxes.Remove(lens)
		} else {
			boxes.Add(lens)
		}
	}

	for boxIdx, box := range boxes {
		boxPower := 0
		for lensIdx, lens := range box {
			boxPower += (boxIdx + 1) * (lensIdx + 1) * lens.FocalLength
		}
		total += boxPower
	}

	return total
}

func getInput(path string) []string {
	input := []string{}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		input = strings.Split(text, ",")
	}

	return input
}

type Lens struct {
	Label       string
	FocalLength int
}

func (l Lens) Hash() (hash int) {
	for _, rune := range l.Label {
		hash = ((hash + int(rune)) * 17) % 256
	}

	return hash
}

func NewLens(label string, focalLength int) Lens {
	return Lens{Label: label, FocalLength: focalLength}
}

type Boxes [][]Lens

func (b Boxes) Add(lens Lens) {
	idx := lens.Hash()
	lensIdx := slices.IndexFunc(b[idx], func(l Lens) bool {
		return l.Label == lens.Label
	})

	if lensIdx == -1 {
		b[idx] = append(b[idx], lens)
	} else {
		b[idx][lensIdx] = lens
	}
}

func (b Boxes) Remove(lens Lens) {
	idx := lens.Hash()
	lensIdx := slices.IndexFunc(b[idx], func(l Lens) bool {
		return l.Label == lens.Label
	})

	if lensIdx != -1 {
		b[idx] = append(b[idx][:lensIdx], b[idx][lensIdx+1:]...)
	}
}
