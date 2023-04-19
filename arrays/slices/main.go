package main

import (
	"fmt"
)

func main() {

	// A slice is a convenient, flexible and powerful wrapper on top of an array.
	// Slices do not own any data on their own. They are just references to existing arrays.
	// Slice Initializing
	slice := []int32{2, 4, 9}
	fmt.Println("Slices before modification and after initializing: ", slice)

	// Any modifications done to the slice will be reflected in the underlying array.
	slice[1] = 5
	fmt.Println("Slices before modification", slice)

	// Length and Capacity
	// The length of the fruitArray is 7. fruitSlice is created from index 1 of fruitArray.
	// Hence, the capacity of fruitSlice is the no of elements in fruitArray starting from index 1 i.e. from orange and that value is 6.
	// Hence, the capacity of fruitSlice is 6. The program prints length of slice 2 capacity 6.
	fruitArray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitSlice := fruitArray[1:3]
	fmt.Printf("length of slice %d capacity %d\n\n", len(fruitSlice), cap(fruitSlice)) //length of fruitSlice is 2 and capacity is 6

	// creating a slice using make
	// func make([]T, len, cap) []T can be used to create a slice by passing the type, length and capacity.
	// The capacity parameter is optional and defaults to the length. The make function creates an array and returns a slice reference to it.

	// One question might be bothering you though. If slices are backed by arrays and arrays themselves are of fixed length then how come a slice
	// is of dynamic length. Well what happens under the hood is, when new elements are appended to the slice, a new array is created.
	// The elements of the existing array are copied to this new array and a new slice reference for this new array is returned. The capacity
	// of the new slice is now twice that of the old slice. Pretty cool right :)

	// The zero value of a slice type is nil. A nil slice has length and capacity 0. It is possible to append values to a nil slice using the append function.
	appendToNilAndEmptySlice()

	appendSliceOfArray()

}

func appendToNilAndEmptySlice() {

	var tempSlice []int32 // zero value of slice is nil. Create nil slice
	if tempSlice == nil {
		fmt.Println("Appending to nil slice. TempSlice:", tempSlice)
		tempSlice = append(tempSlice, 10)
		fmt.Println("After appending to nil slice. TempSlice:", tempSlice)
	}

	tempMakeSlice := make([]int32, 0) // Creates empty slice and this is not nil
	if tempMakeSlice == nil {
		fmt.Println("Appending to nil slice. TempMakeSlice:", tempMakeSlice)
		tempMakeSlice = append(tempMakeSlice, 100)
		fmt.Println("After appending to nil slice. TempMakeSlice:", tempMakeSlice)
	} else {
		fmt.Println("Appending to empty slice. TempMakeSlice:", tempMakeSlice)
		tempMakeSlice = append(tempMakeSlice, 101)
		fmt.Println("After appending to empty slice. TempMakeSlice:", tempMakeSlice)
	}

}

func appendSliceOfArray() {

	// tempArray is [1, 2, 3, 4, 5, 6]
	var tempArray = [6]int32{1, 2, 3, 4, 5, 6}

	// tempSlice is [2, 3, 4]. Slice is the pointer to the underlying array
	tempSlice := tempArray[1:4] // Start from index 1 from start and end at index 4 from start But do not include value of index 4
	fmt.Printf("Array is %+v and Slice is %+v\n", tempArray, tempSlice)

	// Now as we are appending to end of the slice and as slice is made from an array which already have elements after the end of slice
	// so in array the values 5 and 6 are replaced with value 9 and 8 and now end of slice is also end of an array
	tempSlice = append(tempSlice, 9, 8)
	fmt.Printf("After apending in slice, Array is %+v and Slice is %+v\n", tempArray, tempSlice)

}
