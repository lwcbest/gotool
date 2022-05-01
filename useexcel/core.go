package useexcel

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func ReadExcelSheet(filename string, SheetNum int) *xlsx.Sheet {
	file, err := xlsx.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	sheet := file.Sheets[SheetNum]
	return sheet
}

func SaveFile(sheet *xlsx.Sheet, filename string) error {
	return sheet.File.Save(filename)
}
