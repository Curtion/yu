package config

type Identity struct {
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
