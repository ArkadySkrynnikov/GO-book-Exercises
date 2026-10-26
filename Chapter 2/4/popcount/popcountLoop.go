package popcount

func PopCountLoop(x uint64) int {
	res := 0

	for range 64 {
		res += int(x & 1)
		x = x >> 1
	}

	return res
}
