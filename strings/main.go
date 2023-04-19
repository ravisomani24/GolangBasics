package main

import "fmt"

func main() {

	// A string is a slice of bytes in Go. Strings are Immutable in Go. Once a string is created it is not possible to change it

	// Variable name is string and string are created by enclosing a set of characters inside double quotes ""
	name := "Ravi Somani"

	// print string and length of string. Lenth of string is the length of slice of bytes in which the string can be represented.
	// %s format specifier is used to print string and len(s) function to get length of the string
	printStringAndLength(name)

	// Slice of Bytes form the string. %x format specifier to print bytes of the string in hexadecimal and
	// %d format specifier to print bytes of the string in decimal
	printBytes(name)

	// print inidvidual characters of the string name. %c format specifier to print one character of the string
	printChars(name)

	fmt.Printf("\n\nWith new name\n\n")

	// Trying all of the same functions for name Señor. By looking to the string we can say it has length 5. That is 5 characters in the string
	// But this is not the case here
	name = "Señor"

	// The lenght of the string Señor will be 6 as there is one special character 'ñ' which we cannot represent is single uint8 byte.
	// For this special character we need 2 bytes to represent this in uint8 so as length of the string is based on lenght of slice of bytes which
	// represents the string and henc the length will be 6
	printStringAndLength(name)

	// We will see 2 bytes for representing character 'ñ'
	printBytes(name)

	// We will see different character for character 'ñ' while printing individual characters
	printChars(name)

	// Because of this limitation of bytes having uint8 as limit we use rune which has a range of int32
	// A rune is a builin type in go and it is an alias of int32. Rune represent the Unicode code point in Go
	printCharsUsingRune(name)

}

func printStringAndLength(name string) {
	fmt.Printf("Name is: %s and length of string is: %d\n", name, len(name))
}

func printBytes(name string) {

	// Print Bytes in Hexadecimal value
	fmt.Print("Bytes in hexadecimal: ")
	for i := 0; i < len(name); i++ {
		fmt.Printf("%x ", name[i])
	}

	fmt.Printf("\n")

	// Print Bytes in Decimal value
	fmt.Print("Bytes in decimal:     ")
	for i := 0; i < len(name); i++ {
		fmt.Printf("%d ", name[i])
	}

	fmt.Printf("\n")
}

func printChars(name string) {
	fmt.Print("Characters:           ")
	for i := 0; i < len(name); i++ {
		fmt.Printf("%c ", name[i])
	}

	fmt.Printf("\n")
}

func printCharsUsingRune(name string) {

	// tempRunes := []rune{'S', 'e'}
	runes := []rune(name)
	fmt.Printf("Characters with Runes having length %d: ", len(runes))
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}

	fmt.Printf("\n")

	fmt.Printf("Characters with Runes in decimal:      ")
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%d ", runes[i])
	}

	fmt.Printf("\n")

	fmt.Printf("Characters with Runes in hexadecimal:  ")
	for i := 0; i < len(runes); i++ {

		fmt.Printf("%x ", runes[i])
	}

	fmt.Printf("\n")
}
