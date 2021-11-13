package main

func Average(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}

	total := float64(0)
	for _, x := range values {
		total += x
	}

	return total / float64(len(values))
}
