package main

import (
	"fmt"
	"main/2-2/tempconv"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	for _, arg := range args {
		toFloat, err := strconv.ParseFloat(arg, 64)

		if err != nil {
			fmt.Fprintf(os.Stderr, "failed: %v\n", err)
			os.Exit(1)
		}
		fahrenheit := tempconv.Fahrenheit(toFloat)
		celsius := tempconv.Celsius(toFloat)
		meter := tempconv.Meter(toFloat)
		kilo := tempconv.Kilogram(toFloat)

		fmt.Printf("%s = %s, %s = %s\n", fahrenheit, tempconv.FToC(fahrenheit), celsius, tempconv.CToF(celsius))
		fmt.Printf("%s = %s, %s = %s\n", meter, tempconv.MetersToFoots(meter), kilo, tempconv.KiloToPounds(kilo))
	}
}
