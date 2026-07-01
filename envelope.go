package yu

// Version 是协议信封的版本字段，由库统一填充。
const Version = "1.0.0"

// Request 是请求型报文（设备上报或云端下发）：method + params。
type Request struct {
	MsgId   string         `json:"msgId"`
	Version string         `json:"version"`
	Method  string         `json:"method"`
	Params  map[string]any `json:"params"`
}

// Response 是响应型报文（对上报或下发的回复）：code + data + message。
type Response struct {
	MsgId   string         `json:"msgId"`
	Version string         `json:"version"`
	Method  string         `json:"method,omitempty"`
	Code    int            `json:"code"`
	Message string         `json:"message,omitempty"`
	Data    map[string]any `json:"data,omitempty"`
}
