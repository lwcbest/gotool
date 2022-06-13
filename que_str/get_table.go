package que_str

import (
	"fmt"
	"github.com/lwcbest/gotool/useexcel"
	"strconv"
	"time"
)

func BuildTable(name string, dataResource []map[string]interface{}, data ComputeData) string {
	piaos := BuildPiaos(dataResource)
	piaos = ComputePiaos(piaos, data)
	tableStr := "<body><p>" + name + "</p><table class=\"pure-table\">"

	if len(dataResource) > 0 {
		disPlayKeys := GetDisplayKeys()

		//表头
		tableStr += " <thead><tr>"
		tableStr += "<td>" + "score" + "</td>"

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
			tableStr += "<td>" + fmt.Sprintf("%v+%v+%v+%v+%v+%v+%v=%v", piao.HuanshouScore, piao.LiangbiScore, piao.JingzhanzuoScore, piao.WeipipeiScore, piao.ZhulizengcangScore, piao.JingjiajineScore, piao.ZhangfuScore, piao.GetTotalScore()) + "</td>"
			myTime, _ := strconv.Atoi(piao.Shijian[1:3])
			if piao.TotalScore >= 60 {
				if myTime < 11 {
					tableStr += "<td><font color=\"red\">" + piao.Code + "</font></td>"
				} else {
					tableStr += "<td><font color=\"green\">" + piao.Code + "</font></td>"
				}
			} else {
				tableStr += "<td>" + piao.Code + "</td>"
			}
			tableStr += "<td>" + piao.Gupiaojiancheng + "</td>"
			tableStr += "<td>" + fmt.Sprintf("%.2f", piao.Zuixinzhangdiefu) + "</td>"
			tableStr += "<td>" + fmt.Sprintf("%.2f (%v)", piao.Huanshou, piao.HuanshouScore) + "</td>"
			tableStr += "<td>" + fmt.Sprintf("%.2f (%v)", piao.Liangbi, piao.LiangbiScore) + "</td>"
			tableStr += "<td>" + fmt.Sprintf("%.2f (%v)", piao.Jingzhanzuo, piao.JingzhanzuoScore) + "</td>"
			tableStr += "<td>" + fmt.Sprintf("%.2f w", piao.Weipipei/10000) + "</td>"
			tableStr += "<td>" + fmt.Sprintf("%.2f w (%v)", piao.Jingjiajine/10000, piao.JingjiajineScore) + "</td>"
			tableStr += "<td>" + fmt.Sprintf("%.2f (%v)", piao.Weipipei/piao.Jingjiajine, piao.WeipipeiScore) + "</td>"
			tableStr += "<td>" + fmt.Sprintf("%.2f (%v)", piao.Zhulizengcang, piao.ZhulizengcangScore) + "</td>"
			tableStr += "<td>" + fmt.Sprintf("%.2f (%v)", piao.Zhangfu, piao.ZhangfuScore) + "</td>"
			tableStr += "<td>" + fmt.Sprintf("%.2f", piao.Weibi) + "</td>"
			tableStr += "<td>" + fmt.Sprintf("%.2f", piao.Jiban) + "</td>"
			tableStr += "<td>" + fmt.Sprintf("%v", piao.Shijian) + "</td>"
			tableStr += "</tr>"
		}

	} else {
		tableStr += "<tr><td>" + "no data" + "</td></tr>"
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

func SaveTable(filename string, dataResource []map[string]interface{}) error {
	piaos := BuildPiaos(dataResource)
	fmt.Println(len(piaos))
	sheet := useexcel.ReadExcelSheet(filename, 0)
	for _, piao := range piaos {
		row := sheet.AddRow()
		cell := row.AddCell()
		cell.SetString(piao.Code)
		cell = row.AddCell()
		cell.SetString(piao.Gupiaojiancheng)
		cell = row.AddCell()
		cell.SetFloat(piao.Zhangfu)
		cell = row.AddCell()
		cell.SetFloat(piao.Huanshou)
		cell = row.AddCell()
		cell.SetFloat(piao.Liangbi)
		cell = row.AddCell()
		cell.SetFloat(piao.Jingzhanzuo)
		cell = row.AddCell()
		cell.SetFloat(piao.Zhulizengcang)
		cell = row.AddCell()
		cell.SetFloat(piao.Weipipei / 10000)
		cell = row.AddCell()
		cell.SetFloat(piao.Jingjiajine / 10000)
		cell = row.AddCell()
		cell.SetFloat(piao.Weipipei / piao.Jingjiajine)
		cell = row.AddCell()
		cell.SetString(piao.Shijian)
	}

	now := time.Now()
	todayStr := now.Format("20060102")
	newfilename := "data" + todayStr + ".xlsx"
	err := useexcel.SaveFile(sheet, newfilename)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return err
}
