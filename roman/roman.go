package roman

import (
	"bytes"
	"math"
	"strings"
)

func fixRepetition(ls []string, s string) string {
	for i, l := range ls {
		if i != 0 && i%2 == 0 {

			old := strings.Join([]string{ls[i-1], l, l, l, l}, "")
			new := strings.Join([]string{l, ls[i-2]}, "")
			s = strings.Replace(s, old, new, -1)

			old1 := strings.Join([]string{l, l, l, l}, "")
			new1 := strings.Join([]string{l, ls[i-1]}, "")
			s = strings.Replace(s, old1, new1, -1)
		}
	}

	return s
}

func convertedNumberWithRepetition(ls []string, n int) string {
	var buffer bytes.Buffer

	for i, l := range ls {
		var decimalValue float64

		if i%2 == 0 {
			decimalValue = 1
		} else {
			decimalValue = 5
		}

		// This will make for example M becomes 1000; D becomes 500; C becomes 100 etc.
		value := int((decimalValue) * math.Pow10((len(ls)-(i+1))/2))

		for n >= value {
			buffer.WriteString(l)
			n -= value
		}
	}

	return buffer.String()
}

func ConvertNumber(n int) string {
	ls := []string{"M", "D", "C", "L", "X", "V", "I"}
	s := convertedNumberWithRepetition(ls, n)
	s = fixRepetition(ls, s)

	return s
}
