package sort

import "golang.org/x/exp/constraints"

/*
Sorting in ascending order of the list using deviding and conquer method.
It split a list into sublists until the sublist only contain one element,
and then merge all of sublist and sort it at once.

It takes O(n log n) time
*/
func MergeSort[T constraints.Ordered](list []T) []T {
	if len(list) <= 1 {
		return list
	}
	leftHalf, rightHalf := split(list)
	left := MergeSort(leftHalf)
	right := MergeSort(rightHalf)
	return merge(left, right)
}

/*
Devide the unsorted list at midpoint into sublist
Returns two sublists, left and right

Takes overall O(log n) time
*/
func split[T constraints.Ordered](list []T) ([]T, []T) {
	midIndex := len(list) / 2
	return list[:midIndex], list[midIndex:]
}

/*
Merge and sort each element of two sublists
Returns merged and sorted of two sublists

Takes Overall O(n)
*/
func merge[T constraints.Ordered](left, right []T) []T {
	var result []T
	i := 0
	j := 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	for i < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}
