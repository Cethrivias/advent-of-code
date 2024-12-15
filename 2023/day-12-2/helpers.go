package main

func sum(slice []int) (sum int) {
	for _, val := range slice {
		sum += val
	}

	return sum
}
