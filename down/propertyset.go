// Package down 定义云端下发报文的强类型解析结构与回复载荷。
package down

// SetRequest 解析云端下发的属性设置请求。
type SetRequest struct {
	MsgId   string    `json:"msgId"`
	Version string    `json:"version"`
	Method  string    `json:"method"`
	Params  SetParams `json:"params"`
}

// SetParams 是属性设置请求的参数。
type SetParams struct {
	Properties []SetProperty `json:"properties"`
}

// SetProperty 是单个待设置属性（仅 identifier + value）。
type SetProperty struct {
	Identifier string `json:"identifier"`
	Value      any    `json:"value"`
}

// SetResult 用 Res/Msg 表达执行结果，字段与 SetProperty 不同，故独立定义。
type SetResult struct {
	Identifier string `json:"identifier"`
	Res        string `json:"res"` // "1" 成功，"0" 失败
	Msg        string `json:"msg"`
}
