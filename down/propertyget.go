package down

type GetRequest struct {
	MsgId   string    `json:"msgId"`
	Version string    `json:"version"`
	Method  string    `json:"method"`
	Params  GetParams `json:"params"`
}

type GetParams struct {
	Properties GetProperty `json:"properties"`
}

type GetProperty struct {
	Identifier []string `json:"identifier"`
}
