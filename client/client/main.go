package main

import (
	"bufio"
	"client/dianRequest"
	_ "client/dianResponse"
	"client/myconsts"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func main(){
	fmt.Println("客户端启动..")
	conn, err := net.Dial("tcp", myconsts.ServerAddr)
	if err != nil {
		fmt.Println("连接客户端失败..")
		return
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	var sessionId int64 = -1
	loop:
	for true {
		fmt.Println("请输入指令：")
		str := readline()
		data := dianRequest.ParseRaw(str, sessionId)
		if data == nil {
			continue
		}
		buf, err := json.Marshal(*data)
		if err != nil {
			fmt.Println("请求体序列化失败")
			continue
		}
		fmt.Println("请求为：", string(buf))
		if _, err0 := writer.Write(buf); err0 != nil {
			fmt.Println("写入错误：", err0)
			continue
		}
		writer.Flush()
		bufRead := make([]byte, 1024)
		n,_ := reader.Read(bufRead)
		//resp, _ := io.ReadAll(reader)
		fmt.Println("响应为：", string(bufRead[:n]))
		switch data.Method {
		case myconsts.OPTIONS:
			OptionsHandler(bufRead[:n])
		case myconsts.SETUP:
			sessionId = SetupHandler(bufRead[:n])
		case myconsts.PLAY:
			PlayHandler(bufRead[:n], reader, writer)
		case myconsts.TEARDOWN:
			if TeardownHandler(bufRead[:n]) {
				break loop
			}
		}
	}
	fmt.Println("已断开连接..")
}
//1 diantp://127.0.0.1:8080 0.5
func readline() string {
	reader := bufio.NewReader(os.Stdin)
	data , _ := reader.ReadString('\n')
	return data[:len(data)-1]
}