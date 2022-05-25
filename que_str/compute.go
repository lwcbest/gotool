package que_str

import (
	"sort"
)

func ComputePiaos(piaos []*Piao, computeData ComputeData) []*Piao {
	huanshou := computeData.Huanshou
	liangbi := computeData.Liangbi
	jingzhanzuo := computeData.Jingzhanzuo
	weipipei := computeData.Weipipei
	jingjiajine := computeData.Jingjiajine
	zhuli := computeData.Zhuli
	zhangfu := computeData.Zhangfu

	for _, piao := range piaos {
		for _, query := range huanshou {
			if piao.Huanshou >= query[0] && piao.Huanshou < query[1] {
				piao.HuanshouScore = query[2]
				break
			}
		}

		for _, query := range liangbi {
			if piao.Liangbi >= query[0] && piao.Liangbi < query[1] {
				piao.LiangbiScore = query[2]
				break
			}
		}

		for _, query := range jingzhanzuo {
			if piao.Jingzhanzuo >= query[0] && piao.Jingzhanzuo < query[1] {
				piao.JingzhanzuoScore = query[2]
				break
			}
		}

		for _, query := range weipipei {
			value := piao.Weipipei / piao.Jingjiajine
			if value >= query[0] && value < query[1] {
				piao.WeipipeiScore = query[2]
				break
			}
		}

		for _, query := range jingjiajine {
			if piao.Jingjiajine >= query[0]*10000 && piao.Jingjiajine < query[1]*10000 {
				piao.JingjiajineScore = query[2]
				break
			}
		}

		for _, query := range zhuli {
			if piao.Zhulizengcang >= query[0] && piao.Zhulizengcang < query[1] {
				piao.ZhulizengcangScore = query[2]
				break
			}
		}

		for _, query := range zhangfu {
			if piao.Zhangfu >= query[0] && piao.Zhangfu < query[1] {
				piao.ZhangfuScore = query[2]
				break
			}
		}
	}

	//total
	ones := make([]*One, 0)
	for _, piao := range piaos {
		stu := &One{piao.Code, piao.GetTotalScore()}
		ones = append(ones, stu)
	}

	finalPiaos := make([]*Piao, 0)
	sort.Stable(OneList(ones))
	for _, one := range ones {
		for _, piao := range piaos {
			if one.Name == piao.Code {
				finalPiaos = append(finalPiaos, piao)
				break
			}
		}
	}

	return finalPiaos
}

type One struct {
	Name string
	Num  float64
}

type OneList []*One

func (this OneList) Len() int {
	return len(this)
}
func (this OneList) Less(i, j int) bool {
	return this[i].Num > this[j].Num
}
func (this OneList) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
