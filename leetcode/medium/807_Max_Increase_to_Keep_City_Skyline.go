package medium

func maxIncreaseKeepingSkyline(grid [][]int) int {

	var (
		maxTopBottom = make([]int, len(grid[0]))
		maxLeftRight = make([]int, len(grid))
	)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if maxLeftRight[i] < grid[i][j] {
				maxLeftRight[i] = grid[i][j]
			}
			if maxTopBottom[j] < grid[i][j] {
				maxTopBottom[j] = grid[i][j]
			}
		}
	}

	var ret int
	var skyline int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			skyline = maxLeftRight[i]
			if maxTopBottom[j] < maxLeftRight[i] {
				skyline = maxTopBottom[j]
			}

			ret += skyline - grid[i][j]
		}
	}

	return ret
}