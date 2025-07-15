package propertyset

type Identity struct {
	Pk   string `json:"pk"`
	Name string `json:"name"`
}

type PropertySet struct {
	MsgId   string `json:"msgId"`
	Version string `json:"version"`
	Method  string `json:"method"`
	Params  params `json:"params"`
}

type params struct {
	Properties []Propertie `json:"properties"`
	SubDevices []SubDevice `json:"subDevices"`
}

type Propertie struct {
	Identifier string `json:"identifier"`
	Value      any    `json:"value"`
}

type SubDevice struct {
	Identity   Identity    `json:"identity"`
	Properties []Propertie `json:"properties"`
}

type data struct {
	properties []Propertie
	subDevices []SubDevice
	method     string
	productKey string
	deviceName string
	msgId      string
	code       int64
	message    string
}
