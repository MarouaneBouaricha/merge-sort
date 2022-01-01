package main

import "fmt"

func main() {
	fmt.Println(merge([]int{1, 3, 6}, []int{2, 4, 5}))
}

func merge(Left, Right []int) []int {
	A := make([]int, len(Left)+len(Right))
	i, j, k := 0, 0, 0
	// compare left & right side elements before merging
	for i < len(Left) && j < len(Right) {
		if Left[i] < Right[j] {
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
