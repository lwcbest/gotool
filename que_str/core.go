package que_str

import (
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
)

type Config struct {
	Input InputStruct
}

type InputStruct struct {
	Cookie string     `toml:"Cookie"`
	Reqs   []ReqQuery `toml:"reqs"`
}

type ReqQuery struct {
	Name string `toml:"name"`
	Ques string `toml:"ques"`
}

type RequestData struct {
	Question        string `json:"question"`
	Perpage         int    `json:"perpage"`
	Page            int    `json:"page"`
	SecondaryIntent string `json:"secondary_intent"`
	LogInfo         string `json:"log_info"`
	Iwcpro          int    `json:"iwcpro"`
	Source          string `json:"source"`
	Version         string `json:"version"`
	QueryArea       string `json:"query_area"`
	BlockList       string `json:"block_list"`
	AddInfo         string `json:"add_info"`
}

type ResData struct {
	StatusCode int `json:"status_code"`
	Data       struct {
		Answer []struct {
			Txt []struct {
				Content struct {
					Components []struct {
						Data struct {
							Datas []map[string]interface{} `json:"datas"`
						} `json:"data"`
					} `json:"components"`
				} `json:"content"`
			} `json:"txt"`
		} `json:"answer"`
	} `json:"data"`
}

func ReqStr() string {
	var conf Config
	if _, err := toml.DecodeFile("./data.toml", &conf); err != nil {
		fmt.Printf("fail to read config.||err=%v||config=%v", err, conf)
		os.Exit(1)
		return ""
	}

	result := ""
	for _, reqQuery := range conf.Input.Reqs {
		result += httpDo(conf, reqQuery)
	}
	return result
}

func httpDo(conf Config, reqQuery ReqQuery) string {
	client := &http.Client{}

	reqData := &RequestData{}
	reqData.AddInfo = "{\"urp\":{\"scene\":1,\"company\":1,\"business\":1},\"contentType\":\"json\",\"searchInfo\":true}"
	reqData.BlockList = ""
	reqData.QueryArea = ""
	reqData.Version = "2.0"
	reqData.Source = "Ths_iwencai_Xuangu"
	reqData.Iwcpro = 1
	reqData.LogInfo = "{\"input_type\":\"click\"}"
	reqData.SecondaryIntent = "stock"
	reqData.Page = 1
	reqData.Perpage = 50

	reqData.Question = reqQuery.Ques

	req, err := http.NewRequest("POST", "http://www.iwencai.com/customized/chart/get-robot-data", strings.NewReader(JSONStr(reqData)))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Cache-control", "no-cache")
	req.Header.Set("Cookie", conf.Input.Cookie)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	var resData ResData
	err = json.Unmarshal(body, &resData)
	if err != nil {
		return "更新配置文件"
	}
	log.Println(string(body))
	if resData.StatusCode == 0 {
		finalDatas := resData.Data.Answer[0].Txt[0].Content.Components[0].Data.Datas
		return BuildTable(reqQuery.Name, finalDatas)
		//return JSONStr(finalDatas)
	}

	return ""
}

func JSONStr(v interface{}) string {
	if v == nil {
		return ""
	}
	k := reflect.TypeOf(v).Kind()
	switch k {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Slice:
		if reflect.ValueOf(v).IsNil() {
			return ""
		}
	case reflect.String:
		return v.(string)
	}
	JSON, _ := json.Marshal(v)
	return string(JSON)
}

func DecodeJSON(str string, v interface{}) error {
	return json.Unmarshal([]byte(str), v)
}
