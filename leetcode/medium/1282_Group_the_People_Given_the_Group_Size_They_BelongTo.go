package medium

func groupThePeople(groupSizes []int) [][]int {

	var ret [][]int
	group := map[int][]int{}

	for personId, size := range groupSizes {
		group[size] = append(group[size], personId)
		if len(group[size]) == size {
			ret = append(ret, group[size])
			group[size] = make([]int, 0, size)
		}
	}

	return ret
}