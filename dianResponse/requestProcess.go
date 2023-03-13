package dianResponse

import (
	"bufio"
	"encoding/json"
	"fmt"
	"mydiantp/dianRequest"
	"mydiantp/myconsts"
)

func setUp(r *dianRequest.DianRequest, writer *bufio.Writer) {
	port := r.Body["client_port"]
	network := r.Body["transport"]
	fmt.Printf("连接到%v端口,使用%v协议\n", port, network)

	//dest, err := net.Dial(network, fmt.Sprintf("%v:%v", addr, port))
	//if err != nil {
	//	fmt.Println("连接失败：", err)
	//	data := ResponseSETUP{
	//		StatusCode: myconsts.StatusFail,
	//		StatusMsg:  "link fail",
	//		Version:    myconsts.DianVersion,
	//		CSeq:       r.CSeq,
	//		SessionId:  -1,
	//	}
	//	write(writer, data)
	//	return
	//}
	//defer dest.Close()

	jsonData, _ := json.Marshal(*r)
	sessionCount++
	requestRecord[sessionCount] = append(requestRecord[sessionCount], jsonData)
	data := ResponseSETUP{
		StatusCode: myconsts.StatusOK,
		StatusMsg:  "OK",
		Version:    myconsts.DianVersion,
		CSeq:       r.CSeq,
		SessionId:  sessionCount,
	}
	write(writer, data)
	return

}

func options(writer *bufio.Writer, CSeq int64){
	data := ResponseOPTIONS{
		StatusCode: myconsts.StatusOK,
		StatusMsg:  "OK",
		Version:    myconsts.DianVersion,
		CSeq:       CSeq,
		Method:     myconsts.METHODS,
	}
	write(writer, data)
}

func play(writer *bufio.Writer, reader *bufio.Reader, CSeq int64){
	data := ResponsePLAY{
		StatusCode: myconsts.StatusOK,
		StatusMsg:  "OK",
		Version:    myconsts.DianVersion,
		CSeq:       CSeq,
	}
	write(writer, data)


	relay, _ := reader.ReadByte()
	if relay != '1' {
		fmt.Println("未收到确认报文")
		return
	}
	fmt.Println("收到确认报文，准备发送视频数据..")


	videoByte := []byte{48,49,50,51,52,53,54,55,56,57,58,59,60}
	writer.Write(videoByte)
	writer.Flush()
	fmt.Println("已发送视频数据")

}

func teardown(sessionId int64, writer *bufio.Writer, CSeq int64){

	data := &ResponseTEARDOWN{
		StatusCode: myconsts.StatusOK,
		StatusMsg:  "OK",
		Version:    myconsts.DianVersion,
		CSeq:       CSeq,
	}
	write(writer, data)

	fmt.Printf("会话id:%v，共发送了%v个请求，分别如下\n", sessionId, len(requestRecord[sessionId]))
	for index,value := range requestRecord[sessionId] {
		fmt.Printf("第%v个请求为：%v\n", index+1, string(value))
	}


}

func isSetUp(sessionId int64) bool{
	if sessionId == -1{
		//fmt.Println("尚未建立起连接")
		return false
	}
	return true
}

func write(writer *bufio.Writer, resp any){
	respJson, errJson := json.Marshal(resp)
	if errJson != nil {
		fmt.Println("序列化数据错误：", errJson)
		return
	}
	if _, err := writer.Write(respJson);err != nil {
		fmt.Println("写入数据错误：", err)
		return
	}
	writer.Flush()
}