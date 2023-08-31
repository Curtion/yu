package json

const (
	Version = "1.0"
)

type HttpRequest struct {
	MsgId   string                 `json:"msgId"`
	Version string                 `json:"version"`
	Method  string                 `json:"method"`
	Params  map[string]interface{} `json:"params"`
}
