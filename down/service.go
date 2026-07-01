package down

// ServiceRequest 解析云端下发的服务调用请求。
type ServiceRequest struct {
	MsgId   string        `json:"msgId"`
	Version string        `json:"version"`
	Method  string        `json:"method"` // 如 "control"、"1007"
	Params  ServiceParams `json:"params"`
}

// ServiceParams 是服务调用请求的参数。
// 注意 JSON 键为 "value"（沿用线上协议）。
type ServiceParams struct {
	Values map[string]any `json:"value"`
}
