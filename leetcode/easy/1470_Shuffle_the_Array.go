package easy

func shuffle(nums []int, n int) []int {

	output := make([]int, 0, len(nums)*2)
	half   := nums[n:]

	for i := 0; i < n; i++ {
		output = append(output, nums[i], half[i])
	}

	return output
}