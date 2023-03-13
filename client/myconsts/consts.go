package myconsts

const (
	OPTIONS = "OPTIONS"
	SETUP = "SETUP"
	PLAY = "PLAY"
	TEARDOWN = "TEARDOWN"
	StatusOK = 200
	StatusFail = 400
	DianVersion = "0.5"
	DefaultPort = "8080"
	ServerAddr = "127.0.0.1:8080"
)

var METHODS map[string]string = map[string]string{
	"OPTIONS" : "OPTIONS",
	"SETUP" : "SETUP",
	"PLAY" : "PLAY",
	"TEARDOWN" : "TEARDOWN",
}
