package medium

import "strings"

func defangIPaddr(address string) string {

	return solve1(address)
}

func solve1(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}

func solve2(address string) string {
	var builder strings.Builder

	for _, c := range address {
		if c == '.' {
			builder.WriteString("[.]")
			continue
		}

		builder.WriteRune(c)
	}

	return builder.String()
}