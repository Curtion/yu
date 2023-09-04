package service

type Identity struct {
	Pk   string `json:"pk"`
	Name string `json:"name"`
}

type ServiceCall struct {
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
	Identity   Identity
	Properties map[string]interface{}
	Values     map[string]interface{}
	productKey string
	deviceName string
	msgId      string
	code       int64
	message    string
}
