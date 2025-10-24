package main

import (
	"fmt"
	"main/2-1/tempconv"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	arg, err := strconv.ParseFloat(args[1], 64)

	if err != nil {
		fmt.Errorf("ERROR")
	}

	fmt.Println(tempconv.CToK(tempconv.Celsius(arg)))
}
