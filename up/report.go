// Package up 定义设备上行（主动发起）报文的强类型参数。
package up

type Property struct {
	Value any   `json:"value"`
	Time  int64 `json:"time"`
}

type Event struct {
	Value      map[string]any `json:"value"`
	Properties map[string]any `json:"properties"`
	Time       int64          `json:"time"`
}
