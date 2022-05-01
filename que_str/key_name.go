package que_str

import (
	"time"
)

func GetKeyName(key string) string {
	now := time.Now()
	todayStr := now.Format("20060102")
	var yesterStr string
	if now.Weekday() == time.Monday {
		yesterStr = now.AddDate(0, 0, -3).Format("20060102")
	} else {
		yesterStr = now.AddDate(0, 0, -1).Format("20060102")
	}

	keyMap := map[string]string{
		"主力增仓":  "主力增仓占比" + "[" + yesterStr + "]",
		"换手":    "竞价实际换手率" + "[" + todayStr + "]",
		"未匹配":   "竞价未匹配金额" + "[" + todayStr + "]",
		"竞价金额":  "竞价金额" + "[" + todayStr + "]",
		"几板":    "连续涨停天数" + "[" + yesterStr + "]",
		"量比":    "分时量比" + "[" + todayStr + " 09:25" + "]",
		"委比":    "分时委比" + "[" + todayStr + " 09:25" + "]",
		"竞占昨":   "竞价金额占昨日成交额" + "[" + todayStr + "]",
		"涨幅":    "分时涨跌幅:前复权" + "[" + todayStr + " 09:25" + "]",
		"code":  "code",
		"股票简称":  "股票简称",
		"最新涨跌幅": "最新涨跌幅",
	}

	return keyMap[key]
}

func GetDisplayKeys() []string {
	return []string{"code", "name", "latest", "huanshou", "liangbi", "jingzhanzuo", "weipipei", "jingjiajine", "weipipeizhanbi", "zhulizengcang", "zhangfu", "weibi", "jiban"}
}

func GetKeys() []string {
	keys := make([]string, 0)
	for _, disKey := range GetDisplayKeys() {
		key := GetKeyName(disKey)
		keys = append(keys, key)
	}
	return keys
}
