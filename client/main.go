package main

import (
	"bufio"
	"client/dianRequest"
	"client/dianResponse"
	_ "client/dianResponse"
	"client/myconsts"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
)

func main(){
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
		//fmt.Println("请求为：", string(buf))
		if _, err0 := writer.Write(buf); err0 != nil {
			fmt.Println("写入错误：", err0)
			continue
		}
		writer.Flush()
		bufRead := make([]byte, 1024)
		n,_ := reader.Read(bufRead)
		bufRead = bufRead[:n]

		parseJson(bufRead)


		if tmp := dianResponse.ResponseHandler(data.Method, bufRead, reader, writer); tmp != "-1" {
			token = tmp
		}

		if token == "0" {
			break
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

func parseJson(buf []byte) {
	fmt.Println("响应为：")
	str := string(buf[1:len(buf)-1])
	field := strings.SplitN(str, ",", 5)
	for index, value := range field {
		if index == 4 {
			break
		}
		v := strings.SplitN(value, ":", 2)
		v[0] = strings.Trim(v[0], "\"")
		v[1] = strings.Trim(v[1], "\"")
		fmt.Printf("%v %v\t\n", v[0], v[1])
	}
}