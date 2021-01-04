package medium

import "strings"

func generateParenthesis(n int) []string {
	Parentheses = make([]string, 0, 10)
	temp := make([]byte, n*2)

	BackTracking(0,0,temp)

	return Parentheses
}

var Parentheses []string

func BackTracking(left, right int, temp []byte) {

	if len(temp) == left + right {
		var builder strings.Builder

		if left == right {
			for _, v := range temp {
				builder.WriteByte(v)
			}

			Parentheses = append(Parentheses, builder.String())
		}
		return
	}

	if left < right {
		return
	}

	if left == right {
		temp[left+right] = '('
		BackTracking(left+1, right, temp)
		return
	}

	if left > right {
		temp[left+right] = '('
		BackTracking(left+1, right, temp)

		temp[left+right] = ')'
		BackTracking(left, right+1, temp)
		return
	}
}