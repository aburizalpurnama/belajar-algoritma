package sort

import "golang.org/x/exp/constraints"

func VerifySortedAsc[T constraints.Ordered](list []T) bool {
	// if len(list) == 0 || len(list) == 1 {
	// 	return true
	// }

	if len(list) <= 1 {
		return true
	}

	return list[0] < list[1] && VerifySortedAsc(list[1:])
}
