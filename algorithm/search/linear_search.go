package search

/*
Find the index of target with iteratively compare the value of each element with target
until the target found or reach the end of the list.
*/
func IterativeLinearSearch(numbers []int, target int) int {
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == target {
			return i
		}
	}
	return -1
}
