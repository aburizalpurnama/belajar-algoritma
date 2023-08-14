package search

/*
Iteratively find the index of target with deviding the list with two so it becomes some sublist
until the target found, or searching reach the end of the list.

It takes O(log n) time.
*/
func IterativeBinarySearch(numbers []int, target int) int {
	startIndex := 0
	endIndex := len(numbers) - 1
	for startIndex <= endIndex {
		midIndex := (startIndex + endIndex) / 2
		// this operation decrease by devided with 2 each iteration (log2 n + 1)
		if numbers[midIndex] == target {
			return midIndex
		} else if target > numbers[midIndex] {
			startIndex = midIndex + 1
		} else {
			endIndex = midIndex - 1
		}
	}
	return -1
}

/*
Recursively the index of target with deviding the list with two so it becomes some sublist
until the target found, or searching reach the end of the list.

It takes O(log n) time.
*/
func RecursiveBinarySearch(numbers []int, target int) bool {
	if len(numbers) == 0 {
		return false
	}

	midIndex := len(numbers) / 2
	if numbers[midIndex] == target {
		return true
	} else {
		if target > numbers[midIndex] {
			return RecursiveBinarySearch(numbers[midIndex+1:], target)
		} else {
			return RecursiveBinarySearch(numbers[:midIndex], target)
		}
	}
}
