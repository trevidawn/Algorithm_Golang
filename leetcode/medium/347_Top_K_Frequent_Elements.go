package medium

import "sort"

func topKFrequent(nums []int, k int) []int {
	var frequency map[int]int
	frequency = map[int]int{}


	for _, num := range nums {
		if _, ok := frequency[num]; !ok {
			frequency[num] = 1
		} else {
			frequency[num] = frequency[num] + 1
		}
	}

	var t []temp

	for i, v := range frequency {
		t = append(t, temp{i, v})
	}

	sort.Slice(t, func(i, j int) bool {
		return t[i].value > t[j].value
	})

	var ret []int
	for i := 0; i < k; i++ {
		ret = append(ret, t[i].index)
	}

	return ret
}

type temp struct {
	index int
	value int
}