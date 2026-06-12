package luhn

import "unicode"

// Valid determines whether a string is valid according to the Luhn formula.
func Valid(id string) bool {
	sum := 0
	count := 0
	double := false

	// Iterate backwards through the string
	for i := len(id) - 1; i >= 0; i-- {
		r := rune(id[i])
		if unicode.IsSpace(r) {
			continue
		}
		if !unicode.IsDigit(r) {
			return false
		}

		digit := int(r - '0')
		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		count++
		double = !double
	}

	return count > 1 && sum%10 == 0
}
