package propertyget

type Identity struct {
	Pk   string `json:"pk"`
	Name string `json:"name"`
}

type PropertyGet struct {
	MsgId   string `json:"msgId"`
	Version string `json:"version"`
	Method  string `json:"method"`
	Params  params `json:"params"`
}

type params struct {
	Properties Propertie   `json:"properties"`
	SubDevices []SubDevice `json:"subDevices"`
}

type Propertie struct {
	Identifier []string `json:"identifier"`
}

type SubDevice struct {
	Identity   Identity  `json:"identity"`
	Properties Propertie `json:"properties"`
}

type SubDeviceRes struct {
	Identity   Identity              `json:"identity"`
	Properties map[string]Properties `json:"properties"`
}

type Properties struct {
	Value string `json:"value"`
	Time  int64  `json:"time"`
}

type data struct {
	properties map[string]Properties
	subDevices []SubDeviceRes
	method     string
	productKey string
	deviceName string
	msgId      string
	code       int64
	message    string
}
