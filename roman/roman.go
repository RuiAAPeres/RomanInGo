package roman

import (
	"bytes"
	"fmt"
	"math"
	"strings"
)

func convertedNumberWithPossibleRepetition(array []string, inputNumber int) string {
	var buffer bytes.Buffer

	for index, romanLetter := range array {
		var decimalValue float64

		// There is no ternary operator in GO
		if index%2 == 0 {
			decimalValue = 1
		} else {
			decimalValue = 5
		}

		value := int((decimalValue) * math.Pow10((len(array)-(index+1))/2)) // This will make for example M becomes 1000; D becomes 500; C becomes 100 etc.
		fmt.Println(value)

		for inputNumber >= value {
			buffer.WriteString(romanLetter)
			inputNumber -= value
		}
	}

	return buffer.String()
}

func ConvertNumber(inputNumber int) string {

	array := []string{"M", "D", "C", "L", "X", "V", "I"}

	x := convertedNumberWithPossibleRepetition(array, inputNumber)

	for index, romanLetter := range array {
		if index != 0 && index%2 == 0 {

		}
	}

	fmt.Println(x)

	return x
}
