package medium

func subsets(nums []int) [][]int {

	selected = make([]bool, len(nums))
	subsets := make([][]int, 0)

	for i := 0; i <= len(nums); i++ {
		getSubSet(nums, &subsets, 0, 0, i)
	}

	return subsets
}

var selected []bool

func getSubSet(nums []int, subsets *[][]int, idx, n, r int) {

	if idx == len(nums) {
		if n == r {
			temp := make([]int, 0)
			for i, s := range selected {
				if s == true {
					temp = append(temp, nums[i])
				}
			}

			*subsets = append(*subsets, temp)
		}

		return
	}


	selected[idx] = true
	getSubSet(nums, subsets, idx+1, n+1, r)
	selected[idx] = false
	getSubSet(nums, subsets, idx+1, n, r)
}