package lc

import "fmt"

func TestMin() {
	strs := []string{"cba", "daf", "ghi"}

	result := minDeletionSize(strs)
	fmt.Println(result)
}

func minDeletionSize(strs []string) int {
	xMax := len(strs[0])
	yMax := len(strs)

	badx := make([]int, 0)
	for x := 0; x < xMax; x++ {
		for y := 0; y < yMax-1; y++ {
			a := strs[y][x]
			b := strs[y+1][x]
			if a > b {
				badx = append(badx, x)
				break
			}
		}
	}

	return len(badx)
}
