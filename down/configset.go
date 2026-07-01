package down

// ConfigSetRequest 解析云端下发的配置下发请求。
type ConfigSetRequest struct {
	MsgId   string          `json:"msgId"`
	Version string          `json:"version"`
	Method  string          `json:"method"`
	Params  ConfigSetParams `json:"params"`
}

// ConfigSetParams 是配置下发请求的参数。
type ConfigSetParams struct {
	Version string         `json:"version"`
	Config  map[string]any `json:"config"`
}
