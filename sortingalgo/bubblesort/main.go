package main

import "fmt"

func main() {

	numbers := []int{45, 24, 13, 8, 68, 98, 54, 68}

	fmt.Printf("Number list before bubble sort: %+v\n", numbers)
	bubbleSort(numbers)
	fmt.Printf("Number list after bubble sort: %+v\n", numbers)

}

// In Bubble sort we compare the two adjacent elements and order them in sorted manner
// 1. For every value in slice we compare with it's next element.
// 2. If the next element is smaller than current element. We swap these two and make them sorted according to the element.
// 3. By doing this in each iteration we kind of get the largest value to the rightmost position of unsorted array
// 4. We can also improve the bubble sort in best case by checking if there are any swaps or not for the given iteration.
// 5. If there are no swaps that means all elements are already sorted and we can stop at that iteration.
func bubbleSort(numbers []int) {

	// This loop tells that for first iteration we have to compare all the values. And after first iteration we place the biggest number to thr rightmost
	// corner. So we do not need the comparison of that number.
	for uptoIteration := len(numbers) - 1; uptoIteration > 0; uptoIteration-- {

		// This will help to stop if there are no swap for the iteration. If there are no swaps that means the list is already sorted
		didSwap := false

		for start := 0; start < uptoIteration; start++ {

			// Swap current value and it's next value if current value is greater than it's next value
			if numbers[start] > numbers[start+1] {
				numbers[start], numbers[start+1] = numbers[start+1], numbers[start]
				didSwap = true
			}

		}

		if !didSwap {
			break
		}

	}

}
