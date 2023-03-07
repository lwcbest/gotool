package core

import (
	"fmt"
	"github.com/tidwall/gjson"
	"os"
	"strings"
)

func GenerateSql(rawFileName string, targetFileName string) {
	tableName := "openplatform_item_upload"

	bytes, err := os.ReadFile(rawFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	result := gjson.Parse(string(bytes))
	columnsResult := result.Get("data.0.data.0.columns")
	rowsResult := result.Get("data.0.data.0.rows")
	fmt.Println(columnsResult)
	fmt.Println(rowsResult)
	sqls := make([]string, 0)
	if rowsResult.IsArray() {
		rowsArrayResult := rowsResult.Array()
		for _, row := range rowsArrayResult {
			fieldsResults := row.Get("result").Array()

			sql := fmt.Sprintf("INSERT INTO %s (", tableName)
			if columnsResult.IsArray() {
				columnsArrayResult := columnsResult.Array()
				for i, column := range columnsArrayResult {
					if i == len(columnsArrayResult)-1 {
						sql += column.String() + ")"
					} else {
						sql += column.String() + ", "
					}
				}
			} else {
				return
			}

			sql += "VALUES ("

			for i, field := range fieldsResults {
				if i == len(fieldsResults)-1 {
					sql += field.Raw + ");"
				} else {
					sql += field.Raw + ", "
				}
			}
			fmt.Println(sql)
			sqls = append(sqls, sql)
		}
	}

	file, err := os.Create(targetFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	finalDataString := strings.Join(sqls, "\n")
	file.Write([]byte(finalDataString))
}
