package config

type ConfigSetCall struct {
	MsgId   string          `json:"msgId"`
	Version string          `json:"version"`
	Method  string          `json:"method"`
	Params  configSetParams `json:"params"`
}

type configSetParams struct {
	Version    string                 `json:"version"`
	Config     map[string]interface{} `json:"config"`
	SubDevices []SubDevice            `json:"subDevices"`
}

type ConfigCheckCall struct {
	MsgId   string            `json:"msgId"`
	Version string            `json:"version"`
	Method  string            `json:"method"`
	Params  configCheckParams `json:"params"`
}

type configCheckParams struct {
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

type Identity struct {
	Pk   string `json:"pk"`
	Name string `json:"name"`
}

type data struct {
	msgId      string
	productKey string
	deviceName string
	Version    string
	Config     map[string]interface{}
	SubDevices []SubDevice
}
