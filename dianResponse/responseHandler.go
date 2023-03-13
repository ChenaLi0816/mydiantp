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

// This function can invoke corresponding functions to handle requests.
func RequestHandler(r *dianRequest.DianRequest, sessionId int64, writer *bufio.Writer, reader *bufio.Reader) int64 {
	switch r.Method {

	case myconsts.OPTIONS:
		//requestRecord[sessionId] = append(requestRecord[sessionId], *r)
		options(writer, r.CSeq)
		return sessionId

	case myconsts.SETUP:
		if isSetUp(sessionId){
			data := ResponseSETUP{
				StatusCode: myconsts.StatusFail,
				StatusMsg:  "fail, have set up already",
				Version:    myconsts.DianVersion,
				CSeq:       r.CSeq,
				SessionId:  sessionId,
			}
			write(writer, data)
			return sessionId
		}
		setUp(r, writer)
		return SessionCount

	case myconsts.PLAY:
		if !isSetUp(sessionId) {
			data := ResponsePLAY{
				StatusCode: myconsts.StatusFail,
				StatusMsg:  "fail, please set up first",
				Version:    myconsts.DianVersion,
				CSeq:       r.CSeq,
			}
			write(writer, data)
			return sessionId
		}
		jsonData, _ := json.Marshal(*r)
		requestRecord[sessionId] = append(requestRecord[sessionId], jsonData)
		ntp, _ := strconv.Atoi(r.Body["ntp"])
		play(writer, reader, r.CSeq, ntp)
		return sessionId


	case myconsts.TEARDOWN:
		if !isSetUp(sessionId) {
			data := ResponseTEARDOWN{
				StatusCode: myconsts.StatusOK,
				StatusMsg:  "OK",
				Version:    myconsts.DianVersion,
				CSeq:       r.CSeq,
			}
			write(writer, data)
			return 0
		}
		jsonData, _ := json.Marshal(*r)
		requestRecord[sessionId] = append(requestRecord[sessionId], jsonData)
		teardown(sessionId, writer, r.CSeq)
		return 0

	}
	return sessionId
}

