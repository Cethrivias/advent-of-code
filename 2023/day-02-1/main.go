package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var strDigits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

// too high 55725

func main() {
	total := 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		digits := make([]int, 0)

		for i := 0; i < len(text); {
			digit, res := parseByteDigit(text[i])
			if res {
				digits = append(digits, digit)
				i++
				continue
			}

			digit, res = parseStrDigit(text[i:])
			if res {
				digits = append(digits, digit)
			}

			i++
		}

		num, err := strconv.Atoi(strconv.Itoa(digits[0]) + strconv.Itoa(digits[len(digits)-1]))
		if err != nil {
			log.Fatal(err.Error())
		}

		total += num
	}

	fmt.Printf("Total: %d\n", total)
}

func parseByteDigit(r byte) (int, bool) {
	if r < '0' || r > '9' {
		return 0, false
	}
	return int(r - '0'), true
}

func parseStrDigit(s string) (int, bool) {
	for k, v := range strDigits {
		if strings.HasPrefix(s, k) {
			return v, true
		}
	}
	return 0, false
}
