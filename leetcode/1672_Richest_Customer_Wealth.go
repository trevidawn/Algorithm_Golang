package leetcode

func maximumWealth(accounts [][]int) int {

	var (
		wealth int
		temp   int
	)

	for _, banks := range accounts {
		temp = 0
		for _, money := range banks {
			temp += money
		}

		if wealth < temp {
			wealth = temp
		}
	}

	return wealth
}

