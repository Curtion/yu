package down

// SystemRequest 解析云端下发的系统指令请求。
type SystemRequest struct {
	MsgId   string         `json:"msgId"`
	Version string         `json:"version"`
	Method  string         `json:"method"` // 如 "reload"、"reset"
	Params  map[string]any `json:"params"`
}
