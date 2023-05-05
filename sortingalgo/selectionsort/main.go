package main

import (
	"fmt"
)

func main() {

	numbers := []int32{45, 24, 13, 8, 68, 98, 54, 68}

	fmt.Printf("Number list before selection sort: %+v \n", numbers)
	selectionSort(numbers)
	fmt.Printf("Number list after selection sort: %+v \n", numbers)

}

// Time Complexity of selection sort is O(n^2)

// 1. In array/slice select the smallest number and swap it with start of unsorted part of the slice
// 2. At start as all the elements are unsorted, we select the smallest number from slice and swap it with the first index i.e 0th index of the slice
// 3. As now 0th index is sorted our unsorted part of the slice starts with 1st index
// 4. We perform this selection of smallest number from unsorted part of the slice and swap with the first index of the unsorted slice and the move the
// start of the unsorted slice to the next element as that element gets sorted
func selectionSort(numbers []int32) {

	// for loop to get the minimum value from the unsorted part and put it at beginning of unsorted part. Not performing of the last number as single
	// element will always be sorted
	for firstUnsortedIndex := 0; firstUnsortedIndex < len(numbers)-1; firstUnsortedIndex++ {

		// This will give the smallestNumber Index from the unsorted part of the slice
		smallestNumberIndex := firstUnsortedIndex

		// Get the smallestNumber Index from the unsorted part of the slice
		for startOfUnsortedNumber := firstUnsortedIndex; startOfUnsortedNumber < len(numbers); startOfUnsortedNumber++ {

			if numbers[startOfUnsortedNumber] < numbers[smallestNumberIndex] {
				smallestNumberIndex = startOfUnsortedNumber
			}

		}

		// Swap the first Index and the smallest Number Index of the unsorted part of the slice. Now afetr this swap the first Element of unsorted part is
		// now sorted and our unsorted part of array gets shrinked
		swap(&numbers, firstUnsortedIndex, smallestNumberIndex)

	}

}

func swap(numbers *[]int32, numberOne, numberTwo int) {
	// Here passing the address of the slice and not the slice. Can pass slice as well here. But trying to save more memory by passing address of the slice
	// So to access any index of the slice first we need to de-reference the slice and then use index over it
	(*numbers)[numberOne], (*numbers)[numberTwo] = (*numbers)[numberTwo], (*numbers)[numberOne]
}
