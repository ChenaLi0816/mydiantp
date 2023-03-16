package dianRequest

import (
	"bufio"
	"mydiantp/dianResponse"
	"mydiantp/myconsts"
	"strconv"
	"strings"
)

var requestRecord map[int64][]*DianRequest = make(map[int64][]*DianRequest)
var SessionCount int64 = 0

// This function can invoke corresponding functions to handle different requests.
func RequestHandler(r *DianRequest, writer *bufio.Writer, reader *bufio.Reader, addr string) bool {

	switch r.Method {

	case myconsts.OPTIONS:
		token, exist := r.Body["token"]
		if exist && tokenValid(token) {
			requestRecord[allToken[token]] = append(requestRecord[allToken[token]], r)
		}

		options(writer, r.CSeq)
		return false

	case myconsts.SETUP:
		token, exist := r.Body["token"]
		if exist && tokenValid(r.Body["token"]) {
			requestRecord[allToken[token]] = append(requestRecord[allToken[token]], r)
			data := dianResponse.ResponseSETUP{
				StatusCode: myconsts.StatusFail,
				StatusMsg:  "fail, have set up already",
				Version:    myconsts.DianVersion,
				CSeq:       r.CSeq,
				Token:      r.Body["token"],
			}
			write(writer, data)
			return false
		}
		str := strings.Split(addr, ":")
		r.Body["addr"] = str[0] + ":" + r.Body["client_port"]
		setUp(r, writer)
		return false

	case myconsts.PLAY:
		token, exist := r.Body["token"]
		if !exist || !tokenValid(token) {
			data := dianResponse.ResponsePLAY{
				StatusCode: myconsts.StatusFail,
				StatusMsg:  "fail, please set up first",
				Version:    myconsts.DianVersion,
				CSeq:       r.CSeq,
			}
			write(writer, data)
			return false
		}
		requestRecord[allToken[token]] = append(requestRecord[allToken[token]], r)
		ntp, _ := strconv.Atoi(r.Body["ntp"])
		play(writer, reader, r.CSeq, ntp)
		return false

	case myconsts.TEARDOWN:
		token, exist := r.Body["token"]
		if !exist || !tokenValid(token) {
			data := dianResponse.ResponseTEARDOWN{
				StatusCode: myconsts.StatusOK,
				StatusMsg:  "OK",
				Version:    myconsts.DianVersion,
				CSeq:       r.CSeq,
			}
			write(writer, data)
			return true
		}
		requestRecord[allToken[token]] = append(requestRecord[allToken[token]], r)
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
