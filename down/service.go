package down

type ServiceRequest struct {
	MsgId   string        `json:"msgId"`
	Version string        `json:"version"`
	Method  string        `json:"method"` // 如 "control"、"1007"
	Params  ServiceParams `json:"params"`
}

// ServiceParams 的 JSON 键为 "value"（沿用线上协议），非 "values"，勿"修正"。
type ServiceParams struct {
	Values map[string]any `json:"value"`
}
