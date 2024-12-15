package main

import (
	"fmt"
)

var cache = Cache{}

func Memo(
	records string,
	checksum []int,
	segmentLength int,
	f func(records string, checksum []int, segmentLength int) int,
) int {
	if val, ok := cache.Get(records, checksum, segmentLength); ok {
		return val
	}

	val := f(records, checksum, segmentLength)

	cache.Set(records, checksum, segmentLength, val)

	return val
}

func NewCache() Cache {
	return Cache{}
}

type Cache map[string]int

func (c Cache) Clear() {
	clear(c)
}

func (c Cache) Get(records string, checksum []int, segmentLength int) (int, bool) {
	val, ok := c[hash(records, checksum, segmentLength)]
	return val, ok
}

func (c Cache) Set(records string, checksum []int, segmentLength int, val int) {
	c[hash(records, checksum, segmentLength)] = val

}

func hash(records string, checksum []int, segmentLength int) string {
	return fmt.Sprintf("%s|%+v|%d", records, checksum, segmentLength)
}
