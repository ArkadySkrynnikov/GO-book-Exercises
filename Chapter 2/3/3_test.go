package main

import (
	"main/3/popcount"
	"testing"
)

func BenchmarkLoopPopcount(b *testing.B) {
	popcount.PopCountLoop(122)
}

func BenchmarkTablePopcount(b *testing.B) {
	popcount.PopCountOrig(122)
}
