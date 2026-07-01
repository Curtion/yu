package down

type SystemRequest struct {
	MsgId   string         `json:"msgId"`
	Version string         `json:"version"`
	Method  string         `json:"method"` // 如 "reload"、"reset"
	Params  map[string]any `json:"params"`
}
