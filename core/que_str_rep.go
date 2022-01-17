package core

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func ReqStr(){
	fmt.Println("请输入问句：")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	req:=input.Text()
	fmt.Println("你输入的问句是：", req)

	fmt.Println("请输入日期：")
	input.Scan()
	myDate:=input.Text()
	fmt.Println("你输入的日期是：", myDate)
	myDates:=strings.Split(myDate,".")
	year,_:=strconv.ParseInt(myDates[0],10,64)
	month,_:=strconv.ParseInt(myDates[1],10,64)
	day,_:=strconv.ParseInt(myDates[2],10,64)
	loc, _ := time.LoadLocation("Local")
	today:=time.Date(int(year),time.Month(month),int(day),0,0,0,0,loc)
	yesterday := today.Add(-24*time.Hour)

	yesterdayWeek:=yesterday.Weekday()
	if yesterdayWeek==time.Sunday{
		yesterday = yesterday.Add(-24*time.Hour)
		yesterday = yesterday.Add(-24*time.Hour)
	}

	todayStr:=strconv.Itoa(today.Year())+"年"+strconv.Itoa(int(today.Month()))+"月"+strconv.Itoa(today.Day())+"日"
	yesterdayStr:=strconv.Itoa(yesterday.Year())+"年"+strconv.Itoa(int(yesterday.Month()))+"月"+strconv.Itoa(yesterday.Day())+"日"
	fmt.Println("今天是：", todayStr)
	fmt.Println("昨天是：", yesterdayStr)

	str2 := strings.Replace(req, "今日", todayStr, -1)
	str3 := strings.Replace(str2, "昨日", yesterdayStr, -1)

	fmt.Println("替换后问句是：", str3)
}
