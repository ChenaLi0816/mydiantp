package dianRequest

import (
	"fmt"
	"mydiantp/myconsts"
	"strconv"
	"strings"
)

var cseqnum int64 = 0

func Pack(method string, url string, version string, body map[string]string) *DianRequest {
	data := &DianRequest{
		Method:  method,
		Url:     url,
		Version: version,
		CSeq:    cseqnum,
		Body:    body,
	}
	return data
}


// raw : ""
func ParseRaw(raw string, sessionId int64) *DianRequest{
	str := strings.Split(raw, " ")
	//method, _ := strconv.ParseInt(str[0], 10, 64)
	method, exist := myconsts.METHODS[str[0]]
	if !exist {
		fmt.Println("没有此方法，请检查")
		return nil
	}
	index := strings.Index(str[1], "//")
	pro := str[1][:index+2]
	fmt.Println("使用的协议：", pro)
	if pro != "diantp://" {
		fmt.Println("并非支持的diantp协议，请检查")
		return nil
	}
	str[1] = str[1][index+2:]
	var url string
	if strings.Contains(str[1], ":") {
		url = str[1]
	} else {
		url = str[1] + ":" + myconsts.DefaultPort
	}

	version := str[2]
	if version != myconsts.DianVersion {
		fmt.Println("并不支持的diantp协议版本，请检查")
		return nil
	}



	body := make(map[string]string)
	switch method {
	case myconsts.SETUP:
		if len(str) < 5 {
			fmt.Println("输入参数不够，请检查")
			return nil
		}
		body["transport"] = str[3]
		body["addr"] = str[4]
	case myconsts.PLAY:
		body["session_id"] = strconv.FormatInt(sessionId, 10)
	case myconsts.TEARDOWN:
		body["session_id"] = strconv.FormatInt(sessionId, 10)
	}

	cseqnum++
	return Pack(method, url, version, body)



}
