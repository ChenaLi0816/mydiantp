package dianResponse

import (
	"bufio"
	"mydiantp/myconsts"
)

func ResponseHandler(method string, bufRead []byte, reader *bufio.Reader, writer *bufio.Writer) string {
	switch method {
	case myconsts.OPTIONS:
		OptionsHandler(bufRead)
	case myconsts.SETUP:
		return SetupHandler(bufRead)
	case myconsts.PLAY:
		PlayHandler(bufRead, reader, writer)
	case myconsts.TEARDOWN:
		if TeardownHandler(bufRead) {
			return "0"
		}
	}
	return "-1"
}
