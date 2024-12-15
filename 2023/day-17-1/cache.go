package main

import "fmt"

type Cache map[string]int

func (cache Cache) Get(prev, curr Position) (int, bool) {
	val, ok := cache[cache.hash(prev, curr)]

	return val, ok
}

func (cache Cache) Add(prev, curr Position, heatLoss int) {
	cache[cache.hash(prev, curr)] = heatLoss
}

func (cache Cache) hash(prev, curr Position) string {
	return fmt.Sprintf("%d:%d|%d:%d", prev.I, prev.J, curr.I, curr.J)
}
