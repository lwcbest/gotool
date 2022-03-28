package que_str

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func ComputeScore(row map[string]interface{}) int {
	now := time.Now()
	todayStr := now.Format("20060102")
	var yesterStr string
	if now.Weekday() == time.Monday {
		yesterStr = now.AddDate(0, 0, -3).Format("20060102")
	} else {
		yesterStr = now.AddDate(0, 0, -1).Format("20060102")
	}

	code := row["code"]

	v_huanshou, ok := row["竞价实际换手率"+"["+todayStr+"]"]
	if !ok {
		fmt.Println("没有key!" + "竞价实际换手率" + "[" + todayStr + "]")
		return -999
	}
	v_jingjiajine, ok := row["竞价金额"+"["+todayStr+"]"]
	if !ok {
		fmt.Println("没有key!" + "竞价金额" + "[" + todayStr + "]")
		return -999
	}
	v_jiban, ok := row["连续涨停天数"+"["+yesterStr+"]"]
	if !ok {
		fmt.Println("没有key!" + "连续涨停天数" + "[" + yesterStr + "]")
		return -999
	}
	v_jingjiaweipipeijine, ok := row["竞价未匹配金额"+"["+todayStr+"]"]
	if !ok {
		fmt.Println("没有key!" + "竞价未匹配金额" + "[" + todayStr + "]")
		return -999
	}
	v_fenshiliangbi, ok := row["分时量比"+"["+todayStr+" 09:25"+"]"]
	if !ok {
		fmt.Println("没有key!" + "分时量比" + "[" + todayStr + " 09:25" + "]")
		return -999
	}

	v_jingzhanzuo, ok := row["竞价金额占昨日成交额"+"["+todayStr+"]"]
	if !ok {
		fmt.Println("没有key!" + "竞价金额占昨日成交额" + "[" + todayStr + "]")
		return -999
	}

	v_zhangfu, ok := row["分时涨跌幅:前复权"+"["+todayStr+" 09:25"+"]"]
	if !ok {
		fmt.Println("没有key!" + "分时涨跌幅:前复权" + "[" + todayStr + " 09:25" + "]")
		return -999
	}

	score := 0
	huanshouScore := computeHuanshou(v_huanshou.(float64))
	jingjiajineScore := computeJine(v_jingjiajine.(float64))
	jibanScore := computeJiban(v_jiban.(float64))
	weipipeiScore := computeWeipipei(v_jingjiaweipipeijine.(float64), v_jingjiajine.(float64))

	v_fenshiliangbiStr := v_fenshiliangbi.(string)
	v_fenshiliangbifloat, _ := strconv.ParseFloat(v_fenshiliangbiStr, 64)
	liangbiScore := computeLiangbi(v_fenshiliangbifloat)

	jingzhanzuoScore := computeJingzhanzuo(v_jingzhanzuo.(float64))

	v_zhangfuStr := v_zhangfu.(string)
	v_zhangfufloat, _ := strconv.ParseFloat(v_zhangfuStr, 64)
	zhangfuScore := computeZhangfu(v_zhangfufloat)

	score = huanshouScore + jingjiajineScore + jibanScore + weipipeiScore + liangbiScore + jingzhanzuoScore + zhangfuScore
	fmt.Printf("%v得分%v,换手%v,竞价金额%v,几板%v,未匹配%v,量比%v,竞占昨%v,涨幅%v \n", code, score, huanshouScore, jingjiajineScore, jibanScore, weipipeiScore, liangbiScore, jingzhanzuoScore, zhangfuScore)
	return score
}

func computeHuanshou(v_huanshou float64) int {
	if v_huanshou >= 1 {
		return 2
	}

	if v_huanshou >= 0.8 && v_huanshou < 1 {
		return -2
	}

	if v_huanshou >= 0.6 && v_huanshou < 0.8 {
		return -4
	}

	if v_huanshou < 0.6 {
		return -6
	}

	return 0
}

func computeJine(v_jingjiajine float64) int {
	score := 0
	w := 10000.0
	if v_jingjiajine < 900.0*w {
		return score - 2
	}

	if v_jingjiajine > 1000*w {
		left := v_jingjiajine - 1000*w
		down := 500 * w
		temp := math.Floor(left / down)
		score += int(temp)
	}

	return score
}

func computeJiban(v_jiban float64) int {
	if v_jiban == 0 {
		return 0
	} else {
		score := -1
		return score - int(v_jiban)
	}
}

func computeWeipipei(v_weipipei float64, v_jingjiajine float64) int {
	w := 10000.0
	if v_weipipei > 0 {
		return 2
	}

	if v_weipipei > -1*w && v_weipipei < 0 {
		return 0
	}

	v_weipipei_abs := math.Abs(v_weipipei)
	bi := v_weipipei_abs / v_jingjiajine
	if bi < 0.05 {
		return -1
	}

	if bi >= 0.05 && bi < 0.1 {
		return -2
	}

	if bi >= 0.1 && bi < 0.15 {
		return -3
	}

	if bi >= 0.15 && bi < 0.2 {
		return -4
	}
	if bi >= 0.2 {
		return -5
	}

	return 0
}

func computeLiangbi(v_liangbi float64) int {
	if v_liangbi < 9 {
		return -2
	}

	if v_liangbi >= 60 && v_liangbi < 80 {
		return -2
	}

	if v_liangbi > 80 {
		return -4
	}

	return 0
}

func computeJingzhanzuo(v_jingzhanzuo float64) int {
	if v_jingzhanzuo < 0.03 {
		return -2
	}

	return 0
}

func computeZhangfu(v_zhangfu float64) int {
	if v_zhangfu < 0.02 {
		return 0
	}

	if v_zhangfu >= 0.02 && v_zhangfu < 0.04 {
		return 2
	}

	if v_zhangfu >= 0.04 && v_zhangfu < 0.06 {
		return 3
	}

	return 0
}
