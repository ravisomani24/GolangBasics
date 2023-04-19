package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "Señor"

	// Ranging over string which will give us a rune of each character
	rangeStringUsingRune(name)
	fmt.Printf("\n\n")

	// Creating a string from a slice of bytes
	bytes := []byte{82, 97, 118, 105} // Ravi
	createStringFromBytes(bytes)
	fmt.Printf("\n\n")

	// Get correct length of string
	getStringLengthFromRune(name)
	fmt.Printf("\n\n")

	// String is Immutable i.e once a string is created it is not possible to change it. Which means name[0] = 'R' will not work in Go
	// To workaround this string immutability, strings are converted to slice of Runes. Then this slice of Runes is mutated with changes and
	// convert back to a new string
	mutateStrings(name)

	fmt.Printf("\n")
}

func rangeStringUsingRune(name string) {
	fmt.Printf("Traversing string using for each loop")
	fmt.Print("Characters are: ")
	for _, charRune := range name {
		fmt.Printf("%c ", charRune)
	}
	fmt.Printf("\n")
}

func createStringFromBytes(bytes []byte) {
	valueString := string(bytes)
	fmt.Printf("String made from slice of bytes is: %s", valueString)
	fmt.Printf("\n")
}

func getStringLengthFromRune(name string) {
	fmt.Printf("String is: %s, for getting length using bytes and using RuneCountInString function\n", name)
	fmt.Printf("Length of the string using length function is: %d\n", len(name))

	// The len() function gives length of string based on number of bytes in the string and not the string length
	// Some Unicode characters have code points that occupy more than 1 byte. Using len() will give incorrect string length
	// So we should use RuneCountInString function of utf8 package
	fmt.Printf("Length of the string using RuneCountInString function is: %d\n", utf8.RuneCountInString(name))
}

func mutateStrings(name string) {

	fmt.Printf("String name before mutating from Rune: %s\n", name)
	name = mutateRunes([]rune(name))
	fmt.Printf("String name after mutating from Rune:  %s\n", name)

}

func mutateRunes(runeName []rune) string {
	runeName[0] = 'R'
	return string(runeName)
}
