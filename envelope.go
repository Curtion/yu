package yu

// Version 固定为 "1.0.0"，对应线上 payload 约定；与协议文档版本号 V1.3.7 无关，勿按文档"修正"。
const Version = "1.0.0"

type Request struct {
	MsgId   string         `json:"msgId"`
	Version string         `json:"version"`
	Method  string         `json:"method"`
	Params  map[string]any `json:"params"`
}

type Response struct {
	MsgId   string         `json:"msgId"`
	Version string         `json:"version"`
	Method  string         `json:"method,omitempty"`
	Code    int            `json:"code"`
	Message string         `json:"message,omitempty"`
	Data    map[string]any `json:"data,omitempty"`
}
