package second

import "fmt"

const boilingF = 212.0

func Second() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("Температура кипения = %g°F или %g°C\n", f, c)
}
