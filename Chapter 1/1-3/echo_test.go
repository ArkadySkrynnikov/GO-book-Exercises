package echo

import (
	"os"
	"testing"
)

// to run go test -benchmem -run=^$ -bench '^BenchmarkEcho.$'

func BenchmarkEcho1(b *testing.B) {
	os.Args = []string{"cmd1", "arg1", "anotherArg1", "andAnotherArg1"}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res := echo1()
		b.Log(res)
	}
}

func BenchmarkEcho2(b *testing.B) {
	os.Args = []string{"cmd2", "arg2", "anotherArg2", "andAnotherArg2"}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res := echo2()
		b.Log(res)
	}
}

func BenchmarkEcho3(b *testing.B) {
	os.Args = []string{"cmd3", "arg3", "anotherArg3", "andAnotherArg3"}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res := echo3()
		b.Log(res)
	}
}
