package main

import (
	"main/5/popcount"
	"testing"
)

func BenchmarkMovePopcount(b *testing.B) {
	popcount.PopCountMove(25612)
}

func BenchmarkTablePopcount(b *testing.B) {
	popcount.PopCountOrig(25612)
}

// 8 4 2 1

// 12
// 1 1 0 0

// 11
// 1 0 1 1

// 12 & 11
// 1 1 0 0
// 1 0 1 1
// 1 0 0 0
// 8

// 8 & 7
// 1 0 0 0
// 0 1 1 1
// 0
