package medium

func permute(nums []int) [][]int {

	permutations = make([][]int, 0)
	Permute(nums, 0, len(nums))

	return permutations
}

var permutations [][]int

func Permute(nums []int, start, end int) {
	if start == end {
		tempNums := make([]int, len(nums))
		copy(tempNums, nums)
		permutations = append(permutations, tempNums)
		return
	}

	for i := start; i < end; i++ {
		Swap(&nums, start, i)
		Permute(nums, start+1, end)
		Swap(&nums, start, i)
	}
}

func Swap(nums *[]int, i, j int) {
	tempI, tempJ := (*nums)[i], (*nums)[j]
	(*nums)[j], (*nums)[i] = tempI, tempJ
}
