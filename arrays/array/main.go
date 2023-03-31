package main

import (
	"fmt"
)

func main() {
	// Array initialization. We need to specify size of array which tells the maximum size of array.
	// It is not possible to increase size of array
	// The size of the array is the part of the type
	var arr [5]int
	fmt.Println("Array before update:", arr)
	arr[0] = 10
	arr[1] = 20
	fmt.Println("Array after update:", arr)

	fmt.Println("***************************************************************")

	// Shorthand initialization
	temp := [5]int{2, 1, 3}
	fmt.Println("Temp Array", temp)

	fmt.Println("***************************************************************")

	// You can even ignore the length of the array in the declaration and replace it with ... and let the compiler find the length for you
	a := [...]int{12, 78, 50} // ... makes the compiler determine the length
	fmt.Println("Array without length declaration", a)

	fmt.Println("***************************************************************")

	// Note: Arrays are value types and not reference types. This means that when they are assigned to a new variable,
	// a copy of original array is assigned to the new variable. If changes are made to the new variable, it will not be reflected
	// in the original array
	// Similarly when arrays are passed to functions as parameters, they are passed by value and the original array is unchanged.

	// Loop over Array using range. This gives us index and value at that index
	for index, value := range temp {
		fmt.Printf("Index is %d and value is %d\n", index, value)
	}

	fmt.Println("***************************************************************")

	// Multidimensional arrays
	multiDimensionalArray := [3][2]string{
		{"lion", "tiger"},
		{"best", "goat"},
		{"saitama", "naruto"},
	}

	// We traverse for each row we print each column value i.e. we print row[0]column[0], row[0]column[1], row[1]column[0]
	for _, rowValue := range multiDimensionalArray {
		for _, columnValue := range rowValue {
			fmt.Printf("%s ", columnValue)
		}
		fmt.Printf("\n")
	}

}
