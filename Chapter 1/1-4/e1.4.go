package main

import (
	"bufio"
	"fmt"
	"os"
)

// Измените программу dup2 так, чтобы она выводила имена всех
// файлов, в которых найдены повторяющиеся строки.
// go run .\e1.4.go testfile.txt testfile2.txt

func main() {
	SimplierDup()
	// Dup()
}

func SimplierDup() {
	counts := make(map[string]int)
	resStr := ""
	files := os.Args[1:]

	if len(files) == 0 {
		dupStdin := hasDupCount(os.Stdin, counts)
		if dupStdin {
			resStr += "Duplicate in - stdin"
		}

	} else {
		for _, filename := range files {

			file, err := os.Open(filename)

			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			dup := hasDupCount(file, counts)

			if dup {
				resStr += "Duplicate in - " + filename
			}

			file.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	fmt.Println(resStr)
}

func hasDupCount(f *os.File, counts map[string]int) bool {
	input := bufio.NewScanner(f)
	hasDup := false

	for input.Scan() {
		if input.Text() != "" {
			counts[input.Text()]++

			if counts[input.Text()] > 1 {
				hasDup = true
			}
		}
	}
	return hasDup
}

func Dup() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines("stdin", os.Stdin, counts)
	} else {
		for _, arg := range files {

			file, err := os.Open(arg)

			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLines(arg, file, counts)

			file.Close()
		}
	}

	for line := range counts {
		for _, dups := range counts[line] {
			if dups > 1 {
				fmt.Println("Duplicate in ", line)
				break
			}
		}
	}
}

func countLines(name string, f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)

	if counts[name] == nil {
		counts[name] = make(map[string]int)
	}

	for input.Scan() {
		counts[name][input.Text()]++
	}
}
