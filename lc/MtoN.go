package lc

import "fmt"

// 把M个相同的水果放在N个同样的盘子里，允许有的盘子空着不放，问不同的放法数K是多少？请注意，5，1，1和1，5，1 是同一种放法。
func MToN() {
	inputs := [5][2]int{{1, 2}, {7, 3}, {3, 3}, {5, 3}, {20, 5}}
	outputs := [5]int{0, 0, 0, 0, 0}
	for id, input := range inputs {
		total := dfsWork(input[0], input[1])
		outputs[id] = total
	}
	fmt.Println(outputs)

	outputs2 := [5]int{0, 0, 0, 0, 0}
	for id, input := range inputs {
		total := dpWork(input[0], input[1])
		outputs2[id] = total
	}
	fmt.Println(outputs2)
}

func dfsWork(lA, lP int) int {
	if lA == 0 || lP == 1 {
		//苹果分没了就分完了，或者盘子就剩一个就也算分完了
		return 1
	}

	//苹果小于盘子数量，结果和苹果等于盘子数量是一样的
	if lA < lP {
		return dfsWork(lA, lA)
	}

	//两种情况，把lA个苹果放入lP个盘子的方法，就等于两种情况相加
	//a. 有空盘子的话，相当于算出来把lA个苹果放入lP-1个盘子里面的数量就行了
	//b. 没有空盘子的话，相当于每个盘子先放了一个，然后剩余苹果为lA-lP，计算剩余苹果往lP个盘子里面放的数量就行了
	a := dfsWork(lA, lP-1)
	b := dfsWork(lA-lP, lP)
	return a + b
}

func dpWork(lA, lP int) int {
	data := make([][]int, lA+1)
	//init
	for i := 0; i <= lA; i++ {
		data[i] = make([]int, lP+1)
	}

	for i := 0; i <= lA; i++ {
		data[i][0] = 0
		data[i][1] = 1
	}

	for j := 0; j <= lP; j++ {
		data[0][j] = 1
	}

	for i := 1; i <= lA; i++ {
		for j := 1; j <= lP; j++ {
			if data[i][j] > 0 {
				continue
			}
			if i < j {
				data[i][j] = data[i][i]
			} else {
				data[i][j] = data[i][j-1] + data[i-j][j]
			}
		}
	}

	return data[lA][lP]
}
