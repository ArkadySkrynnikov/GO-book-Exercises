package dup

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Dup() {
	input := bufio.NewScanner(os.Stdin)
	counts := make(map[string]int)

	for input.Scan() {
		fmt.Println("input scan WORKS", input.Text())
		counts[input.Text()]++
	}

	fmt.Println("||       input scan STOPS        ||")

	fmt.Println("V-----------SCANNED---------------V")
	fmt.Println(counts)
	fmt.Println("-----------------------------------")

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func Dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]

	fmt.Println(files)
	if len(files) == 0 {

		countLines(os.Stdin, counts)

	} else {

		for _, arg := range files {

			f, err := os.Open(arg)

			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func Dup3() {
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)

		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}

	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
	}
}
