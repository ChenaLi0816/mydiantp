package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"mydiantp/dianRequest"
	"mydiantp/dianResponse"
	"mydiantp/myconsts"
	"net"
	"os"
)

func main() {
	fmt.Println("客户端启动..")
	conn, err := net.Dial("tcp", myconsts.ServerAddr)
	if err != nil {
		fmt.Println("连接服务端失败..")
		return
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	var token string = "-1"
	for true {
		fmt.Println("请输入指令：")
		str := readline()
		data := dianRequest.ParseRaw(str, token)
		if data == nil {
			continue
		}
		buf, err := json.Marshal(*data)
		if err != nil {
			fmt.Println("请求体序列化失败")
			continue
		}

		if _, err0 := writer.Write(buf); err0 != nil {
			fmt.Println("写入错误：", err0)
			continue
		}
		writer.Flush()
		bufRead := make([]byte, 1024)
		n, _ := reader.Read(bufRead)
		bufRead = bufRead[:n]
		if tmp := dianResponse.ResponseHandler(data.Method, bufRead, reader, writer); tmp != "-1" {
			token = tmp
		}

		if token == "0" {
			break
		}
	}
	fmt.Println("已断开连接..")
}

// read line without '\n'
func readline() string {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	return data[:len(data)-1]
}
