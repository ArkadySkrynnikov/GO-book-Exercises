package popcount

var pcLoop [256]byte

func init() {
	for i := range pcLoop {
		pcLoop[i] = pcLoop[i/2] + byte(i&1)
	}
}

func PopCountLoop(x uint64) int {
	sum := 0
	for i := range 7 {
		sum += int(pcLoop[byte(x>>(i*8))])
	}
	return sum
}
