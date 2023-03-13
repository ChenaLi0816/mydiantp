package dianRequest

type DianRequest struct {
	Method  string            `json:"method"`
	Url     string    `json:"url"`
	Version string            `json:"version"`
	CSeq    int64            `json:"cseq"`
	Body    map[string]string `json:"body"`
}

