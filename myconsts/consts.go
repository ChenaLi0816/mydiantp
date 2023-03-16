package myconsts

const (
	OPTIONS     = "OPTIONS"
	SETUP       = "SETUP"
	PLAY        = "PLAY"
	TEARDOWN    = "TEARDOWN"
	StatusOK    = 200
	StatusFail  = 400
	DianVersion = "0.5"
	DefaultPort = "8080"
	ServerAddr  = "127.0.0.1:8080"
	DefaultPro  = "diantp://"
)

var METHODS map[string]string = map[string]string{
	"OPTIONS":  "查看方法，-d 可使用默认服务端ip及版本号",
	"SETUP":    "建立连接，入口参数为：网络协议(tcp/udp) 端口号",
	"PLAY":     "开始播放，入口参数为：开始播放的时间",
	"TEARDOWN": "断开连接",
}
