package second

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func UnrealMix() {
	// s := Celsius.String(12)
	// fmt.Println(s)
	val1 := CToF(15)
	val2 := FToC(15)
	// fmt.Println(val1 + Fahrenheit(val2))
	fmt.Println(val2 + Celsius(val1))
	fmt.Println(val2)
	// fmt.Println(float64(val1) + float64(val2))
}
