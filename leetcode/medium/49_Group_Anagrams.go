package medium

import (
	"crypto/sha256"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	var anagrams = map[string][]string{}

	for _, str := range strs {
		runeSlice := make([]rune, 0)
		for _, ch := range str {
			runeSlice = append(runeSlice, ch)
		}

		sort.Slice(runeSlice, func(i, j int) bool {
			return runeSlice[i] < runeSlice[j]
		})

		h := sha256.New()
		for _, ch := range runeSlice {
			h.Write([]byte{byte(ch)})
		}
		key := h.Sum(nil)
		anagrams[string(key)] = append(anagrams[string(key)], str)
	}

	var ret [][]string

	for _, v := range anagrams {
		ret = append(ret, v)
	}
	return ret
}