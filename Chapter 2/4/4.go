package main

import (
	"fmt"
	"main/3/popcount"
)

func main() {
	// 122
	// 128 64 32  16 8 4 2 1
	//  0   1  1  1  1 0 1 0
	fmt.Println(popcount.PopCountOrig(122))
	fmt.Println(popcount.PopCountLoop(122))
}
