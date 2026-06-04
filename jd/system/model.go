package system

type SystemCall struct {
	MsgId   string                 `json:"msgId"`
	Version string                 `json:"version"`
	Method  string                 `json:"method"`
	Params  map[string]interface{} `json:"params"`
}

type data struct {
	msgId      string
	method     string
	productKey string
	deviceName string
	data       map[string]interface{}
	code       int64
	message    string
}
