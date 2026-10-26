package popcount

func PopCountMove(x uint64) int {
	res := 0

	for x > 0 {
		x = x & (x - 1)
		res++
	}

	return res
}
