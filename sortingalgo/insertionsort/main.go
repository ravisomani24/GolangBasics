package main

import "fmt"

func main() {

	numbers := []int{45, 24, 13, 8, 68, 98, 54, 68}

	fmt.Printf("Number list before insertion sort: %+v \n", numbers)
	insertionSort(numbers)
	fmt.Printf("Number list after insertion sort: %+v \n", numbers)

}

// 1. Suppose there is sorted part and unsorted part of slice.
// 2. Pick the first element from unsorted part of the slice and place it at correct position in sorted part of the slice
// 3. To place it at correct position we should start comparison from end of sorted to find the correct position. This is becuse the end of sorted part will
// be  surely at the position of first position of unsorted part
func insertionSort(numbers []int) {

	// Starting this from 1st Index and not 0th Index as 0th Index can be considered as one element is sorted part of the slice.
	// This will help in reducing one iteration for first element. This is not necessary. We can start from 0th Index as well
	for firstIndexOfUnsortedPart := 1; firstIndexOfUnsortedPart < len(numbers); firstIndexOfUnsortedPart++ {

		for correctIndexPostion := firstIndexOfUnsortedPart; correctIndexPostion > 0; correctIndexPostion-- {
			if numbers[correctIndexPostion] < numbers[correctIndexPostion-1] {
				// swap these values
				numbers[correctIndexPostion], numbers[correctIndexPostion-1] = numbers[correctIndexPostion-1], numbers[correctIndexPostion]
			} else {
				break
			}
		}

	}

}
