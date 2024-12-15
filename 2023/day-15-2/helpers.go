package main

import (
	"log"
	"strconv"
)

func partInt(val string) int {
	result, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
