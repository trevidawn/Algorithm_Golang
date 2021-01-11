package medium

func productExceptSelf(nums []int) []int {
	left := make([]int, len(nums))
	right := make([]int, len(nums))

	left[0] = 1
	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	right[len(nums)-1] = 1
	for i := len(nums)-2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	product := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		product[i] = left[i] * right[i]
	}

	return product
}