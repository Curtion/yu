package config

type Identity struct {
	Pk   string `json:"pk"`
	Name string `json:"name"`
}

type ConfigCall struct {
	MsgId   string `json:"msgId"`
	Version string `json:"version"`
	Method  string `json:"method"`
	Params  params `json:"params"`
}

type params struct {
	Self       string            `json:"self"`
	SubDevices []ParamsSubDevice `json:"subDevices"`
}

type ParamsSubDevice struct {
	Pk   string `json:"pk"`
	Name string `json:"name"`
}

type SubDevice struct {
	Identity Identity               `json:"identity"`
	Version  string                 `json:"version"`
	Config   map[string]interface{} `json:"config"`
}

type data struct {
	msgId      string
	productKey string
	deviceName string
	Version    string
	Config     map[string]interface{}
	SubDevices []SubDevice
}
