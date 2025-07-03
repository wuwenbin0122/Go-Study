package tempconv

//CToF converts a Celsus temperautre to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
