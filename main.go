package main

import "fmt"

func main() {
	fmt.Println(sort([]int{6, 1, 0, 2, 4, 3, 7, 5}, func(a, b int) bool {
		return a < b
	}))
}

type sortFunc func(a, b int) bool

func merge(Left, Right []int, fn sortFunc) []int {
	A := make([]int, len(Left)+len(Right))
	i, j, k := 0, 0, 0
	// compare left & right side elements before merging
	for i < len(Left) && j < len(Right) {
		if fn(Left[i], Right[j]) {
			A[k] = Left[i]
			i++
		} else {
			A[k] = Right[j]
			j++
		}
		k++
	}
	// check if any elements from the left/right side
	// were missed in the comparison section above

	for i < len(Left) {
		A[k] = Left[i]
		i++
		k++
	}
	for j < len(Right) {
		A[k] = Right[j]
		j++
		k++
	}

	return A
}

func sort(A []int, fn sortFunc) []int {
	if fn == nil {
		fn = func(a, b int) bool {
			return a < b
		}
	}
	if len(A) > 1 {
		m := len(A) / 2
		Left := A[:m]
		Right := A[m:]

		// sort the right side
		// sort the left side
		// merge the sorted sides
		// do this recursively till the slice has 1 element only
		A = merge(sort(Right, fn), sort(Left, fn), fn)
	}
	return A
}
