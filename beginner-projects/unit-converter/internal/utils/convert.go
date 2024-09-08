package utils

func ConvertUnits(value float64, fromUnit string, toUnit string, conversionMap map[string]float64) float64 {
	return value * conversionMap[fromUnit] / conversionMap[toUnit]
}