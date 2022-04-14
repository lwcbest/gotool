package que_str

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"strings"
)

type Config struct {
	Input InputStruct
}

type InputStruct struct {
	Cookie      string      `toml:"Cookie"`
	Reqs        []ReqQuery  `toml:"reqs"`
	ComputeData ComputeData `toml:"computedata"`
}

type ComputeData struct {
	Huanshou    [][]float64 `toml:"huanshou"`
	Liangbi     [][]float64 `toml:"liangbi"`
	Jingzhanzuo [][]float64 `toml:"jingzhanzuo"`
	Weipipei    [][]float64 `toml:"weipipei"`
	Jingjiajine [][]float64 `toml:"jingjiajine"`
	Zhuli       [][]float64 `toml:"zhuli"`
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

	result := "<head>\n<title>新魔盒1.8</title>\n<meta charset=\"UTF-8\">\n<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n \n<style type=\"text/css\">\nhtml {\n    font-family: sans-serif;\n    -ms-text-size-adjust: 100%;\n    -webkit-text-size-adjust: 100%;\n}\n \nbody {\n    margin: 10px;\n}\ntable {\n    border-collapse: collapse;\n    border-spacing: 0;\n}\n \ntd,th {\n    padding: 0;\n}\n \n.pure-table {\n    border-collapse: collapse;\n    border-spacing: 0;\n    empty-cells: show;\n    border: 1px solid #cbcbcb;\n}\n \n.pure-table caption {\n    color: #000;\n    font: italic 85%/1 arial,sans-serif;\n    padding: 1em 0;\n    text-align: center;\n}\n \n.pure-table td,.pure-table th {\n    border-left: 1px solid #cbcbcb;\n    border-width: 0 0 0 1px;\n    font-size: inherit;\n    margin: 0;\n    overflow: visible;\n    padding: .5em 1em;\n}\n \n.pure-table thead {\n    background-color: #e0e0e0;\n    color: #000;\n    text-align: left;\n    vertical-align: bottom;\n}\n \n.pure-table td {\n    background-color: transparent;\n}\n \n.pure-table-odd td {\n    background-color: #f2f2f2;\n}\n</style>\n</head>"
	conf.Input.Cookie = conf.Input.Cookie + "v=" + getV()
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
		fmt.Println(err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Cache-control", "no-cache")
	req.Header.Set("Cookie", conf.Input.Cookie)

	resp, err := client.Do(req)
	if err != nil {
		// handle error
		fmt.Println(err.Error())
	}

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

	if resData.StatusCode == 0 {
		finalDatas := resData.Data.Answer[0].Txt[0].Content.Components[0].Data.Datas
		return BuildTable(reqQuery.Name, finalDatas, conf.Input.ComputeData)
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

func getV() string {
	cmd := exec.Command("node", "./usejs/abc.js")
	//cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	v := out.String()
	v = v[:len(v)-1]
	return v
}
