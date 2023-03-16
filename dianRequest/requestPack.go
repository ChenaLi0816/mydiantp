package dianRequest

import (
	"fmt"
	"mydiantp/myconsts"
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

func ParseRaw(raw string, token string) *DianRequest {
	raw = strings.Trim(raw, "\r")
	str := strings.Split(raw, " ")
	if len(str) < 2 {
		fmt.Println("参数不够，请检查")
		return nil
	}

	_, exist := myconsts.METHODS[str[0]]
	if !exist {
		fmt.Println("没有此方法，请检查")
		return nil
	}

	method := str[0]
	if str[1] == "-d" {
		str = append(str, "take a place")
		tmp := []string{method, myconsts.DefaultPro + myconsts.ServerAddr, myconsts.DianVersion}
		tmp = append(tmp, str[2:len(str)]...)
		str = tmp[:len(tmp)-1]
	}
	if len(str) < 3 {
		fmt.Println("参数不够或第二个参数错误，请检查")
		return nil
	}
	index := strings.Index(str[1], "//")
	if index == -1 {
		fmt.Println("并非url格式，请检查")
		return nil
	}
	pro := str[1][:index+2]
	if pro != myconsts.DefaultPro {
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
	if url != myconsts.ServerAddr {
		fmt.Println("并未开启监听的客户端，请检查")
		return nil
	}
	version := str[2]
	if version != myconsts.DianVersion {
		fmt.Println("并不支持的diantp协议版本，请检查")
		return nil
	}

	body := make(map[string]string)
	switch method {
	case myconsts.OPTIONS:
		body["token"] = token
	case myconsts.SETUP:
		if len(str) < 5 {
			fmt.Println("输入参数不够，请检查")
			return nil
		}
		if str[3] != "tcp" && str[3] != "udp" {
			fmt.Println("传输协议不合法，请检查")
			return nil
		}
		body["token"] = token
		body["transport"] = str[3]
		body["client_port"] = str[4]
	case myconsts.PLAY:
		if len(str) < 4 {
			fmt.Println("输入参数不够，请检查")
			return nil
		}
		body["token"] = token
		body["ntp"] = str[3]
	case myconsts.TEARDOWN:
		body["token"] = token
	}

	cseqnum++
	return Pack(method, url, version, body)

}
