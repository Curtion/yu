package down

// GetRequest 解析云端下发的属性查询请求。
type GetRequest struct {
	MsgId   string    `json:"msgId"`
	Version string    `json:"version"`
	Method  string    `json:"method"`
	Params  GetParams `json:"params"`
}

// GetParams 是属性查询请求的参数。
type GetParams struct {
	Properties GetProperty `json:"properties"`
}

// GetProperty 列出待查询的属性标识符。
type GetProperty struct {
	Identifier []string `json:"identifier"`
}
