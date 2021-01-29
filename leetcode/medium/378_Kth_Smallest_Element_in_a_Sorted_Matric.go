package medium

/*
	Runtime: 40 ms, faster than 33.33% of Go online submissions for Kth Smallest Element in a Sorted Matrix.
	Memory Usage: 7 MB, less than 11.90% of Go online submissions for Kth Smallest Element in a Sorted Matrix.
 */
import "sort"

func kthSmallest(matrix [][]int, k int) int {
	var numbers = make([]int, 0)

	for _, row := range matrix {
		numbers = append(numbers, row...)
	}

	sort.Ints(numbers)

	return numbers[k-1]
}