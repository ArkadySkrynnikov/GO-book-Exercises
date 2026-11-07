package ints

import "fmt"

func Run() {
	var u uint8 = 255
	// 255 + 21 = 276
	// 256 128 64 32 16 8 4 2 1

	fmt.Println(u, u+1, u*u) // "255 0 1"
	var i int8 = 127
	fmt.Println(i, i+1, i*i) // "127 -128 1"
	// 0    0   0  0  0 0 0 1 0
	// 0    0   0  1  0 0 0 0 1
	// 34
	var x uint8 = 1<<1 | 1<<5
	// 0    0   0  0  0 0 0 1 0
	// 0    0   0  0  0 0 1 0 0
	// 6
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b, %[1]d\n", x)
	fmt.Printf("%08b, %[1]d\n", y)
	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x|y)
	fmt.Printf("%08b\n", x^y)
	fmt.Printf("%08b\n", x&^y)
}

//32
