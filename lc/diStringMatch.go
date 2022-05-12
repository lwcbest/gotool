package lc

func diStringMatch(s string) []int {
	min := 0
	max := len(s)
	result := make([]int, 0)
	for _, tar := range s {
		if string(tar) == "I" {
			result = append(result, min)
			min++
		} else {
			result = append(result, max)
			max--
		}
	}

	return result
}
