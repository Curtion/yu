package down

type ConfigSetRequest struct {
	MsgId   string          `json:"msgId"`
	Version string          `json:"version"`
	Method  string          `json:"method"`
	Params  ConfigSetParams `json:"params"`
}

type ConfigSetParams struct {
	Version string         `json:"version"`
	Config  map[string]any `json:"config"`
}
