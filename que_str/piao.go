package que_str

import "strconv"

type Piao struct {
	Code             string
	Gupiaojiancheng  string
	Zuixinzhangdiefu float64
	Huanshou         float64
	Liangbi          float64
	Jingzhanzuo      float64
	Weipipei         float64
	Jingjiajine      float64
	Zhulizengcang    float64
	Zhangfu          float64
	Weibi            float64
	Jiban            float64
	Shijian          string

	HuanshouScore      float64
	LiangbiScore       float64
	JingzhanzuoScore   float64
	WeipipeiScore      float64
	JingjiajineScore   float64
	ZhulizengcangScore float64
	ZhangfuScore       float64
	TotalScore         float64
}

func (p *Piao) Input(row map[string]interface{}) {
	p.Code = row["code"].(string)
	p.Gupiaojiancheng = row[GetKeyName("股票简称")].(string)
	p.Zuixinzhangdiefu, _ = strconv.ParseFloat(row[GetKeyName("最新涨跌幅")].(string), 64)
	p.Huanshou = row[GetKeyName("换手")].(float64)
	p.Liangbi, _ = strconv.ParseFloat(row[GetKeyName("量比")].(string), 64)
	p.Jingzhanzuo = row[GetKeyName("竞占昨")].(float64)
	p.Weipipei = row[GetKeyName("未匹配")].(float64)
	p.Jingjiajine = row[GetKeyName("竞价金额")].(float64)
	if row[GetKeyName("主力增仓")] != nil {
		p.Zhulizengcang = row[GetKeyName("主力增仓")].(float64)
	}

	if row[GetKeyName("涨幅")] != nil {
		p.Zhangfu, _ = strconv.ParseFloat(row[GetKeyName("涨幅")].(string), 64)
	}
	if row[GetKeyName("委比")] != nil {
		p.Weibi, _ = strconv.ParseFloat(row[GetKeyName("委比")].(string), 64)
	}
	if row[GetKeyName("几板")] != nil {
		p.Jiban = row[GetKeyName("几板")].(float64)
	}
	if row[GetKeyName("最终涨停时间")] != nil {
		p.Shijian = row[GetKeyName("最终涨停时间")].(string)
	}
}

func (p *Piao) GetTotalScore() float64 {
	if p.TotalScore == 0 {
		p.TotalScore = p.HuanshouScore + p.LiangbiScore + p.JingzhanzuoScore + p.WeipipeiScore + p.ZhulizengcangScore + p.JingjiajineScore + p.ZhangfuScore
	}

	return p.TotalScore
}
