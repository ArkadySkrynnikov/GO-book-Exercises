package main

import (
	"main/4/popcount"
	"testing"
)

func BenchmarkLoopPopcount(b *testing.B) {
	popcount.PopCountLoop(25612)
}

func BenchmarkTablePopcount(b *testing.B) {
	popcount.PopCountOrig(25612)
}
