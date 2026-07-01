// Package up 定义设备上行（主动发起）报文的强类型参数。
package up

// Property 是单个属性的值与时间戳，用于属性上报与属性查询回复。
type Property struct {
	Value any   `json:"value"`
	Time  int64 `json:"time"`
}

// Event 是单个事件的载荷。
type Event struct {
	Value      map[string]any `json:"value"`
	Properties map[string]any `json:"properties"`
	Time       int64          `json:"time"`
}
