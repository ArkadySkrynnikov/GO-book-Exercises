package tempconv

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func CToK(c Celsius) Kelvin {
	return Kelvin((c + 273.15))
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func KToC(k Kelvin) Celsius {
	return Celsius((k - 273.15))
}

func MetersToFoots(m Meter) Foot {
	return Foot(m * 3.28084)
}

func KiloToPounds(kilo Kilogram) Pound {
	return Pound(kilo * 2.20462)
}
