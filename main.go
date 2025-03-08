package main

import (
	"fmt"
)

// Map of roman numerals to arabic numerals
var m = map[rune]int{
    'I': 1,
    'V': 5,
    'X': 10,
    'L': 50,
    'C': 100,
    'D': 500,
    'M': 1000,
}

func romanToInt(s string) int {
    // Edge case handling
    if s == "" {
        return 0 // Empty string
    }

    last_digit := m[rune(s[0])] // Initialize to first character's value
    // last_digit could safely be initialised as 1000 as well
    var arabic int
    
	for _, r := range s {
		digit := m[r]
		if last_digit < digit {
            // subtract twice the previous numeral from arabic to cancel the previous
            // addition and subtract it
			arabic -= 2 * last_digit
		}
		last_digit = digit
		arabic += digit
	}

	return arabic
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
    fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
}