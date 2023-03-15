package dianResponse

type ResponseOPTIONS struct {
	StatusCode int               `json:"status_code"`
	StatusMsg  string            `json:"status_msg"`
	Version    string            `json:"version"`
	CSeq       int64             `json:"cseq"`
	Method     map[string]string    `json:"method"`
}

type ResponseSETUP struct {
	StatusCode int               `json:"status_code"`
	StatusMsg  string            `json:"status_msg"`
	Version    string            `json:"version"`
	CSeq       int64             `json:"cseq"`
	Token 	   string			 `json:"token"`
}

type ResponsePLAY struct {
	StatusCode int               `json:"status_code"`
	StatusMsg  string            `json:"status_msg"`
	Version    string            `json:"version"`
	CSeq       int64             `json:"cseq"`
}
type ResponseTEARDOWN struct {
	StatusCode int               `json:"status_code"`
	StatusMsg  string            `json:"status_msg"`
	Version    string            `json:"version"`
	CSeq       int64             `json:"cseq"`
}