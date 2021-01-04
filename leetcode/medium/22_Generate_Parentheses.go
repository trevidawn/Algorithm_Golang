package medium

import "strings"

func generateParenthesis(n int) []string {
	Parentheses = make([]string, 0, 10)
	temp := make([]byte, n*2)

	gen(0,0,temp)

	return Parentheses
}

var Parentheses []string

func gen(left, right int, temp []byte) {

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
		gen(left+1, right, temp)
		return
	}

	if left > right {
		temp[left+right] = '('
		gen(left+1, right, temp)

		temp[left+right] = ')'
		gen(left, right+1, temp)
		return
	}
}

/*
	다른 사람이 푼 코드인데, 훨씬 깔끔해서 놀랐다.
 */

func generateParenthesis2(n int) []string {
	result := []string{}
	util(&result, "", n, n)
	return result
}


func util(result *[]string, s string, left_required_parans, right_required_parans int) {
	if left_required_parans >0 {
		util(result, s+"(", left_required_parans-1, right_required_parans)
	}
	if left_required_parans < right_required_parans  {
		util(result, s+")", left_required_parans, right_required_parans-1)
	}

	if right_required_parans == 0 {
		*result = append(*result, s)
	}
}