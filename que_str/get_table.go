package que_str

import (
	"fmt"
	"strconv"
)

func BuildTable(name string, dataResource []map[string]interface{}) string {
	tableStr := "<body><p>" + name + "</p><table class=\"pure-table\">"

	if len(dataResource) > 0 {
		keys := BuildKey(dataResource[0])

		//表头
		tableStr += " <thead><tr>"
		tableStr += "<td>" + "得分" + "</td>"
		for _, key := range keys {
			tableStr += "<td>"
			tableStr += key
			tableStr += "</td>"
		}
		tableStr += "</tr></thead>"

		//第三行以后
		for index, item := range dataResource {
			if index%2 == 0 {
				tableStr += "<tr class=\"pure-table-odd\">"
			} else {
				tableStr += "<tr>"
			}

			//计算得分
			score := ComputeScore(item)
			tableStr += "<td>" + strconv.Itoa(score) + "</td>"

			for _, key := range keys {
				value := item[key]
				tableStr += "<td>"
				_, ok := value.(float64)
				if ok {
					tableStr += fmt.Sprintf("%.2f", value)
				} else {
					tableStr += fmt.Sprintf("%v", value)
				}
				tableStr += "</td>"
			}
			tableStr += "</tr>"
		}
	} else {
		tableStr += "<tr><td>" + "无数据" + "</td></tr>"
	}

	tableStr += "</table></body>"

	return tableStr
}

func BuildKey(row map[string]interface{}) []string {
	keys := make([]string, 0)
	keys = append(keys, "code", "股票简称")
	for key, _ := range row {
		if key != "code" && key != "股票简称" {
			keys = append(keys, key)
		}
	}

	return keys
}
