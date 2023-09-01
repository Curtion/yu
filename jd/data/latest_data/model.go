package latestData

type Propertie struct {
	Value string `json:"value"`
	Time  int64  `json:"time"`
}

type Event struct {
	Value      map[string]interface{} `json:"value"`
	Properties map[string]interface{} `json:"properties"`
	Time       int64                  `json:"time"`
}

type Identity struct {
	Pk   string `json:"pk"`
	Name string `json:"name"`
}

type SubDevice struct {
	Identity   Identity `json:"identity"`
	Properties map[string]Propertie
}

type data struct {
	Properties map[string]Propertie
	Events     map[string]Event
	SubDevices []SubDevice
	productKey string
	deviceName string
	msgId      string
}
