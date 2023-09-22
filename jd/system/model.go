package system

type Identity struct {
	Pk   string `json:"pk"`
	Name string `json:"name"`
}

type SystemCall struct {
	MsgId   string `json:"msgId"`
	Version string `json:"version"`
	Method  string `json:"method"`
	Params  params `json:"params"`
}

type params struct {
	Identity Identity               `json:"identity"`
	Values   map[string]interface{} `json:"value"`
}

type data struct {
	msgId      string
	productKey string
	deviceName string
	data       map[string]interface{}
	code       int64
	message    string
}
