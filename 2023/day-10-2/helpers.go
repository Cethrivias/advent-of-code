package main

import "fmt"

func printMap(m Map) {
	mCopy := make(Map, len(m))

	for i := range m {
		mCopy[i] = make([]Tile, len(m[i]))
		for j, tile := range m[i] {
			switch m[i][j].Type {
			case RegularTile:
				mCopy[i][j].Value = 'O'
			case EnclosedTile:
				mCopy[i][j].Value = 'I'
			default:
				mCopy[i][j] = tile
			}
		}
	}

	for _, row := range mCopy {
		print("> ")
		for _, tile := range row {
			fmt.Printf("%s", string(tile.Value))
		}
		println()
	}
}

func sliceContains(slice []Coords, c Coords) bool {
	for _, it := range slice {
		if it == c {
			return true
		}
	}
	return false
}
