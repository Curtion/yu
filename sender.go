package yu

import "context"

// Sender 是 Device 唯一的传输注入点；连接管理、重试等由业务负责，库只依赖此最小接口。
type Sender interface {
	Publish(ctx context.Context, topic string, payload []byte) error
}
