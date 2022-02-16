package utils

import (
	"fmt"
	"strconv"
)

//TODO - Mover isso para outro lugar e validar com Tom ou Tiago
func Truncate(value float64, decimals int) float64 {

	valueText := fmt.Sprintf("%f", value)
	valueRune := []rune(valueText)
	var resultString string

	for index, char := range valueRune {
		if fmt.Sprintf("%U", char) == "U+002E" {
			resultString = string(valueRune[0 : index+decimals+1])
			break
		}
	}

	resultFloat, err := strconv.ParseFloat(resultString, 64)
	if err != nil {
		return value
	}
	return float64(resultFloat)
}
