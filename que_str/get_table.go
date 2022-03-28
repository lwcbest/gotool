package que_str

import (
	"fmt"
	"strconv"
)

func BuildTable(name string, dataResource []map[string]interface{}) string {
	tableStr := "<head>\n<title>新魔盒1.0</title>\n<meta charset=\"UTF-8\">\n<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n \n<style type=\"text/css\">\nhtml {\n    font-family: sans-serif;\n    -ms-text-size-adjust: 100%;\n    -webkit-text-size-adjust: 100%;\n}\n \nbody {\n    margin: 10px;\n}\ntable {\n    border-collapse: collapse;\n    border-spacing: 0;\n}\n \ntd,th {\n    padding: 0;\n}\n \n.pure-table {\n    border-collapse: collapse;\n    border-spacing: 0;\n    empty-cells: show;\n    border: 1px solid #cbcbcb;\n}\n \n.pure-table caption {\n    color: #000;\n    font: italic 85%/1 arial,sans-serif;\n    padding: 1em 0;\n    text-align: center;\n}\n \n.pure-table td,.pure-table th {\n    border-left: 1px solid #cbcbcb;\n    border-width: 0 0 0 1px;\n    font-size: inherit;\n    margin: 0;\n    overflow: visible;\n    padding: .5em 1em;\n}\n \n.pure-table thead {\n    background-color: #e0e0e0;\n    color: #000;\n    text-align: left;\n    vertical-align: bottom;\n}\n \n.pure-table td {\n    background-color: transparent;\n}\n \n.pure-table-odd td {\n    background-color: #f2f2f2;\n}\n</style>\n</head>"
	tableStr += "<body><table class=\"pure-table\">"
	//第一行名字
	tableStr += "<tr class=\"pure-table-odd\"><td>" + name + "</td></tr>"

	if len(dataResource) > 0 {
		keys := BuildKey(dataResource[0])

		//第二行表头
		tableStr += "<tr class=\"pure-table-odd\">"
		tableStr += "<td>" + "得分" + "</td>"
		for _, key := range keys {
			tableStr += "<td>"
			tableStr += key
			tableStr += "</td>"
		}
		tableStr += "</tr>"

		//第三行以后
		for _, item := range dataResource {
			tableStr += "<tr class=\"pure-table-odd\">"
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
