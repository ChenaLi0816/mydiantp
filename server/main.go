package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"mydiantp/dianRequest"
	"mydiantp/dianResponse"
	"mydiantp/myconsts"
	"net"
	"strings"
)

func main(){
	fmt.Println("服务端监听启动...监听地址为：", myconsts.ServerAddr)
	server, err := net.Listen("tcp", myconsts.ServerAddr)
	if err != nil {
		fmt.Println("监听失败,错误：", err)
	}

	client, err := server.Accept()
	fmt.Println("已监听到客户端,地址为：", client.RemoteAddr())
	process(client)
	fmt.Println("结束监听")
}


func process(conn net.Conn){
	defer conn.Close()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	var sessionId int64 = -1
	for true {
		var buf []byte = make([]byte, 1024)
		//length, errRead := conn.Read(buf)
		//buf, errRead :=  io.ReadAll(reader)
		//n, errRead := io.ReadFull(reader, buf)
		n, errRead := reader.Read(buf)
		if errRead != nil {
			fmt.Println("读取数据错误：", errRead)
			return
		}
		//fmt.Println("开始读取")

		fmt.Println("读取到的数据为：", string(buf))


		var req dianRequest.DianRequest
		if errJson := json.Unmarshal(buf[:n], &req);errJson != nil {
			fmt.Println("解析数据错误：", errJson)
			return
		}
		parseJson(buf[:n])
		//fmt.Println("解析后的请求为：",req)
		shutdown := dianResponse.RequestHandler(&req, sessionId, writer, reader)
		writer.Flush()
		if shutdown == 0 {
			break
		}

		sessionId = shutdown
	}
}

func parseJson(buf []byte) {
	fmt.Println("请求为：")
	str := string(buf[1:len(buf)-1])
	field := strings.SplitN(str, ",", 5)
	for _, value := range field {
		v := strings.SplitN(value, ":", 2)
		v[0] = strings.Trim(v[0], "\"")
		v[1] = strings.Trim(v[1], "\"")
		fmt.Printf("%v %v\t\n", v[0], v[1])
	}
}