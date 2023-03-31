package main

import "fmt"

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
	fmt.Printf("length of slice %d capacity %d", len(fruitSlice), cap(fruitSlice)) //length of fruitSlice is 2 and capacity is 6

	//creating a slice using make
	//func make([]T, len, cap) []T can be used to create a slice by passing the type, length and capacity.
	// The capacity parameter is optional and defaults to the length. The make function creates an array and returns a slice reference to it.

}
