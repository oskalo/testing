package benchmarks

// Calculate returns x + 2.
func Calculate(x int) (result int) {
	result = x + 2
	return result
}

/*func Calculate(x int) (result int) {
	result += x
	for i := 0; i < 25; i++ {
		result++
	}
	return result
}*/
