package main

import (
	"bufio"
	"client/dianResponse"
	"client/myconsts"
	"encoding/json"
	"fmt"
)


func OptionsHandler(buf []byte){
	resp := dianResponse.ResponseOPTIONS{}
	if !unmarshal(buf, &resp) {
		return
	}
	if resp.StatusCode != myconsts.StatusOK {
		fmt.Println("响应错误，错误信息为：", resp.StatusMsg)
		return
	}
	fmt.Println("响应成功！")
	fmt.Printf("共有方法%v种，分别为\n", len(resp.Method))
	for key, value := range resp.Method {
		fmt.Printf("%v:%v\n", key, value)
	}
}

func SetupHandler(buf []byte) int64 {
	resp := dianResponse.ResponseSETUP{}
	if !unmarshal(buf, &resp) {
		return -1
	}
	if resp.StatusCode != myconsts.StatusOK {
		fmt.Println("响应错误，错误信息为：", resp.StatusMsg)
		return -1
	}
	fmt.Println("响应成功！")
	fmt.Printf("已建立连接，会话ID为：%v\n", resp.SessionId)
	return resp.SessionId
}

func PlayHandler(buf []byte, reader *bufio.Reader, writer *bufio.Writer) {
	resp := dianResponse.ResponsePLAY{}
	if !unmarshal(buf, &resp) {
		return
	}
	if resp.StatusCode != myconsts.StatusOK {
		fmt.Println("响应错误，错误信息为：", resp.StatusMsg)
		return
	}
	fmt.Println("响应成功！")
	acp := []byte{'1'}
	writer.Write(acp)
	writer.Flush()
	fmt.Println("已回复确认报文")
	fmt.Println("开始读取字节流...")
	bufRead := make([]byte, 1024)
	n, err := reader.Read(bufRead)
	if err != nil {
		fmt.Println("读取错误：", err)
		return
	}
	fmt.Println("读取到的字节流为：", bufRead[:n])
	return

}
func TeardownHandler(buf []byte) bool {
	resp := dianResponse.ResponseTEARDOWN{}
	if !unmarshal(buf, &resp) {
		return false
	}
	if resp.StatusCode != myconsts.StatusOK {
		fmt.Println("响应错误，错误信息为：", resp.StatusMsg)
		return false
	}
	fmt.Println("响应成功！")
	fmt.Println("即将关闭连接...")
	return true
}

func unmarshal(buf []byte, dest any) bool{
	if err := json.Unmarshal(buf, dest); err != nil {
		fmt.Println("解析数据错误：", err)
		return false
	}
	return true
}
