package easy

func kidsWithCandies(candies []int, extraCandies int) []bool {

	var maxCandy int

	for _, candy := range candies {
		if maxCandy < candy {
			maxCandy = candy
		}
	}

	output := make([]bool, 0, len(candies))
	for _, candy := range candies {
		if candy + extraCandies >= maxCandy {
			output = append(output, true)
			continue
		}

		output = append(output, false)
	}

	return output
}