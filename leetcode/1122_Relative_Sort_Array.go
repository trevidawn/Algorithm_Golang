package leetcode

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {

	var (
		count map[int]int
		ret   []int
	)

	count = map[int]int{}
	for _, val := range arr1 {
		count[val]++
	}

	for _, val := range arr2 {
		for i := 0; i < count[val]; i++ {
			ret = append(ret, val)
		}

		count[val] = 0
	}

	var left []int
	for val, cnt := range count {
		for i := 0; i < cnt; i++ {
			left = append(left, val)
		}
	}

	sort.Ints(left)

	return append(ret, left...)
}
