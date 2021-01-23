package easy

func climbStairs(n int) int {
	if n < 3 {
		return n
	}


	var D = make([]int, n+1)
	D[1] = 1
	D[2] = 2

	for i := 3; i <= n; i++ {
		D[i] = D[i-2] + D[i-1]
	}

	return D[n]
}

// D[x] =  D[x-2] + D[x-1]