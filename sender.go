package yu

import "context"

// Sender 由业务实现并注入 Device，用于发布 MQTT 报文。
// 传输层（paho 连接管理、重试等）完全由业务负责，yu 仅依赖此最小接口。
type Sender interface {
	Publish(ctx context.Context, topic string, payload []byte) error
}
