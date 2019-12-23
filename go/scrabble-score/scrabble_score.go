package scrabble

import (
	"strings"
)

//aplphabet with values assigned to each letter
var alphabet = map[string]int{
	"a": 1,
	"e": 1,
	"i": 1,
	"o": 1,
	"u": 1,
	"l": 1,
	"n": 1,
	"r": 1,
	"s": 1,
	"t": 1,
	"d": 2,
	"g": 2,
	"b": 3,
	"c": 3,
	"m": 3,
	"p": 3,
	"f": 4,
	"h": 4,
	"v": 4,
	"w": 4,
	"y": 4,
	"k": 5,
	"j": 8,
	"x": 8,
	"q": 10,
	"z": 10,
}

//Score will calculate and return score of the given word
//based on the value assigned to each letter in alphabet
func Score(word string) int {
	result := 0
	word = strings.ToLower(word)

	for letter, value := range alphabet {
		result += strings.Count(word, letter) * value
	}

	return result
}
