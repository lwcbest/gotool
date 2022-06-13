package lc

import "math"

func MinEatingSpeed(piles []int, h int) int {
	//4,11,20,23,30
	min := 1
	max := 0
	for _, pile := range piles {
		if pile > max {
			max = pile
		}
	}

	result := max

	for min < max {
		speed := (max-min)/2 + min
		time := 0
		for _, pile := range piles {
			time += int(math.Ceil(float64(pile) / float64(speed)))

		}
		if time > h {
			min = speed + 1
		} else {
			result = speed
			max = speed
		}
	}

	return result
}
