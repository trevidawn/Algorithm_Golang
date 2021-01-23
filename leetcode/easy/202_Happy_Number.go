package easy

/*
	Runtime: 0 ms, faster than 100.00% of Go online submissions for Happy Number.
	Memory Usage: 2.1 MB, less than 77.06% of Go online submissions for Happy Number.
*/

func isHappy(n int) bool {
	var numbers = map[int]struct{}{}

	for {
		n = Process(n)
		if n == 1 {
			break;
		}
		if _, ok := numbers[n]; ok {
			return false
		}
		numbers[n] = struct{}{}
	}

	return true
}

func Process(n int) int {
	var nums = 0

	for n > 0 {
		nn := (n%10)
		nums += nn * nn

		n /= 10
	}

	return nums
}
