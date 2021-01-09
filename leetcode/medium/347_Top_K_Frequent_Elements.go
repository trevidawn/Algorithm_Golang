package medium

import "sort"

func topKFrequent(nums []int, k int) []int {
	var (
		frequency = make(map[int]int)
		keys      []int
	)

	for _, num := range nums {
		frequency[num]++
	}

	for key := range frequency {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return frequency[keys[i]] > frequency[keys[j]]
	})

	return keys[:k]
}
