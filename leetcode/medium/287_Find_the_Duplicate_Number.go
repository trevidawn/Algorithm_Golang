package medium

func findDuplicate(nums []int) int {
	var numSet = map[int]struct{}{}

	var ret int
	for _, num := range nums {
		if _, ok := numSet[num]; ok {
			ret = num
			break
		}

		numSet[num] = struct{}{}
	}

	return ret
}