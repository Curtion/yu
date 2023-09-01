package jd

type HttpRequest struct {
	MsgId   string                 `json:"msgId"`
	Version string                 `json:"version"`
	Method  string                 `json:"method"`
	Params  map[string]interface{} `json:"params"`
}

type HttpRequest2 struct {
	MsgId   string                 `json:"msgId"`
	Version string                 `json:"version"`
	Code    int64                  `json:"code"`
	Data    map[string]interface{} `json:"data"`
	Message string                 `json:"message"`
}

type HttpRequest3 struct {
	MsgId   string        `json:"msgId"`
	Version string        `json:"version"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}
