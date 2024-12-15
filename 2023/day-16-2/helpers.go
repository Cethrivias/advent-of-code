package main

func filter(coords []Coords, f func(Coords) bool) []Coords {
	coordsNew := []Coords{}
	for i := range coords {
		if f(coords[i]) {
			coordsNew = append(coordsNew, coords[i])

		}
	}

	return coordsNew
}
