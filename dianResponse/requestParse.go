package dianResponse

import (
	"bufio"
	"encoding/json"
	"mydiantp/dianRequest"
	"mydiantp/myconsts"
	"strconv"
)


var requestRecord map[int64][][]byte = make(map[int64][][]byte)
var SessionCount int64 = 0

// This function can invoke corresponding functions to handle different requests.
func RequestHandler(r *dianRequest.DianRequest, writer *bufio.Writer, reader *bufio.Reader) bool {
	switch r.Method {

	case myconsts.OPTIONS:
		//requestRecord[sessionId] = append(requestRecord[sessionId], *r)
		options(writer, r.CSeq)
		return false

	case myconsts.SETUP:

		// 是否需要先检验token的存在
		if tokenValid(r.Body["token"]){
			data := ResponseSETUP{
				StatusCode: myconsts.StatusFail,
				StatusMsg:  "fail, have set up already",
				Version:    myconsts.DianVersion,
				CSeq:       r.CSeq,
				Token:      r.Body["token"],
			}
			write(writer, data)
			return false
		}
		setUp(r, writer)
		return false

	case myconsts.PLAY:
		token, exist := r.Body["token"]
		if !exist || !tokenValid(token) {
			data := ResponsePLAY{
				StatusCode: myconsts.StatusFail,
				StatusMsg:  "fail, please set up first",
				Version:    myconsts.DianVersion,
				CSeq:       r.CSeq,
			}
			write(writer, data)
			return false
		}
		jsonData, _ := json.Marshal(*r)
		requestRecord[allToken[token]] = append(requestRecord[allToken[token]], jsonData)
		ntp, _ := strconv.Atoi(r.Body["ntp"])
		play(writer, reader, r.CSeq, ntp)
		return false


	case myconsts.TEARDOWN:
		token, exist := r.Body["token"]
		if !exist || !tokenValid(token) {
			data := ResponseTEARDOWN{
				StatusCode: myconsts.StatusFail,
				StatusMsg:  "fail, please set up first",
				Version:    myconsts.DianVersion,
				CSeq:       r.CSeq,
			}
			write(writer, data)
			return false
		}
		jsonData, _ := json.Marshal(*r)
		requestRecord[allToken[token]] = append(requestRecord[allToken[token]], jsonData)
		teardown(allToken[token], writer, r.CSeq)
		return true
	}
	return true
}

func tokenValid(token string) bool {
	for key, _ := range allToken {
		if key == token {
			return true
		}
	}
	return false
}