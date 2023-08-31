package latestData

type propertie struct {
	Value string `json:"value"`
	Time  int64  `json:"time"`
}

type event struct {
	Value      map[string]interface{} `json:"value"`
	Properties map[string]interface{} `json:"properties"`
	Time       int64                  `json:"time"`
}

type data struct {
	Properties map[string]propertie `json:"properties"`
	Events     map[string]event     `json:"events"`
}
