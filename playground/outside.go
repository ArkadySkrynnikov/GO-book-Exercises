package main

import (
	"fmt"
	"os"
	"strings"
)

var outside = "outside"

func Outside() {
	fmt.Println(strings.Join(os.Args[1:], " "))

	for i, arg := range os.Args {
		fmt.Println(i, " > ", arg)
	}

	fmt.Println(outside + " mutated")
}
