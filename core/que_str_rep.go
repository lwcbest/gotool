package core

import (
	"bufio"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/golang-module/carbon"
	"os"
	"strings"
)

type Config struct {
	Input InputStruct
}

type InputStruct struct {
	Req         string
	TodayString string
	Count       int
}

func ReqStr() {
	var conf Config
	if _, err := toml.DecodeFile("./data.toml", &conf); err != nil {
		fmt.Printf("fail to read config.||err=%v||config=%v", err, conf)
		os.Exit(1)
		return
	}

	input := bufio.NewScanner(os.Stdin)

	req := conf.Input.Req
	myDate := conf.Input.TodayString
	today := carbon.Parse(myDate)
	if !today.IsWeekday() {
		fmt.Println("配置日期今日不能是周末....")
		input.Scan()
		input.Text()
		return
	}

	for i := 0; i < conf.Input.Count; i++ {
		today = BuildFinalStr(today, req)
	}

	fmt.Println("输入任何字母进行退出.....")
	input.Scan()
	input.Text()
}

func BuildFinalStr(today carbon.Carbon, req string) carbon.Carbon {
	yesterday := today.SubDay()
	if yesterday.IsWeekend() {
		yesterday = yesterday.SubDay()
		yesterday = yesterday.SubDay()
	}

	todayStr := today.ToFormatString("Ymd")
	yesterdayStr := yesterday.ToFormatString("Ymd")

	str2 := strings.Replace(req, "今日", todayStr, -1)
	str3 := strings.Replace(str2, "昨日", yesterdayStr, -1)

	fmt.Println(todayStr)
	fmt.Println(str3)
	fmt.Println("----------------")
	return yesterday
}
