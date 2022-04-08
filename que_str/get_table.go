package que_str

import (
	"fmt"
)

func BuildTable(name string, dataResource []map[string]interface{}, data ComputeData) string {
	piaos := BuildPiaos(dataResource)
	piaos = ComputePiaos(piaos, data)
	ComputePiaos2(piaos)
	tableStr := "<body><p>" + name + "</p><table class=\"pure-table\">"

	if len(dataResource) > 0 {
		disPlayKeys := GetDisplayKeys()

		//表头
		tableStr += " <thead><tr>"
		tableStr += "<td>" + "chang得分" + "</td>"
		tableStr += "<td>" + "fan得分" + "</td>"
		for _, disKey := range disPlayKeys {
			tableStr += "<td>"
			tableStr += disKey
			tableStr += "</td>"
		}
		tableStr += "</tr></thead>"

		//第三行以后
		for i := len(piaos); i > 0; i-- {
			piao := piaos[i-1]
			if i%2 == 0 {
				tableStr += "<tr class=\"pure-table-odd\">"
			} else {
				tableStr += "<tr>"
			}

			//计算得分
			tableStr += "<td>" + fmt.Sprintf("%v+%v+%v+%v+%v+%v=%v", piao.HuanshouScore, piao.LiangbiScore, piao.JingzhanzuoScore, piao.WeipipeiScore, piao.ZhulizengcangScore, piao.JingjiajineScore, piao.GetTotalScore()) + "</td>"
			tableStr += "<td>" + fmt.Sprintf("%v+%v+%v=%v", piao.HuanshouScore2, piao.LiangbiScore2, piao.JingzhanzuoScore2, piao.GetTotalScore2()) + "</td>"
			tableStr += "<td>" + piao.Code + "</td>"
			tableStr += "<td>" + piao.Gupiaojiancheng + "</td>"
			tableStr += "<td>" + fmt.Sprintf("%.2f", piao.Zuixinzhangdiefu) + "</td>"
			tableStr += "<td>" + fmt.Sprintf("%.2f", piao.Huanshou) + "</td>"
			tableStr += "<td>" + fmt.Sprintf("%.2f", piao.Liangbi) + "</td>"
			tableStr += "<td>" + fmt.Sprintf("%.2f", piao.Jingzhanzuo) + "</td>"
			tableStr += "<td>" + fmt.Sprintf("%.2f 万", piao.Weipipei/10000) + "</td>"
			tableStr += "<td>" + fmt.Sprintf("%.2f 万", piao.Jingjiajine/10000) + "</td>"
			tableStr += "<td>" + fmt.Sprintf("%.2f", piao.Zhulizengcang) + "</td>"
			tableStr += "<td>" + fmt.Sprintf("%.2f", piao.Zhangfu) + "</td>"
			tableStr += "<td>" + fmt.Sprintf("%.2f", piao.Weibi) + "</td>"
			tableStr += "<td>" + fmt.Sprintf("%.2f", piao.Jiban) + "</td>"
			tableStr += "</tr>"
		}

	} else {
		tableStr += "<tr><td>" + "无数据" + "</td></tr>"
	}

	tableStr += "</table></body>"

	return tableStr
}

func BuildPiaos(dataResource []map[string]interface{}) []*Piao {
	fmt.Println("%+v", dataResource)
	piaos := make([]*Piao, 0)
	for _, item := range dataResource {
		piao := &Piao{}
		piao.Input(item)
		piaos = append(piaos, piao)
	}

	return piaos
}
